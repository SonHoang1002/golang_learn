package server_side

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// download: go get -u github.com/gorilla/mux

func RouteServer() {
	//create new Routes
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		fmt.Fprintf(writer, "Hello, This is %s book", title)
	}).Methods("GET")
	r.HandleFunc("/books/new/", func(writer http.ResponseWriter, request *http.Request) {
		var newParams any
		fmt.Println("Data is %s", json.NewDecoder(request.Body).Decode(newParams))
		fmt.Fprintf(writer, "Hello,data is: %s", request.Body)
	}).Methods("POST")
	port := "8080"
	fmt.Println("Listening route:  ", port)
	http.ListenAndServe(":"+port, r)
}

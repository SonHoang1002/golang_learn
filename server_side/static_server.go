package server_side

import (
	"fmt"
	"net/http"
)

func StaticServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, I am Son Server")
	})
	////
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	////
	port := "8080"
	fmt.Println("Listening static:  ", port)
	http.ListenAndServe(":"+port, nil)
}

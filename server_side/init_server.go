package server_side

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, I am Son Server")
	})
	port := "8080"
	fmt.Println("Listening ", port)
	http.ListenAndServe(":"+port, nil)
}
func GetMusicApi() {
	http.HandleFunc("/music", func(writer http.ResponseWriter, request *http.Request) {
		var data = map[string]interface{}{
			"name":   "Hello world",
			"singer": "you",
		}
		json.NewEncoder(writer).Encode(data)
	})
	port := "8080"
	fmt.Println("Listening ", port)
	http.ListenAndServe(":"+port, nil)
}

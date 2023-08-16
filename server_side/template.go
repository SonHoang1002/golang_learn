package server_side

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Name  string
	Tasks []Task
}
type Task struct {
	Title string
	Done  bool
}

func TemplateServer() {
	templ := template.Must(template.ParseFiles("template/main.html"))
	http.HandleFunc("/ui", func(writer http.ResponseWriter, request *http.Request) {
		data := Todo{
			Name: "HTS",
			Tasks: []Task{
				{
					Title: "Task 1",
					Done:  true,
				},
				{
					Title: "Task 2",
					Done:  false,
				},
				{
					Title: "Task 3",
					Done:  true,
				},
			},
		}
		err := templ.Execute(writer, data)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}

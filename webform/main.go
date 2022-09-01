package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

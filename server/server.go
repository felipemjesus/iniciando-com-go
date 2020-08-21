package main

import (
	"net/http"
	"text/template"
)

type Member struct {
	Name string
}

func hello(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("hello.html"))

	members := []Member{
		{"Felipe"},
		{"João"},
		{"Pedro"},
	}

	tmpl.Execute(w, members)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

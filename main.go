package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {

	http.HandlerFunc("/index", index)
	http.HandlerFunc("/index", about)
	http.HandlerFunc("/contact", contact)
	http.HandlerFunc("/apply", apply)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request)   {}
func about(w http.ResponseWriter, req *http.Request)   {}
func contact(w http.ResponseWriter, req *http.Request) {}
func apply(w http.ResponseWriter, req *http.Request)   {}

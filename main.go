package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/process", processor)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func processor(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fname := req.FormValue("first_name")
	lname := req.FormValue("last_name")

	d := struct {
		First string
		Last  string
	}{

		First: fname,
		Last:  lname,
	}

	err := tpl.ExecuteTemplate(w, "processor.gohtml", d)
	if err != nil {
		log.Println("Log", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println("No errors found")
	fmt.Println(req.URL.Path)
}

func index(w http.ResponseWriter, req *http.Request) {

	pd := pageData{

		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("Log", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println("No errors found")
	fmt.Println(req.URL.Path)
}
func about(w http.ResponseWriter, req *http.Request) {

	pd := pageData{

		Title: "About Page",
	}
	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func contact(w http.ResponseWriter, req *http.Request) {

	pd := pageData{

		Title: "Contact Page",
	}
	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func apply(w http.ResponseWriter, req *http.Request) {

	pd := pageData{

		Title: "Apply Page",
	}

	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

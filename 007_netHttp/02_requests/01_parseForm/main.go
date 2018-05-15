package main

import (
	"net/http"
	"log"
	"html/template"
)

type myHandler int

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d myHandler
	http.ListenAndServe(":8080", d)
}

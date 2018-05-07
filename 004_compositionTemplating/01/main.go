package main

import (
	"text/template"
	"os"
	"log"
)

// lets define a Person to pass to template

type person struct {
	Name string
	Age int
}

// we initialise the pointer to the template
var tpl *template.Template

// we parse all the templates and store in the pointer to the template
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Name: "Luis Benavides",
		Age: 26,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}

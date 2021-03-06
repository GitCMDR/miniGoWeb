package main

import (
	"text/template"
	"os"
	"log"
)

type person struct {
	Name string
	Age int
}

type doubleZero struct {
	person
	LicenceToKill bool
}


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleZero{
		person {
			Name: "Luis",
			Age: 26,
		},
		false,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
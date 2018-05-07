package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// Create a slice (like Python list) to pass to variables.
	vips := []string{"Luis", "Ellie", "Novee", "Mish"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", vips)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println() // New line
}

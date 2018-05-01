package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("one.gohtml") // parse file into tpl variable
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) // execute tpl
	if err != nil {
		log.Fatalln(err)
	}

	// parse new files into tpl
	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// call specific parsed file from tpl container, write to stdout with nil data
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

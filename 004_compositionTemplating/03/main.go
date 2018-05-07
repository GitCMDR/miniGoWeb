package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"IFN600", "Research", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"IFN700", "Project Management", "4"},
				{"IFN701", "Project 1", "4"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{"IFN702", "Project 2", "4"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}

}

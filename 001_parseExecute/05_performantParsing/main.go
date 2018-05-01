// It is important to be performing in programming and that
// is why we only want to parse the templates once (at the
// beginning of our program). We can do this by putting the parsing
// inside the func init() of our program. We will call parse glob
// within the func Must() - func Must() comes from within the "text/
// template" package and it takes care of the error checking so that
// we avoid doing the error checking ourselves.

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template


// template.ParseGlob takes a folder and returns *Template and Error
// template.Must takes a *Template and Error and returns *Template only
// template.Must saves us doing Error checking ourselves.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
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

}

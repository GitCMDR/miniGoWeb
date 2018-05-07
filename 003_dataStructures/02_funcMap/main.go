package main

import (
	"text/template"
	"strings"
)

// We can pass functions to the template so that we can use them there
// We do that using template.FuncMap

var tpl *template.Template

var functionsPassed = template.FuncMap{
	"uc": strings.ToUpper,  // Predefined and globally available.
	"ft": firstFour, // A function that I create for passing
}

func init() {
	tpl = template.Must(template.New("").Funcs(functionsPassed).ParseFiles("tpl.gohtml"))
}

func firstFour(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}


package main

import (
	"net/http"
	"io"
)

type hotdog int

func (d hotdog) ServeHTTP (rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "dog dog dogggy")
}

type hotcat int

func (c hotcat) ServeHTTP (rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "cat cat kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
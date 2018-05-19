package main

import (
	"net/http"
	"io"
)

//instead of defining super weird types or calling the ServeHTTP explicitly we can do something
//way more elegant!

// create a function for each view (yaaas)
func helloWorld(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "Hello World! I did this myself :V")
}

func main() {
	http.HandleFunc("/", helloWorld)

	http.ListenAndServe(":8080", nil) // you pass a nil handler so that http uses default mux
}

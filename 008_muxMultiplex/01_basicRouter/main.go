package _1_basicRouter

import (
	"net/http"
	"io"
)

// multiplexing is just routing (if this url in header then this code)

type hotdog int

func (m hotdog) ServeHTTP (rw http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(rw, "doggy doggy doggy")
	case "/cat":
		io.WriteString(rw, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
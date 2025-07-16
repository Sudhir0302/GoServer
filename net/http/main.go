package main

import (
	"fmt"
	"net/http"
)

type server struct {
	addr string
}

// actual implementation of ServeHTTP(http.ResponseWriter, *http.Request) occurs.Since this method statisfies that handler func.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello from server"))
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("hello from home endpoint"))
			return
		case "/test":
			w.Write([]byte("hello from test endpoint"))
			return
		default:
			w.Write([]byte("404 not found"))
			return
		}
	}
}

// so http.ListenAndServe() takes two arguments ,one is port no and the other is a handler which handles the
// incoming req. so here i have manually created a struct which has addr(port number)
// and that struct actually implements the handler func (ServeHTTP)
func main() {
	fmt.Println("net/http")

	// http.ListenAndServe(":8080", nil)
	s := &server{
		addr: ":8080",
	}
	http.ListenAndServe(s.addr, s)
}

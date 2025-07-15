package main

import (
	"fmt"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from server"))
}
func main() {
	fmt.Println("net/http")

	// http.ListenAndServe(":8080", nil)
	s := &server{
		addr: ":8080",
	}
	http.ListenAndServe(s.addr, s)
}

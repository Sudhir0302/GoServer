package main

import (
	"fmt"
	"net/http"
)

func call(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from call func"))
}

func main() {
	fmt.Println("HandlerFunc")

	//adapter pattern
	// res := compute(8, 2, func(a, b int) int { return a - b })
	// res := compute(8, 2, add)
	// fmt.Println(res)

	//creating a new ServeMux router
	// mux := http.NewServeMux()

	http.HandleFunc("/call", call) //registers in the Servemux's map(a place where all endpoints are mapped and stored inside the mux struct)

	// http.ListenAndServe(":8080", &custom{})
	http.ListenAndServe(":8080", nil)

}

// Interfaces in Go are implemented by any type — struct, function type, or even custom type aliases — as long as the type defines all the methods in the interface. No explicit declaration is needed.
// type custom struct {
// }

//custom ServerHTTP handler
// func (*custom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello from custom handler"))
// }

// adapter pattern
type operation func(a int, b int) int

func add(a int, b int) int {
	return a + b
}

func compute(a int, b int, op operation) int {
	return op(a, b)
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
// 	if r.RequestURI == "*" {
// 		if r.ProtoAtLeast(1, 1) {
// 			w.Header().Set("Connection", "close")
// 		}
// 		w.WriteHeader(StatusBadRequest)
// 		return
// 	}
// 	var h Handler
// 	if use121 {
// 		h, _ = mux.mux121.findHandler(r)
// 	} else {
// 		h, r.Pattern, r.pat, r.matches = mux.findHandler(r)
// 	}
// 	h.ServeHTTP(w, r)
// }

// The above code is where the actual function is found and called.
// Like if we define http.HandleFunc("/home",func()..) means the go's default ServeMux registers the
// handler func with its endpoint inside a map.Then if a req hits for a specific endpoint means ServeMux's default
// ServeHTTP method is called, where it finds the crct handler func using several built-in methods(look inside server.go file) ,
// if found it calls the respective method

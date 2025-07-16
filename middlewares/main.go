package main

import (
	"fmt"
	"net/http"
)

func groot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from groot"))
}

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Req: ", r.URL)
		next.ServeHTTP(w, r)
	})
}

// so http.HandleFunc() directly registers the handler in mux.m and http.HandlerFunc() returns
// only the ServeHTTP(w http.ResponseWriter, r *http.Request) i.e http.Handler, so we can pass
// it as an actual handler to the http.HandleFunc() which registers.

func main() {
	fmt.Println("middlewares")
	// wrapped := mw(call)
	// fmt.Println(wrapped("sudhir"))

	// http.HandleFunc("/groot", groot)
	handler := log(http.HandlerFunc(groot))
	http.Handle("/groot", handler)
	http.ListenAndServe(":8080", nil)
}

func call(name string) string {
	return name + "!!!"
}
func mw(f func(string) string) func(string) string {
	return func(name string) string {
		fmt.Println("hello ", name)
		res := f(name)
		return res
	}
}

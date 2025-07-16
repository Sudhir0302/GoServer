package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login successfull"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from home"))
}

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log: ", r.URL)
		next.ServeHTTP(w, r)
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth middlewareee")
		next.ServeHTTP(w, r)
	})
}

func chain(h http.Handler, middle ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middle) - 1; i >= 0; i-- {
		h = middle[i](h) //it's like calling auth(h) ,here h is initially the finalHandler (here home).
		// at the end of the loop ,h will be wrapped with chain middlewares passed.
		// for eg:log(auth(home))
	}
	return h
}

func main() {

	fmt.Println("Server up and running")

	mux := http.NewServeMux()

	mux.HandleFunc("/login", login)

	// handler := auth(http.HandlerFunc(home))
	// http.Handle("/home", handler)

	mux.Handle("/home", chain(
		http.HandlerFunc(home),
		log,
		auth,
	))

	http.ListenAndServe(":8080", mux)
}

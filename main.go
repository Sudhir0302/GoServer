package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("serverrr")
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8081", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","Application/json")
	w.Write([]byte("hello from golang"))
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {

	//setting header
	w.Header().Set("Content-Type","application/json")

	fmt.Println("greet functionn")
	//sends the response automatically
	json.NewEncoder(w).Encode("hello from golang")
}
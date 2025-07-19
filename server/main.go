package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server..")

	http.HandleFunc("/",home)

	http.ListenAndServe(":8080",nil)
}

func home(w http.ResponseWriter,r *http.Request){
	
	msg:=map[string]string{"message":"hello from server"}

	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(msg)
}

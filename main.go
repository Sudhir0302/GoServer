package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	w.Header().Set("Content-Type", "application/json")
	data := make(map[string]string)
	data["name"] = "sudhir"
	// json.NewEncoder(w).Encode(data)
	fmt.Fprintf(w, `{"name":"sudhir"}`)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"msg":"welcome to goo"}`)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := Student{"kumar", 21}
	bytedata, _ := json.MarshalIndent(s, "", "\t")
	fmt.Println(string(bytedata))
	json.NewEncoder(w).Encode(s)
}

func main() {
	fmt.Println("go serverr")
	http.HandleFunc("/test", test)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/getUser", getUser)
	http.ListenAndServe(":8080", nil)
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

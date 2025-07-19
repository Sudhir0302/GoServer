package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greet", greetHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

// GET /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Go REST API!"))
}

// POST /greet
func greetHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	name := r.URL.Query().Get("name")
	if name == "" {
		// fallback to body
		var req struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Name == "" {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		name = req.Name
	}

	res := map[string]string{"message": "Hello, " + name + "!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

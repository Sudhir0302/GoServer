package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("gorilla muxx")

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	data:=make(map[string]string)
	data["msg"]="hellooo"
	// w.Write([]byte(data))  wrong way to send json data

	//type 1
	// jsonbytes,_:=json.Marshal(data)
	// w.Write(jsonbytes)

	// type 2

	json.NewEncoder(w).Encode(data)
}


// 🔄 Flow: From Client to Handler
// Browser/User visits http://localhost:8080/

//          ⬇️
// Go's internal HTTP server accepts the connection.

//          ⬇️
// It checks the registered route handlers (like "/").

//          ⬇️
// Finds that "/" is handled by `home(w, r)`.

//          ⬇️
// Calls:
//    home(w, r)
// Where:
//    w = ResponseWriter (used to write response back to client)
//    r = *Request (contains all HTTP request data)
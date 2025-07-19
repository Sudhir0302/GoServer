package main

import (
	"fmt"
	"log"
	"mongodbapi/routes"
	"net/http"
)

func main() {
	// fmt.Println("hello dbb")
	fmt.Println("server is running")
	r:=routes.Router()
	log.Fatal(http.ListenAndServe(":8080",r))

}

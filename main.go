package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sudhir0302/GoServer/routes"
)

func main() {
	fmt.Println("go serverrr startss")
	log.Fatal(http.ListenAndServe(":8080",routes.UserRoute()))
}

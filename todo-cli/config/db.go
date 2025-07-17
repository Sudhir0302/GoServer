package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:sudhir@tcp(127.0.0.1:3306)/todo_cli")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected")
	return db
}

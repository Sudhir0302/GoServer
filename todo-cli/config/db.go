package config

import (
	// database/sql is the standard interface layer, and the chosen driver (here mysql) is the engine that connects to the actual database.
	"database/sql"
	"fmt"
	"log"

	// use underscore( _ ) to run only side effects like init()
	// init() calls sql.Register("mydriver", ...)
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

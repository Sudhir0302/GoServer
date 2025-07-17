package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Sudhir0302/Go/todo-cli/models"
)

type TodoController struct {
	DB *sql.DB
}

func (repo *TodoController) Create() {
	check, _ := repo.DB.Exec(`SHOW TABLES LIKE "todo"`)
	if check != nil {
		return
	}
	query := `CREATE TABLE todo(
		id int auto_increment primary key,
		todo varchar(200) not null,
		status boolean default false
	);`
	_, err := repo.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created")
}

func (repo *TodoController) Add(todo string) string {
	stmt, _ := repo.DB.Prepare("INSERT INTO todo (todo,status) VALUES(?,false)")
	res, _ := stmt.Exec(todo)
	if res != nil {
		return "succes"
	}
	return "failed"
}

func (repo *TodoController) View() {
	row, _ := repo.DB.Query(`select * from todo`)

	// var

	for row.Next() {
		// fmt.Println(res.Next())
		var t models.Todo
		row.Scan(&t.ID, &t.Task, &t.Status)
		fmt.Println(t)
	}
}

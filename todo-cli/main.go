package main

import (
	"fmt"

	"github.com/Sudhir0302/Go/todo-cli/config"
	"github.com/Sudhir0302/Go/todo-cli/controllers"
)

func main() {

	db := config.Conn()
	defer db.Close()

	cont := controllers.TodoController{DB: db}
	cont.Create()

	fmt.Println("todo cli")
	fmt.Println("Enter a number to continue: ")
	// fmt.Println("1 - view todos 2-add 3-update 4-delete 5-exit")
	for {
		fmt.Println("1 - view todos 2-add 3-update 4-delete 5-exit")
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("view todos")
			cont.View()
		case 2:
			fmt.Println("enter a task: ")
			var todo string
			fmt.Scan(&todo)
			res := cont.Add(todo)
			fmt.Println(res)
		case 3:
			fmt.Println("update")
		case 4:
			fmt.Println("delete")
		case 5:
			fmt.Println("exit-ingg")
			return
		}
	}
}

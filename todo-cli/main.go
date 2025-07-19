package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sudhir0302/Go/todo-cli/config"
	"github.com/Sudhir0302/Go/todo-cli/controllers"
)

func main() {

	db := config.Conn()
	reader := bufio.NewReader(os.Stdin)
	defer db.Close()

	cont := controllers.TodoController{DB: db}
	cont.Create()

	fmt.Println("todo cli")
	// fmt.Println("Enter a number to continue: ")
	// fmt.Println("1 - view todos 2-add 3-update 4-delete 5-exit")
	for {
		fmt.Println("Enter a number to continue: 1 - view todos, 2-add, 3-update, 4-delete, 5-exit")
		var op int
		fmt.Scan(&op)
		reader.ReadString('\n')
		switch op {
		case 1:
			fmt.Println("view todos")
			cont.View()
		case 2:
			fmt.Println("enter a task: ")
			todo, _ := reader.ReadString('\n')
			todo = strings.TrimSpace(todo)
			res := cont.Add(todo)
			fmt.Println(res)
		case 3:
			// fmt.Println("enter a task to update: ")
			fmt.Print("enter task name: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			res := cont.Update(task)
			fmt.Println(res)
			// fmt.Println("update")
		case 4:
			fmt.Print("enter a task to delete: ")
			todo, _ := reader.ReadString('\n')
			// fmt.Println(todo)
			todo = strings.TrimSpace(todo)
			res := cont.Delete(todo)
			fmt.Println(res)
		case 5:
			fmt.Println("exit-ingg")
			time.Sleep(2 * time.Second) //2 sec delay
			return
		default:
			fmt.Println("please enter a valid number to continue")
		}
	}
}

package models

import "fmt"

type Todo struct {
	ID     int
	Task   string
	Status bool
}

//to change the default behavior printing struct
func (t Todo) String() string {
	return fmt.Sprintf("Task #%d %s %t", t.ID, t.Task, t.Status)
}

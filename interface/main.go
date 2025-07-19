package main

import "fmt"

type Name interface {
	getName()
}
type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name string
	Age  int
}

func (std Student) getName() {
	fmt.Println(std.Name)
}

func (tch Teacher) getName() {
	fmt.Println(tch.Name)
}

func main() {
	// fmt.Println("interfacess")

	std := Student{"kumar", 12}
	tch := Student{"teacher", 32}
	// std.getName()
	// tch.getName()
	obj := []Name{std, tch}
	for _, d := range obj {
		d.getName()
	}
	obj1 := dummy{}
	obj1.call()
}
func (d dummy) call() {
	fmt.Println("hii")
}

type Custom interface {
	call()
}

type dummy struct {
}

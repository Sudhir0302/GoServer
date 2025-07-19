package main

import "fmt"

func main() {
	fmt.Println("pointerss")
	var a int=2
	var ptr *int=&a
	// var ptr1 **int=&ptr;
	fmt.Println(ptr)
	fmt.Println("a:",a)
	call(&a)
	fmt.Println("a:",a)
	name:=fmt.Sprint("hello ",a)
	fmt.Println(name)

	var p *int=&a
	if p!=nil {
		fmt.Println(p)
	}

	fmt.Println(fmt.Sprint("hello"))
}

func call(n *int) int{
	*n+=1
	return 1;
}
package main

import "fmt"

func main() {
	defer fmt.Println("helloo")
	fmt.Println("deferss")
	defer call()
}


func call(){
	for i:=range(5) {
		defer fmt.Println(i)
	}
}

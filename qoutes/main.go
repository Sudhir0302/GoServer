package main

import (
	"fmt"
	"math/rand"

	"rsc.io/sampler"
)

func main() {
	fmt.Println("qoutess")
	fmt.Println(sampler.Hello())
	a,b:=Call()
	fmt.Println(a,b)

	arr:=[]int{}
	arr=append(arr, 1,2,3)
	fmt.Println(arr)

	fmt.Println(rand.Intn(10))
}


func Call()(string,string){
	return "1","2"
}
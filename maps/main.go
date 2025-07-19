package main

import "fmt"

func main() {
	fmt.Println("mapss")

	data:=make(map[string]string)
	data["kumar"]="tir"
	data["kumar1"]="cbe"
	data["kumar"]="cbe"

	fmt.Println(data["kumar"])

	for k,v :=range(data){
		fmt.Println(k,v)
	}
}

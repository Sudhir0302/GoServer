package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("hello ddosss")

	// for i:=1 ; i<=2 ; i++{
	// 	res,err:=http.Get("http://localhost:8080/test");
	// 	if err!=nil{
	// 		panic(err)
	// 	}
	// 	bytedata,_:=io.ReadAll(res.Body)
	// 	fmt.Println(string(bytedata))
	// }
	for{
		res,err:=http.Get("http://localhost:8080/test");
		if err!=nil{
			panic(err)
		}
		bytedata,_:=io.ReadAll(res.Body)
		fmt.Println(string(bytedata))
	}
}

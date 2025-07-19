package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("filesss")
	content:="hello,this is from golang"

	file,err:=os.Create("files1.txt");
	
	if err!=nil {
		panic(err)
	}

	length,err:=io.WriteString(file,content)

	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(length);
	defer file.Close();
	readFile("files1.txt")
}

func readFile(filepath string){
	data,err:=os.ReadFile(filepath)   //ioutil.ReadFile() is depreacated
	if(err!=nil){
		panic(err)
	}
	fmt.Println(string(data))
}

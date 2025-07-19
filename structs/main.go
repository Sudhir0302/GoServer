package main

import "fmt"

func main() {
	fmt.Println("structss")
	user1:=User{"sudhir","sudhir@gmail.com",21}
	fmt.Println(user1)
	fmt.Printf("%+v\n",user1)
	fmt.Printf("%+v\n",user1.Name)
	fmt.Println(add(1546,2324))
	add1(1,2,2,3)
	fmt.Println("User Name : ",user1.getName());

	var user User
	user.Name="sudhir"
	fmt.Println(user) 

	var user2 *User=&user
	user2.Name="kumar"
	fmt.Println(user)
	fmt.Printf("%p\n",user2)  //prints the memory address of the value which the pointer is pointing..
	fmt.Printf("%p\n",&user2)  //prints the memory address of the pointer itself

}

type User struct {
	Name  string 
	Email string
	age   int
}


func (u *User) getName()string{
	// fmt.Println(u.Name);
	return u.Name
}


func add(a ,b int) int {
	return a+b
}

func add1(a...int){
	for _,d:=range(a){
		fmt.Println(d);
	}
}
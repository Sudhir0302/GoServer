package main

import (
	"fmt"
	"net/url"
	"strconv"
)

const myurl="https://jsonplaceholder.typicode.com/users?id=1"

func main() {
	fmt.Println("urlss")
	fmt.Println(myurl)

	//parsing
	res,_:=url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Port())
	params:=res.Query();
	// fmt.Println(params["id"])
	id, _ := strconv.Atoi(params.Get("id"))
	fmt.Println(id)

	
	//construct url

	// partsofurl:=&url.URL{
	// 	Scheme: "https",
	// 	Host:	"jsonplaceholder.typicode.com",
	// 	Path: "/users",
	// }

	var partsofurl *url.URL=&url.URL{
		Scheme: "https",
		Host:	"jsonplaceholder.typicode.com",
		Path: "/users",
	}

	// partsofurl:=url.URL{
	// 	Scheme: "https",
	// 	Host:	"jsonplaceholder.typicode.com",
	// 	Path: "/users",
	// }

	myurl2:=partsofurl.String();
	fmt.Println(myurl2)
}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// const url="https://jsonplaceholder.typicode.com/users"
// func main() {
// 	fmt.Println("web reqq")
// 	res,err:=http.Get(url)
// 	if(err!=nil){
// 		panic(err)
// 	}
// 	data,err:=io.ReadAll(res.Body)
// 	if(err!=nil){
// 		panic(err)
// 	}
// 	fmt.Println(string(data));
// }

// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func main() {
// 	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
// 		users := []User{{1, "Alice"}, {2, "Bob"}}
// 		json.NewEncoder(w).Encode(users)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "strings"
)

// const myurl="http://localhost:8081/greet";
const myurl="https://jsonplaceholder.typicode.com/users";
func main(){
	fmt.Println("req to go server")
	// content:=strings.NewReader(`
	// 	{ 
	// 		"msg":"hi from req"
	// 	}
	// `)
	// res,_:=http.Post(myurl,"Application/json",content)
	res,_:=http.Get(myurl)

	data,_:=io.ReadAll(res.Body)
	// fmt.Println(data)
	//data will be in byte format

	// fmt.Println(string(data))

	var user []map[string]interface{}

	var err=json.Unmarshal(data,&user)

	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(user[1]);
	

	defer res.Body.Close();  //callers responsibility to close the connection..

}
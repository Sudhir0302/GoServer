package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const goserver = "http://localhost:8081"

func main() {
	fmt.Println("helloo")
	getdata()
}

func getdata() {

	res, err := http.Get(goserver + `/greet`)

	if err != nil {
		panic(err)
	}

	bytedata, _ := io.ReadAll(res.Body)

	content := string(bytedata)

	fmt.Println(content)

	//using map[string]interface{}
	var data map[string]any //any is an alias for interface{} and is equivalent to interface{} in all ways.

	_ = json.Unmarshal(bytedata, &data)

	fmt.Println(data["message"])

	//marshal

	jsondata := Data{Message: "hello"}

	// jsonbytes,_:=json.Marshal(jsondata)
	jsonbytes, _ := json.MarshalIndent(jsondata, "", " \t")

	jsonString := string(jsonbytes)

	fmt.Println(jsonString)

	fmt.Println(json.Valid(jsonbytes)) //to check if the data is valid json or not

}

type Data struct {
	Message string `json:"message"`
}

// 2 types to read json data, using struct if json structure known means,
// and the other is using map[string]interface{} when json structure unknown means...

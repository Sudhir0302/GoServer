package main

import (
	// "encoding/json"
	"crypto/sha256"
	"fmt"
	// "time"
)

const ( //just like enums in java
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	// fmt.Println("pracc")

	// a := []int{}
	// a = append(a, 1, 2, 3)
	// for _, d := range a {
	// 	fmt.Print(d, " ")
	// }
	// fmt.Println()
	// fmt.Print(a)
	// fmt.Println(call())

	// var num int = 10
	// fmt.Println(num)
	// var p *int = &num
	// fmt.Println(*p)
	// add(p)
	// fmt.Println(*p)

	// jsondata := []byte{97, 98}
	// fmt.Println(string(jsondata))
	// data := []byte(`{"name":"kumar"}`)

	// var a map[string]any
	// _ = json.Unmarshal(data, &a)
	// fmt.Println(a["name"])

	// fmt.Println(A, B, C)
	// a := 1
	// fmt.Printf("%T", a)

	// std := Student{"kumar", 32}
	// fmt.Println(std)
	// std.call()
	// std.changeName()
	// fmt.Println(std)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%d ", i)
	// }

	// call(10)

	// //just like try-catch
	// defer func() {
	// 	if r := recover(); r != nil {  //recover() catchs that error
	// 		fmt.Println(r)
	// 	}
	// }()
	// panic("error occured") //throw new error

	// sum1 := sum(10, call)
	// fmt.Println(sum1(10))

	//marshal
	// std := Student{"sudhir", 21}
	// res, _ := json.MarshalIndent(std, "", "\t")
	// fmt.Println(string(res))

	//unmarshal
	// jsondata := `{"name":"sudhir"}`
	// var data map[string]any
	// json.Unmarshal([]byte(jsondata), &data)
	// fmt.Println(data["name"])

	// start := time.Now()
	// for i := 1; i <= 1e9; i++ {

	// }
	// end := time.Since(start)
	// fmt.Println(end)

	// var sm sync.Map
	// sm.Store("name", "kumar")
	// val, ok := sm.Load("key")
	// fmt.Println(val, ok)

	// op := os.Getenv("Path")
	// fmt.Println(op)
	// call := outer()
	// call()

	hash := sha256.Sum256([]byte("sudhir"))
	fmt.Printf("%x", hash)
}

func outer() func() {
	x := 10
	return func() { fmt.Println(x) } // x escapes
}

func sum(a int, call func(int)) func(int) int {
	// call(10)
	return func(x int) int {
		return x + a
	}
}

func call(a int) {
	if a == 0 {
		return
	}
	fmt.Println(a)
	call(a - 1)
}

type Teacher struct {
	Name string
	Age  int
}
type Student struct {
	Name string
	Age  int
}

func (std Student) call() {
	fmt.Println(std.Name)
}

func (std *Student) changeName() {
	std.Name = "ajayyy"
}

// func call() string {
// 	fmt.Println("\ncall")
// 	return "hhii"
// }

func add(a *int) {
	*a = *a + 10
}

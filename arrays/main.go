package main

import "fmt"

func main() {
	// a := [3]int{1, 2, 3}
	// var a = []int{1, 2}
	// var b=[]int{10,20}
	// a=append(a, 1,23);
	// a=append(a,b...);
	// fmt.Println(a)
	// fmt.Println(len(a))
	// for i,d:=range a{
	// 	fmt.Println(i," ",d)
	// }
	
	// slc:=a[:len(a)-1]
	// fmt.Println(slc)
	// hello()

	// In Go, the ... is called the variadic expansion operator.


	// a:=[]int{1,2}
	// a=append(a, []int{4,5}...)
	// ind:=2;
	// a=append(a[:ind],append([]int{3},a[ind:]...)...);  //inserting a value at specific index using slice.
	// fmt.Println(a)
	// fmt.Printf("%#v",a)


	// data:=[]int{}
	// data=append(data, 1,2,3)
	// fmt.Println(data)
	// fmt.Println(data[1:])
	// data=append(data, data[1:]...)

	// anydata:=make([]any, len(data))

	// for i,d :=range data {
	// 	anydata[i]=d
	// }
	// fmt.Println(anydata...)

	// var a=[3]int{1,3,2}
	// fmt.Println(a)

	var b=[]int{1}
	b=append(b, []int{1,2,4}...)
	fmt.Println(b)

	b=[]int{};
	fmt.Println(b)	
} 

func hello(){
	fmt.Println("hello");
}
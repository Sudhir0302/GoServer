package main

import (
	"fmt"
	"sync"
)

func main() {

	// defer fmt.Println("hiii")
	fmt.Println("mutexxx")

	wg:=&sync.WaitGroup{}
	var score=[]int{0}
	//anonymous 
	
	wg.Add(3) //defining how many go routines
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		fmt.Println("1 R")
		score=append(score, 1)	
	}(wg)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		fmt.Println("2 R")
		score=append(score, 2)	
	}(wg)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		fmt.Println("3 R")
		score=append(score, 3)	
	}(wg)

	wg.Wait()
	fmt.Println(score)
}

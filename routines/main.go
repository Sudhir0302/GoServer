package main

import (
	"fmt"
	"sync"
)

// threads - managed by os.Fixed size - 1mb
// goroutines - managed by Go runtime.Flexible size - 2kb,if exceeds doubles the size.
// goroutines and threads are M:N mapping, i.e one or more goroutine is managed by a single os thread.
func main() {
	fmt.Println("routiness")
	var wg sync.WaitGroup
	wg.Add(1)
	go call("helloo", &wg)
	for i := 1; i <= 1e9; i++ {
		// fmt.Print(i)
		if i == 1090190 {
			fmt.Println(i)
		}
	}
	fmt.Println()
	wg.Wait()
}

func call(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(name)
	}
}

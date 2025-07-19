package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	for val := range ch { // keep receiving until channel is closed
		fmt.Println("Received:", val)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("Worker done.")
}

func main() {
	ch := make(chan int) // unbuffered channel

	go worker(ch) // launch worker goroutine

	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		ch <- i // blocks until worker receives
	}

	close(ch) // signal worker that no more values will be sent

	time.Sleep(time.Second) // give time for last print
	fmt.Println("Main done.")
}

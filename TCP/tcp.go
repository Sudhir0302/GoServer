package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//opens a tcp socket at port 8081
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Server is running")
	//waits for another tcp socket to connect
	conn, _ := ln.Accept()
	defer conn.Close()

	// buf := make([]byte, 4096)
	// n, _ := conn.Read(buf)
	// fmt.Println(string(buf[:n]))
	n, _ := conn.Write([]byte("hii from tcp socket"))
	fmt.Println(n)
}

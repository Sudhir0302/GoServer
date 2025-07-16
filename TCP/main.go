package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Server is running")
	conn, _ := ln.Accept()
	defer conn.Close()

	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("net")

	//connects to the tcp socket with given address and returns conn
	conn, _ := net.Dial("tcp", "localhost:8081")

	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.LocalAddr().Network())

	//http req
	// fmt.Fprintf(conn, "GET /test HTTP/1.1\r\nHost: localhost\r\n\r\n")
	// fmt.Println(len("GET /test HTTP/1.1\r\nHost: localhost\r\n\r\n"))
	// n1, _ := conn.Write([]byte("GET /test HTTP/1.1\r\nHost: localhost\r\n\r\n"))
	// fmt.Println(n1)

	//read req
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf[:n]))
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			fmt.Println(string(buf))
		}
		if err != nil {
			break
		}
	}
}

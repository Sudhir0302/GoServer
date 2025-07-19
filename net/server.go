package main

import (
	"bufio"
	"fmt"
	"net"
)

var clients []net.Conn

func handle(conn net.Conn) {
	defer conn.Close()
	clients = append(clients, conn)

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected:", conn.RemoteAddr())
			return
		}

		// Broadcast to all
		for _, c := range clients {
			if c != conn {
				c.Write([]byte(msg))
			}
		}
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	fmt.Println("Chat server started on :8080")

	for {
		conn, _ := ln.Accept()
		fmt.Println("New connection:", conn.RemoteAddr())
		go handle(conn)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error binding port 6379 - ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Gedis started!")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error stablishing connection - ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		if _, err := conn.Read([]byte{}); err != nil {
			fmt.Println("Error reading from client: ", err.Error())
			continue
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}

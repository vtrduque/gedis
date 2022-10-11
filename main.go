package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	conn.Write([]byte("hey! :D"))
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Error binding port 6379 - ", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error stablishing connection - ", err.Error())
		}

		go handleConnection(conn)
	}
}

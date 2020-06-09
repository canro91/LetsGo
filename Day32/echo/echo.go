package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handler(conn net.Conn) {
	conn.Write([]byte("Welcome! Say anything...\n"))
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Close()
			break
		}

		trimmed := strings.TrimSpace(message)
		if trimmed == "ping" {
			conn.Write([]byte("pong\n"))
		} else {
			fmt.Printf(message)
			conn.Write([]byte(fmt.Sprintf("You said: %s", message)))
		}
	}
}

func main() {
	fmt.Println("Starting...")
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}
}

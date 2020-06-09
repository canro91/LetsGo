package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting...")
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listener.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf(message)
	}
}

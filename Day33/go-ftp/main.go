package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
)

// TODO: Take a look at tp.ReadResponse
func readResponse(conn net.Conn) (string, string, error) {
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	var code string
	line, _ := tp.ReadLine()
	if line[3:4] == "-" {
		code = line[:3]
		for {
			nextLine, err := tp.ReadLine()
			if err != nil {
				return "", "", err
			}

			line += "\n" + strings.TrimRight(nextLine, "\r\n")
			if nextLine[:3] == code && nextLine[3:4] != "-" {
				break
			}
		}
	}

	return code, line, nil
}

func main() {
	connection, err := net.Dial("tcp", "mirror.us.leaseweb.net:21")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	fmt.Println("INFO Welcome")
	code, message, err := readResponse(connection)
	if err != nil || code != "220" {
		log.Fatal(err)
	}

	fmt.Printf("*** Connected %s\n%s\n", code, message)
}

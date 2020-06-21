package main

import (
	"bufio"
	"fmt"
	"github.com/canro91/30DaysOfGo/Day33/go-ftp/client"
	"io"
	"log"
	"net"
	"net/textproto"
	"regexp"
	"strconv"
	"strings"
)

const (
	OpenData       = 150
	Ok             = 200
	Ready          = 220
	PasswordNeeded = 331
	LoggedIn       = 230
	PassiveMode    = 227
)

func checkResponse(conn net.Conn, responseCode int) (int, string, error) {
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	return tp.ReadResponse(responseCode)
}

func parsePassiveModeResponse(response string) (string, int) {
	re := regexp.MustCompile(`(\d+),(\d+),(\d+),(\d+),(\d+),(\d+)`)
	values := re.FindStringSubmatch(response)
	address := strings.Join(values[1:5], ".")

	a, _ := strconv.Atoi(values[5])
	b, _ := strconv.Atoi(values[6])
	port := a*256 + b
	return address, port
}

func ls(conn net.Conn) (string, error) {
	err := sendCommand(conn, "TYPE A")
	if err != nil {
		return "", err
	}
	_, _, err = checkResponse(conn, Ok)
	if err != nil {
		return "", err
	}

	err = sendCommand(conn, "PASV")
	if err != nil {
		return "", err
	}
	_, message, err := checkResponse(conn, PassiveMode)
	if err != nil {
		return "", err
	}

	address, port := parsePassiveModeResponse(message)
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	// Notice: LIST is send to the first socket
	err = sendCommand(conn, "LIST")
	if err != nil {
		return "", err
	}
	_, _, err = checkResponse(conn, OpenData)
	if err != nil {
		return "", err
	}

	message, err = readTextResponse(connection)
	if err != nil {
		return "", err
	}

	return message, nil
}

func sendCommand(conn net.Conn, command string) error {
	_, err := conn.Write([]byte(command + "\r\n"))
	return err
}

func readTextResponse(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	var line string
	for {
		nextLine, err := tp.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		line += "\n" + strings.TrimRight(nextLine, "\r\n")
	}

	return line, nil
}

func main() {
	client := client.NewClient("mirror.us.leaseweb.net", 21)
	connection, err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("*** Connected")

	err = client.AnonymousLogin()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("*** LoggedIn")

	fmt.Printf("*** Listing files:\n")
	message, err := ls(connection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", message)
}

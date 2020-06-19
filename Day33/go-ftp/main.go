package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strconv"
	"strings"
)

const (
	Ready          = 202
	PasswordNeeded = 331
	LoggedIn       = 230
)

func login(conn net.Conn, user, password string) error {
	err := sendCommand(conn, fmt.Sprintf("USER %s", user))
	if err != nil {
		return err
	}
	_, _, err = checkResponse(conn, PasswordNeeded)
	if err != nil {
		return err
	}

	err = sendCommand(conn, fmt.Sprintf("PASS %s", password))
	if err != nil {
		return err
	}
	_, _, err = checkResponse(conn, LoggedIn)
	if err != nil {
		return err
	}

	return nil
}

func sendCommand(conn net.Conn, command string) error {
	_, err := conn.Write([]byte(command + "\r\n"))
	return err
}

func checkResponse(conn net.Conn, responseCode int) (int, string, error) {
	code, message, err := readResponse(conn)
	if err != nil {
		return 0, "", err
	}
	if code != responseCode {
		return 0, "", errors.New(fmt.Sprintf("Expected Code %d, but was %d", responseCode, code))
	}
	return code, message, nil
}

// TODO: Take a look at tp.ReadResponse
func readResponse(conn net.Conn) (int, string, error) {
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	line, _ := tp.ReadLine()
	code, err := strconv.Atoi(line[:3])
	if err != nil {
		return 0, "", errors.New(fmt.Sprintf("Invalid response. Expected int, but was %s", line))
	}

	if line[3:4] == "-" {
		for {
			nextLine, err := tp.ReadLine()
			if err != nil {
				return 0, "", err
			}

			line += "\n" + strings.TrimRight(nextLine, "\r\n")
			if nextLine[:3] == string(code) && nextLine[3:4] != "-" {
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

	code, _, err := checkResponse(connection, Ready)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("*** Connected %d\n", code)

	usr := "anonymous"
	pwd := "anonymous@"
	err = login(connection, usr, pwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("*** LoggedIn %s\n", usr)
}

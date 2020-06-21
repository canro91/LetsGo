package client

import (
	"bufio"
	"fmt"
	"io"
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

type FtpClient struct {
	conn     net.Conn
	Hostname string
	Port     int
}

func NewClient(hostname string, port int) *FtpClient {
	c := &FtpClient{
		Hostname: hostname,
		Port:     port,
	}
	return c
}

func (c *FtpClient) Connect() error {
	address := fmt.Sprintf("%s:%d", c.Hostname, c.Port)
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	c.conn = connection

	_, _, err = c.checkResponse(Ready)
	if err != nil {
		return err
	}

	return nil
}

func (c *FtpClient) Close() {
	c.conn.Close()
}

func (c *FtpClient) AnonymousLogin() error {
	return c.Login("anonymous", "anonymous@")
}

func (c *FtpClient) Login(username, password string) error {
	_, err := c.sendCommand(fmt.Sprintf("USER %s", username), PasswordNeeded)
	if err != nil {
		return err
	}
	_, err = c.sendCommand(fmt.Sprintf("PASS %s", password), LoggedIn)
	if err != nil {
		return err
	}

	return nil
}

func (c *FtpClient) List() (string, error) {
	_, err := c.sendCommand("TYPE A", Ok)
	if err != nil {
		return "", err
	}
	pasvResponse, err := c.sendCommand("PASV", PassiveMode)
	if err != nil {
		return "", err
	}

	address, port := parsePassiveModeResponse(pasvResponse)
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return "", err
	}
	defer connection.Close()

	// Notice: LIST is send to the first socket
	_, err = c.sendCommand("LIST", OpenData)
	if err != nil {
		return "", err
	}

	message, err := readTextResponse(connection)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (c *FtpClient) sendCommand(command string, responseCode int) (string, error) {
	err := c.writeCommand(command)
	if err != nil {
		return "", err
	}
	_, message, err := c.checkResponse(responseCode)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (c *FtpClient) writeCommand(command string) error {
	_, err := c.conn.Write([]byte(command + "\r\n"))
	return err
}

func (c *FtpClient) checkResponse(responseCode int) (int, string, error) {
	reader := bufio.NewReader(c.conn)
	tp := textproto.NewReader(reader)
	return tp.ReadResponse(responseCode)
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

func parsePassiveModeResponse(response string) (string, int) {
	re := regexp.MustCompile(`(\d+),(\d+),(\d+),(\d+),(\d+),(\d+)`)
	values := re.FindStringSubmatch(response)
	address := strings.Join(values[1:5], ".")

	a, _ := strconv.Atoi(values[5])
	b, _ := strconv.Atoi(values[6])
	port := a*256 + b
	return address, port
}

package client

import (
	"net/textproto"
	"bufio"
	"fmt"
	"net"
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
	conn net.Conn
	Hostname string
	Port int
}

func NewClient(hostname string, port int) (*FtpClient) {
	c := &FtpClient{
		Hostname: hostname,
		Port: port,
	}
	return c
}

func (c *FtpClient) Connect() (net.Conn, error) {
	address := fmt.Sprintf("%s:%d", c.Hostname, c.Port)
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	c.conn = connection

	_, _, err = c.checkResponse(Ready)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (c *FtpClient) checkResponse(responseCode int) (int, string, error) {
	reader := bufio.NewReader(c.conn)
	tp := textproto.NewReader(reader)
	return tp.ReadResponse(responseCode)
}


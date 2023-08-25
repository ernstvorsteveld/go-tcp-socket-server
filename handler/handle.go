package handler

import (
	"bufio"
	"fmt"
	"go-tcp-socket-server/arguments"
	"net"
	"strings"
	"time"
)

type ConnectionDetails struct {
	args arguments.Arguments
	c    net.Conn
}

func NewConnectionDetails(c net.Conn, args arguments.Arguments) ConnectionDetails {
	return ConnectionDetails{
		args: args,
		c:    c,
	}
}

func (conn ConnectionDetails) Handle() {
	fmt.Printf("Server %s is serving %s\n", conn.args.Server, conn.c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(conn.c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(netData)

		message := strings.TrimSpace(string(netData))
		if message == "STOP" {
			break
		}

		result := message + " - at - " + time.Now().String() + "\n"
		conn.c.Write([]byte(string(result)))
	}
	conn.c.Close()
}

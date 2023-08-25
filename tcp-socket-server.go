package main

import (
	"flag"
	"fmt"
	"go-tcp-socket-server/arguments"
	"go-tcp-socket-server/handler"
	"math/rand"
	"net"
	"time"
)

func main() {
	s := flag.String("s", "localhost", "Server name or ip address")
	p := flag.String("p", "8080", "Port number")
	flag.Parse()
	arguments := arguments.HandleArguments([]string{*s, *p})

	l, err := net.Listen("tcp4", ":"+arguments.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		connDetails := handler.NewConnectionDetails(c, arguments)
		go connDetails.Handle()
	}
}

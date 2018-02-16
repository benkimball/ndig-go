package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	network = "tcp"
	port    = ":2525"
)

func main() {
	ln, _ := net.Listen(network, port)
	defer ln.Close()

	fmt.Printf("Listening on %s%s\n", network, port)

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	msg, err := rw.ReadString('\n')
	if err != nil {
		return
	}

	rw.WriteString(msg)
	rw.Flush()
}

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("server running on port", ln.Addr())
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		if conn != nil {
			fmt.Println("New Client Connected :", conn.LocalAddr())
		}

		io.WriteString(conn, fmt.Sprint("hello prathyusha..", "\n", "server Connected \n", time.Now()))

		conn.Close()
	}
}

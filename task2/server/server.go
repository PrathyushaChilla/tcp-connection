package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	x = make(map[string]*list.List)
)

var (
	openConnection = make(map[net.Conn]bool)
	newConnection  = make(chan net.Conn)
	deadConnection = make(chan net.Conn)
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	logFatal(err)

	defer ln.Close()

	go func() {

		for {
			conn, err := ln.Accept()
			logFatal(err)

			openConnection[conn] = true
			newConnection <- conn

		}
	}()

	for {
		select {
		case conn := <-newConnection:
			go Store(conn)
		case conn := <-deadConnection:
			for item := range openConnection {
				if item == conn {
					break
				}
			}
		}
	}

}

func Store(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		usernameAndMessage, err := reader.ReadString('\n')
		logFatal(err)

		//fmt.Print(usernameAndMessage)

		parts := strings.Split(usernameAndMessage, ":-")
		username := parts[0]
		message := parts[1]

		//fmt.Println(parts[0])

		if _, ok := x[username]; ok {
			//fmt.Println("User exits in Connection")
			x[username].PushBack(message)
			fmt.Println("Messages From ", username, ":-")
			for e := x[username].Front(); e != nil; e = e.Next() {
				fmt.Print(e.Value)
			}
		} else {
			x[username] = list.New()
			x[username].PushBack(message)
			fmt.Println("Messages From ", username, ":-")
			for e := x[username].Front(); e != nil; e = e.Next() {
				fmt.Print(e.Value)
			}
		}
	}

	deadConnection <- conn

}

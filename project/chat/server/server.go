package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	out  chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// timeout
var timeout time.Duration

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.out <- msg

			}

		case cli := <-entering:
			clients[cli] = true

			cli.out <- "Present:"
			for c := range clients {
				cli.out <- c.name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)

		}
	}
}

func handleConn(conn net.Conn) {
	in := make(chan string, 10)
	out := make(chan string)

	// Cleint write
	go clientWrite(conn, out)

	// Client read
	go clientRead(conn, in)

	// start timer
	nameTimer := time.NewTimer(timeout)

	out <- "Enter you name"

	var who string
	select {
	case name := <-in:
		who = name
	case <-nameTimer.C:
		conn.Close()
		return

	}

	// user
	cli := client{out, who}
	out <- "you are " + who

	// msg
	messages <- who + "has arrived"

	entering <- cli

	idle := time.NewTimer(timeout)

loop:
	for {
		select {
		case msg := <-in:
			messages <- who + ":" + msg
			idle.Reset(timeout)
		case <-idle.C:
			conn.Close()
			break loop

		}
	}
	leaving <- cli

	messages <- who + "has left"
	conn.Close()
}

func clientRead(conn net.Conn, in chan<- string) {
	input := bufio.NewScanner(conn)

	for input.Scan() {
		in <- input.Text()
	}
}

func clientWrite(conn net.Conn, out chan string) {
	for msg := range out {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	var (
		listener net.Listener
		conn     net.Conn
		err      error
	)

	// start  tcp

	listener, err = net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	// goroutine do

	go broadcaster()

	// accept

	for {
		if conn, err = listener.Accept(); err != nil {
			log.Print(err)
			continue
		}

		// 处理连接
		go handleConn(conn)

	}

}

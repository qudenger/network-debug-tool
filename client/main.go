package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func send(conn net.Conn) {
	for i := 0; i < 5; i++ {
		words := strconv.Itoa(i) + "This is a test for long conn"
		conn.Write([]byte(words))
		time.Sleep(1 * time.Second)

	}
	fmt.Println("send over")
	defer conn.Close()
}

func main() {
	server := "localhost:12345"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	send(conn)

}

package main

import (
	"log"
	"net"
)

// Client holds info about connection
type Client struct {
	conn   net.Conn
	Server *server
}

// TCP server
type server struct {
	address                  string // Address to open connection: localhost:9999
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message []byte)
}

// Read client data from channel
func (c *Client) listen() {

	buf := make([]byte, 1024)

	for {
		lenght, err := c.conn.Read(buf)
		// if checkError(err, "Connection") == false {
		// 	conn.Close()
		// 	break
		// }
		if err != nil {
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		//fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		//reciveStr := string(buf[0:lenght])
		c.Server.onNewMessage(c, buf[0:lenght])
	}

	// reader := bufio.NewReader(c.conn)

	// for {

	// 	message, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		c.conn.Close()
	// 		c.Server.onClientConnectionClosed(c, err)
	// 		return
	// 	}
	// 	c.Server.onNewMessage(c, message)
	// }
}

// Send text message to client
func (c *Client) Send(message []byte) error {
	_, err := c.conn.Write(message)
	return err
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called right after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, message []byte)) {
	s.onNewMessage = callback
}

// Start network server
func (s *server) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("==============================000000000000")
			log.Println(err)
			// if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
			// 	log.Printf("NOTICE: temporary Accept() failure - %s", err.Error())
			// 	runtime.Gosched()
			// 	continue
			// }
			// // theres no direct way to detect this error because it is not exposed
			// if !strings.Contains(err.Error(), "use of closed network connection") {
			// 	log.Printf("ERROR: listener.Accept() - %s", err.Error())
			// }
			// break
		}
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
		s.onNewClientCallback(client)
	}
}

// Creates new tcp server instance
func New(address string) *server {
	log.Println("Creating server with address", address)
	server := &server{
		address: address,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message []byte) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}

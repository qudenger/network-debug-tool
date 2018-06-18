package main

import (
	"log"
	"net"
	"strings"
)

var timeout = 600 // 超时60秒，自动断开连接

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
		length, err := c.conn.Read(buf)
		if err != nil {
			// 这种属于正常断开连接的情况：发送了EOF
			LogErr(c.conn.RemoteAddr().String(), " connection error: ", err)
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		if length > 0 {
			buf[length] = 0
		}
		log.Println("-----=======" + c.conn.RemoteAddr().String())
		Data := (buf[:length])
		messnager := make(chan byte)
		//心跳计时 没有加入websocket代理的定时发送心跳机制

		if !strings.HasPrefix(c.conn.RemoteAddr().String(), "127.0.0.1") {
			log.Println("不是127的")
			go HeartBeating(c.conn, messnager, timeout)
			//检测每次Client是否有数据传来 // todo: 排除掉对127.0.0.1的连接
			go GravelChannel(c, Data, messnager)
		} else {
			c.Server.onNewMessage(c, buf[0:length])
		}

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
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
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

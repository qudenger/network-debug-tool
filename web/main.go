package main

import (
	"io"
	"log"
	"net"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

var (
	tcpAddress string = "127.0.0.1:12345"
)

func copyWorker(dst io.Writer, src io.Reader, doneCh chan<- bool) {
	io.Copy(dst, src)
	doneCh <- true
}

func hello(c echo.Context) error {

	websocket.Handler(func(ws *websocket.Conn) {
		conn, err := net.Dial("tcp", tcpAddress)
		if err != nil {
			log.Printf("[ERROR] %v \n", err)
			return
		}

		for {
			ws.PayloadType = websocket.BinaryFrame
			doneCh := make(chan bool)
			go copyWorker(conn, ws, doneCh)
			go copyWorker(ws, conn, doneCh)
			<-doneCh
			conn.Close()
			ws.Close()
			<-doneCh
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.Any("/ws", hello)
	e.Logger.Fatal(e.Start(":1234"))
}

package main

import (
	"net"
	"time"
)

//HeartBeating, determine if client send a message within set time by GravelChannel
// 心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息

func HeartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case _ = <-readerChannel:
		Log(conn.RemoteAddr().String(), "get message, keeping heartbeating...")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * 5):
		// todo:应该排除掉对 127.0.0.1的连接对象
		conn.Close() // 会触发OnClientConnectionClosed事件
	}

}

func GravelChannel(c *Client, n []byte, mess chan byte) {
	for _, v := range n {
		mess <- v
	}
	c.Server.onNewMessage(c, n)
	close(mess)
}

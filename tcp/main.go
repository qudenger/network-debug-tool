package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
	msgProto "github.com/qudenger/network-debug-tool/tcp/proto"
)

// 广播所有的
func echoHandler(conns *map[string]net.Conn, msg string) {
	for key, value := range *conns {
		fmt.Println("connection is connected from ...", key)
		_, err := value.Write([]byte(msg))
		if err != nil {
			fmt.Println(err.Error())
			delete(*conns, key)
		}
	}
}

func doGetClientList(receiver string, conns *map[string]net.Conn) error {
	fmt.Println("doGetClientList")

	clients := []string{}
	for key, _ := range *conns {
		if !strings.HasPrefix(key, "127.0.0.1") {
			clients = append(clients, key)
		}
	}

	clientsArray, err := json.Marshal(clients)
	if err != nil {
		clientsArray = nil
	}

	msg := &msgProto.Message{ // 使用辅助函数设置域的值
		Cmd:      "GetClientList",
		Sender:   "tcp-server",
		Receiver: receiver,
		Body:     clientsArray,
	}

	result, _ := proto.Marshal(msg)

	_, err = (*conns)[receiver].Write(result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func doAddNewClient(conns *map[string]net.Conn, clientAddr string) {
	for key, value := range *conns {
		// 如果是本地web下的客户端：
		if strings.HasPrefix(key, "127.0.0.1") {

			msg := &msgProto.Message{ // 使用辅助函数设置域的值
				Cmd:      "AddNewClient",
				Sender:   "tcp-server",
				Receiver: key,
				Body:     []byte(clientAddr),
			}
			rsp, _ := proto.Marshal(msg)
			_, err := value.Write(rsp)
			if err != nil {
				fmt.Println(err.Error())
				//delete(*conns, key)
			}
		}
	}
}

func doClientClose(conns *map[string]net.Conn, clientAddr string) {
	for key, value := range *conns {
		// 如果是本地web下的客户端：
		if strings.HasPrefix(key, "127.0.0.1") {
			msg := &msgProto.Message{ // 使用辅助函数设置域的值
				Cmd:      "ClientClose",
				Sender:   "tcp-server",
				Receiver: key,
				Body:     []byte(clientAddr),
			}
			rsp, _ := proto.Marshal(msg)

			_, err := value.Write([]byte(rsp))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	delete(*conns, clientAddr)
}

func doReceiveMsg(conns *map[string]net.Conn, from string, body []byte) error {
	// todo: 判断这些参数不能为空的情况
	fmt.Println("doReceiveMsg")
	fmt.Println("from:" + from)

	for key, value := range *conns {
		// 如果是本地web下的客户端：
		if strings.HasPrefix(key, "127.0.0.1") {
			msg := &msgProto.Message{ // 使用辅助函数设置域的值
				Cmd:      "ReceiveMsg",
				Sender:   from,
				Receiver: key,
				Body:     body,
			}
			rsp, _ := proto.Marshal(msg)
			_, err := value.Write([]byte(rsp))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	return nil
}

func doSendMsgToClient(conns *map[string]net.Conn, target string, body []byte) error {
	// todo: 判断这些参数不能为空的情况
	fmt.Println("target:" + target)

	_, err := (*conns)[target].Write(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func main() {

	server := New("0.0.0.0:12345")

	conns := make(map[string]net.Conn)
	// 如果是127.0.0.1的客户端就不要注册了，属于tcp server的代理；
	server.OnNewClient(func(c *Client) {
		conns[c.conn.RemoteAddr().String()] = c.conn
		if !strings.HasPrefix(c.conn.RemoteAddr().String(), "127.0.0.1") {
			doAddNewClient(&conns, c.conn.RemoteAddr().String())
		}
	})
	server.OnNewMessage(func(c *Client, message []byte) {

		// 只要是来自127.0.0.1的连接都是websocket端，那么按照protobuf进行解析
		fromProxy := strings.HasPrefix(c.conn.RemoteAddr().String(), "127.0.0.1")

		if fromProxy {
			fmt.Println("来自本机的客户端")

			req := &msgProto.Message{}
			parserErr := proto.Unmarshal(message, req)
			if parserErr != nil {
				fmt.Println(parserErr.Error())
				//return
			}
			fmt.Println(req.Cmd)
			switch req.Cmd {
			case "GetClientList":
				doGetClientList(c.conn.RemoteAddr().String(), &conns)
				break
			case "SendMsgToClient":
				doSendMsgToClient(&conns, req.Receiver, req.Body)
				break
			case "ReceiveMsg":
				doReceiveMsg(&conns, req.Sender, req.Body)
				break
			}
		} else {
			from := c.conn.RemoteAddr().String()
			doReceiveMsg(&conns, from, message)

		}

		// if strings.HasPrefix(c.conn.RemoteAddr().String(), "127.0.0.1") {
		// 	fmt.Println("来自本机的客户端")
		// 	//都该有:{"cmd":"***"}
		// 	var request map[string]interface{}
		// 	err := json.Unmarshal([]byte(message), &request)
		// 	if err == nil {
		// 		if request["cmd"].(string) == "GetClientList" {
		// 			fmt.Println("0000")
		// 			c.Send(GetClientList(&conns))
		// 		} else if request["cmd"].(string) == "SendMsgToClient" {
		// 			fmt.Println("1111")
		// 			target := request["target"].(string)
		// 			msg := request["msg"].(string)
		// 			SendMsgToClient(&conns, target, msg)
		// 		}
		// 	}
		// } else {
		// 	from := c.conn.RemoteAddr().String()
		// 	ReceiveMsg(&conns, from, message)
		// }
		// fmt.Println(message)

		fmt.Println(message)
	})
	server.OnClientConnectionClosed(func(c *Client, err error) {
		// connection with client lost
		//c.Send("closed")
		from := c.conn.RemoteAddr().String()
		doClientClose(&conns, from)

	})

	server.Listen()
}

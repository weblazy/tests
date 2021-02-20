package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	Start(":7123")
}

func Start(tcpAddrStr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpAddrStr)
	if err != nil {
		log.Printf("Resolve tcp addr failed: %v\n", err)
		return
	}

	// 向服务器拨号
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Dial to server failed: %v\n", err)
		return
	}

	// 向服务器发消息
	go SendMsg(conn)

	// 接收来自服务器端的广播消息
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}

		fmt.Println(string(buf[0:length]))
	}
}

// 向服务器端发消息
func SendMsg(conn net.Conn) {
	username := conn.LocalAddr().String()
	for {
		var input string

		// 接收输入消息，放到input变量中
		fmt.Scanln(&input)

		if input == "/q" || input == "/quit" {
			fmt.Println("Byebye ...")
			conn.Close()
			os.Exit(0)
		}

		// 只处理有内容的消息
		if len(input) > 0 {
			msg := username + " say:" + input
			_, err := conn.Write([]byte(msg))
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}

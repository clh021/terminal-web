package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func tcpSend(conn net.Conn, content string) {
	_, err := conn.Write([]byte(content + "\r\n"))
	if err != nil {
		return
	}
	fmt.Println("[TCP]Send:", content)
}
func tcpRecv(conn net.Conn) {
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	}
	fmt.Println(string(buf[:n]))
}
func tcpSender(sendto string, port int) {
	// 建立连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", sendto, port))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer conn.Close()

	// 自动测试
	for i := 0; i < 5; i++ {
		tcpSend(conn, time.Now().Local().GoString())
		time.Sleep(time.Second * 2)
		tcpRecv(conn)
	}

	// 手动测试
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("[TCP]Input something to send(q to quite):")
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		tcpSend(conn, inputInfo)
		tcpRecv(conn)
	}
}

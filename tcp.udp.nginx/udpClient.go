package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func udpSend(conn net.Conn, content string) {
	_, err := conn.Write([]byte(content))
	if err != nil {
		fmt.Println("[UDP]SendError: ", err)
		return
	}
	fmt.Println("[UDP]Send:", content)
}
func udpRecv(conn *net.UDPConn) {
	data := make([]byte, 4096)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("[UDP]ReadError: ", err)
	}
	fmt.Printf("[UDP]read %d from <%s> recv:%v \n", n, remoteAddr, data[:n])
}
func udpSender(sendto string, port int, inputMode bool) {
	// 建立连接
	ip := net.ParseIP(sendto)
	if ip == nil {
		addr, err := net.ResolveIPAddr("ip", sendto)
		if err != nil {
			fmt.Println("[UDP]ParseIPError: ", err)
		}
		fmt.Println("[UDP]ParseIP: ", sendto, " => ", addr.String())
		ip = net.ParseIP(addr.String())
	}
	log.Println("udp sendto ip:", ip)
	// srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: port}
	conn, err := net.DialUDP("udp", nil, dstAddr)
	if err != nil {
		fmt.Println("[UDP]ConnectError: ", err)
		return
	}
	defer conn.Close()

	// 自动测试
	for i := 0; i < 5; i++ {
		udpSend(conn, time.Now().Local().GoString())
		time.Sleep(time.Second * 2)
		udpRecv(conn)
	}

	if inputMode {
		// 手动测试
		inputReader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("[UDP]Input something to send(q to quite):")
			input, _ := inputReader.ReadString('\n')
			inputInfo := strings.Trim(input, "\r\n")
			if strings.ToUpper(inputInfo) == "Q" {
				return
			}
			udpSend(conn, inputInfo)
			udpRecv(conn)
		}
	}
}

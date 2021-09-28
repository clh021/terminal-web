package main

import (
	"fmt"
	"net"
	"os"
)

func udpListen(conn_host string, conn_port int) {
	addr := net.UDPAddr{
		Port: conn_port,
		IP:   net.ParseIP(conn_host),
	}
	lis, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	defer lis.Close()
	fmt.Printf("[UDP]Listen: <%s> \n", lis.LocalAddr().String())

	for {
		var data [4096]byte
		n, remoteAddr, err := lis.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("[UDP]ReadError: ", err)
			continue
		}
		fmt.Println("[UDP]Recv: bytes=", n, " from=", remoteAddr, " data:", string(data[:n]))
		_, err = lis.WriteToUDP([]byte("hello"), remoteAddr) // 发送数据
		if err != nil {
			fmt.Println("[UDP]WriteError: ", err)
			continue
		}
	}
}
func tcpListen(conn_host string, conn_port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conn_host, conn_port))
	if err != nil {
		fmt.Println("[TCP]Errorlis:", err.Error())
		os.Exit(1)
	}
	defer lis.Close()
	fmt.Println("[TCP]Listen: ", lis.Addr().String())
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("[TCP]Erroraccepting: ", err.Error())
			// os.Exit(1)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("[TCP]Error reading:", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("[TCP]Recv: bytes=", n, " data:", recv)
		conn.Write([]byte(fmt.Sprintf("[TCP]Message received. data: %v", recv)))
	}
	conn.Close()
}

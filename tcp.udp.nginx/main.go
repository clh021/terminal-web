package main

import (
	"flag"
	"fmt"
)

func main() {
	var stype = flag.String("stype", "server", "[server]tcp&udp server listen.\r\n[client]tcp&udp sender.")
	var sendto = flag.String("sendto", "127.0.0.1", "host addr tcp udp send to.")
	var port = flag.Int("port", 7990, "tcp&udp server listen port.")
	flag.Parse()
	switch *stype {
	case "server":
		fmt.Println("Start Type: server")
		go tcpListen("tcp", "0.0.0.0", *port)
		udpListen("0.0.0.0", *port)
		return
	case "client":
		fmt.Println("Start Type: client")
		udpSender(*sendto, *port)
		tcpSender(*sendto, *port)
		return
	default:
		fmt.Println("Not support Server Type:", *stype)
		return
	}
}

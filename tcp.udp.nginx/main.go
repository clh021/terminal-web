package main

import (
	"flag"
	"fmt"
)

func main() {
	var stype = flag.String("stype", "server", "[server]tcp&udp server listen.\r\n[client]tcp&udp sender.")
	var port = flag.Int("port", 7990, "tcp&udp server listen port.")
	var toHost = flag.String("toHost", "127.0.0.1", "host addr tcp udp send to.")
	var toPort = flag.Int("toPort", 7990, "port addr tcp udp send to.")
	var inputMode = flag.Bool("inputMode", false, "enable to test by input string for client.")
	flag.Parse()
	switch *stype {
	case "server":
		fmt.Println("Start Type: server")
		fmt.Println("Server Port: ", *port)
		go tcpListen("0.0.0.0", *port)
		udpListen("0.0.0.0", *port)
		return
	case "client":
		fmt.Println("Start Type: client")
		fmt.Println("send Host: ", *toHost)
		fmt.Println("send Port: ", *toPort)
		udpSender(*toHost, *toPort, *inputMode)
		tcpSender(*toHost, *toPort, *inputMode)
		return
	default:
		fmt.Println("Not support Server Type:", *stype)
		return
	}
}

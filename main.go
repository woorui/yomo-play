package main

import (
	"fmt"
	"log"
	"net"

	"github.com/woorui/yomo-play/source"
	"github.com/yomorun/yomo"
)

var (
	zipperaddr = "127.0.0.1:9000"
	tcpAddr    = "0.0.0.0:8080"
)

func main() {
	s := yomo.NewSource("avg", zipperaddr)

	if err := s.Connect(); err != nil {
		log.Fatalln("source connect error:", err)
	}

	listener, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Fatalln("listen error:", err)
	}

	fmt.Println("SERVER UP")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept error:", err)
		}

		go source.PipeToSource(conn.RemoteAddr().String(), conn, s)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/woorui/yomo-play/source"
	"github.com/yomorun/yomo"
)

var (
	broker     = ""
	zipperaddr = "127.0.0.1:9000"
	tcpAddr    = "0.0.0.0:8080"
)

func main() {
	flag.StringVar(&broker, "broker", "yomo", "yomo or http")
	flag.Parse()

	if broker == "yomo" {
		runYomoBroker()
	} else {
		runHttpBroker()
	}
}

func runHttpBroker() {
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

		// http srv
		go source.PostToHttpSrv(conn.RemoteAddr().String(), conn)
	}
}

func runYomoBroker() {
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

		// yomo source
		go source.PipeToSource(conn.RemoteAddr().String(), conn, s)
	}
}

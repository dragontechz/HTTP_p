package main

import (
	"fmt"
	"log"
	"net"
)

var buff []byte = make([]byte, 1024)

type proxy struct {
	port string
}

func main() {
	proxy_server := proxy{"9999"}

	proxy_server.listen()
}

func (p *proxy) listen() {
	lsiten, err := net.Listen("tcp", ":"+p.port)
	if err != nil {
		log.Println("ERROR LISTENING: ", err)
	}
	fmt.Println("LISTENING ON PORT ", p.port)

	for {
		conn, err := lsiten.Accept()
		if err != nil {
			log.Println("ERROR ACCEPTING: ", err)
		}
		conn.Read(buff)
		conn.Write(buff)
	}

}

func handle_conn(conn net.Conn) {
	conn.Read(buff)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}

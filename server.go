package main

import (
	"fmt"
	"log"
	"net"
)

type proxy struct {
	port string
}

var token string = "1234"

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
		go handle_conn(conn)

		continue
	}

}

func handle_conn(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("ERROR RECVING: ", err)
	}
	data := string(buffer[:n])
	log.Println("RECVED :", data, "of length", n, "bytes")
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}

func handle_traffic(data int) {

}

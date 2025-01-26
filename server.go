package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port string = ":8080"
var buff []byte = make([]byte, 1024)

type sshProxy interface {
	Write([]byte) (int, error)
	Read([]byte) (int, error)
	Close() error
}

type proxy struct {
	adrr, dest_addr string
}

func main() {
	port := flag.String("p", "8080", "port to run the server on")
	flag.Parse()
	p := proxy{":" + *port, ":22"}

	p.HTTP_proxy()
}

func (p *proxy) HTTP_proxy() {
	sock, err := net.Listen("tcp", p.adrr)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	fmt.Println("server listening on port ", p.adrr)

	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Println("ERROR: ", err)
		}
		log.Println("new connection established by:", conn.RemoteAddr().String())

		conn.Read(buff)
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

	}

}

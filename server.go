package main

import (
	"fmt"
	"io"
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
		log.Println("handling new incomming connection from :", conn.RemoteAddr().String())

		conn_remote, err := net.Dial("tcp", ":8888")
		if err != nil {
			log.Println("cannot dial to remote port")
		}
		go io.Copy(conn_remote, conn)

		go io.Copy(conn, conn_remote)
	}

}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()
	conn_remote, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("cannot dial to remote port")
	}
	buffer := make([]byte, 1024)
	n, err := clientConn.Read(buffer)
	if err != nil {
		log.Println("couldnot read from remote client")
	}
	buf := make([]byte, 1024)
	conn_remote.Write(buffer[:n])
	conn_remote.Read(buf)
	conn_remote.Write(buffer[:n])

	n, err = conn_remote.Read(buffer)
	if err != nil {
		log.Println("cannot dial to remote port")
	}
	clientConn.Write(buffer[:n])
	conn_remote.Close()
}

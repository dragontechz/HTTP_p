package main

import (
	"fmt"
	"log"
	"net"
)

type http_proxy struct {
	port string
}
type client struct {
	conn net.Conn
}

func main() {

	proxy := http_proxy{"7777"}

	proxy.run()

}

func (p *http_proxy) run() {
	listener, err := net.Listen("tcp", ":"+p.port)
	fmt.Println("server listening on port: ", p.port)
	if err != nil {
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
		}

		client := client{conn}

		log.Println("client created \nHandling client")
		go client.handle()
		continue
	}

}

func (c *client) handle() {
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		log.Println("ERROR RECEIVING: ", err)
	}
	data := string(buffer[:n])
	log.Println("RECVED number of byte: ", n, "\nwith value : ", data)

	_, err = c.conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\n"))
	if err != nil {
		log.Println("ERROR COULDN'T WRITE TO CLIENT: ", err)
	}
}

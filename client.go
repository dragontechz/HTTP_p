package main

import (
	"log"
	"net"
)

type client struct {
	listening_port, dst_port string
}

func main() {
	client_server := client{":9090", ":8888"}

	client_server.start()
}
func (c *client) start() {
	listener, err := net.Listen("tcp", c.listening_port)
	if err != nil {
	}
	log.Println("LISTENING ON PORT ", c.listening_port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("couldn't accept incomimg conn")
		}
		go c.handle(conn)
	}
}

func (c *client) handle(client net.Conn) {
	for {
		log.Println("handling client")
		buff := make([]byte, 1024*8)

		n, err := client.Read(buff)
		if err != nil {
		}
		if n < 10 {

		}
		c.send_handle_req(buff, n, client)
		log.Println("waiting for handle")
	}

}

func (c *client) send_handle_req(buff []byte, n int, client net.Conn) {
	server, err := net.Dial("tcp", c.dst_port)
	if err != nil {
		log.Println("ERROR IN CONNECTING TO REMOTE HOST")
	}
	data := string(buff[:n])
	log.Println("DATA TO SEND : ", data)

	_, err = server.Write(buff[:n])
	if err != nil {
		log.Println("CANN'T WRITE TO SERVER : ", err)
	}
	log.Println("SUCCESSFULLY FORWARD DATA TO SERVER")
	buffer := make([]byte, 1024)
	for {
		n, err = server.Read(buffer)

		if n < 1 {
			continue
		}
		if err != nil {
			log.Println(" COULDN'T RECV FROM SERVER : ", err)
		}

		data = string(buffer[:n])
		log.Println("DATA RECVED : ", data)

		client.Write(buffer[:n])
		log.Println("SUCCESFULLY FORWARD DATA TO CLIENT")
		break
	}
}

func (c *client) handle_recv(client, server_conn net.Conn) {
	for {
		buffer := make([]byte, 1024*8)
		n, err := server_conn.Read(buffer)
		if n < 1 {
			continue
		}
		if err != nil {
			log.Println(" COULDN'T RECV FROM SERVER : ", err)
		}

		data := string(buffer[:n])
		log.Println("DATA RECVED : ", data)

		client.Write(buffer[:n])
		log.Println("SUCCESFULLY FORWARD DATA TO CLIENT")
		break
	}
}

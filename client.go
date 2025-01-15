package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

//	"time"

var buff []byte = make([]byte, 1024/2)

var mu sync.Mutex
var writeOrder int

type client_presistent_conn struct {
	dst string
}

func main() {
	client := client_presistent_conn{"170.205.31.126:9999"}
	client.send_and_recv()
	//	}
}

var i int = 1

func (c *client_presistent_conn) send_and_recv() {
	for {
		conn, err := net.Dial("tcp", c.dst)
		defer conn.Close()
		if err != nil {
			log.Println("ERROR DIALING: ", err)
		}
		_, err = conn.Write([]byte("1234"))

		if err != nil {
			log.Println("ERROR : ", err)
		}
		conn.Read(buff)
		conn.Write([]byte("HTTP NUMBER :" + fmt.Sprint(i)))
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("ERROR RECEIVING: ", err)
		}
		data := string(buff[:n])
		log.Println("received from server: ", data)
		i++
		conn.Close()
	}
}

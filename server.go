package main

import (
	"log"
	"net"
)

func main() {

	addr := net.UDPAddr{
		Port: 8081,
		IP:   net.ParseIP("localhost"),
	}

	for {
		log.Println("listening on port 8081")
		conn, err := net.ListenUDP("udp", &addr)

		if err != nil {

		}

		defer conn.Close()

		for {
			buff := make([]byte, 1024)
			n, client_addr, err := conn.ReadFromUDP(buff)
			if err != nil {

			}

			log.Printf("message recv from %s : %s\n", client_addr.String(), string(buff[:n]))
			go send(conn, client_addr)
		}
	}
}
func send(conn *net.UDPConn, client_addr *net.UDPAddr) {
	messageReponse := []byte("xbyte 200 ok")
	_, err := conn.WriteToUDP(messageReponse, client_addr)
	if err != nil {
		log.Println("Erreur lors de l'envoi de la réponse :", err)
	}
	log.Println("SUCCESSFULLY SENT RESPONSE TO CLIENT")
}

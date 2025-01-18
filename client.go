package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Définir l'adresse du serveur
	serverAddr := net.UDPAddr{
		Port: 8081,                     // Assurez-vous que le port correspond à celui du serveur
		IP:   net.ParseIP("170.205.31.126"), // Remplace par l'adresse IP du serveur si nécessaire
	}

	// Créer le socket UDP
	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		fmt.Println("Erreur lors de la connexion :", err)
		os.Exit(1)
	}
	defer conn.Close()

	go recv(conn)

	// Message à envoyer
	message := []byte("GET {?query} HTTP/1.1//end")
	for {
		// Envoyer le message au serveur
		_, err = conn.Write(message)
		if err != nil {
			fmt.Println("Erreur lors de l'envoi du message :", err)
			return
		}
		fmt.Println("Message envoyé au serveur :", string(message))

		// Préparer un buffer pour la réponse
	}
}

func recv(conn *net.UDPConn) {
	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse :", err)
			return
		}

		fmt.Printf("Réponse du serveur : %s\n", string(buffer[:n]))
	}
}

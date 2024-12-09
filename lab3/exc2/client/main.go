package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

var clientKey string

func main() {
	serverAddr, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	if !authenticate(conn) {
		fmt.Println("Authentication failed. Exiting...")
		return
	}

	playGame(conn)
}

func authenticate(conn *net.TCPConn) bool {
	fmt.Print(read(conn))
	var username string
	fmt.Scanf("%s", &username)
	send(conn, username)

	fmt.Print(read(conn))
	var password string
	fmt.Scanf("%s", &password)
	send(conn, password)

	response := read(conn)
	fmt.Println(response)

	if strings.Contains(response, "AUTH:OK") {
		parts := strings.Split(response, ":")
		clientKey = parts[len(parts)-1]
		return true
	}
	return false
}

func playGame(conn *net.TCPConn) {
	var playing bool = true
	for playing {
		var input string
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%s", &input)
		send(conn, input)

		response := read(conn)
		fmt.Println(response)

		if strings.Contains(response, "Correct") {
			fmt.Print("Play again? (yes/no): ")
			fmt.Scanf("%s", &input)

			switch strings.ToLower(input) {
			case "yes":
				send(conn, "yes")
			case "no":
				send(conn, "no")
				fmt.Println(read(conn))
				playing = false
			default:
				fmt.Println("Invalid input. Exiting...")
				send(conn, "no")
				playing = false
			}
		}
	}
}

func send(conn *net.TCPConn, s string) {
	message := fmt.Sprintf("%s_%s", clientKey, s)
	_, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		panic(err)
	}
}

func read(conn *net.TCPConn) string {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(buffer[:n]))
}

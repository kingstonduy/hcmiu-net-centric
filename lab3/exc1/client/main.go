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

	var playing bool = true
	for playing {
		var input string
		// Get user input for the guess
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%s", &input)

		// Send the guess to the server
		send(conn, input)

		response := read(conn)
		fmt.Println(response)
		if strings.Contains(strings.ToLower(response), "correct") {
			fmt.Scanf("%s", &input)
			switch strings.ToLower(input) {
			case "yes":
				send(conn, "yes")
			case "no":
				send(conn, "no")
				response := read(conn)
				fmt.Println(response)
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
	_, err := conn.Write([]byte(s + "\n"))
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
	s := string(buffer[:n])
	return s
}

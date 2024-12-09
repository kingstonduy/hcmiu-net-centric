package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is running...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleGame(conn)
	}
}

func handleGame(conn net.Conn) {
	defer conn.Close()
	for {
		target := rand.Intn(100) + 1
		fmt.Println("New game started with target number:", target)

		for {
			s := read(conn)
			guess, err := strconv.Atoi(s)
			if err != nil {
				conn.Write([]byte("Invalid input. Please enter a number between 1 and 100, or -1 to quit.\n"))
				continue
			}

			if guess == -1 {
				conn.Write([]byte("Game stopped. Connection will now close.\n"))
				fmt.Println("Client ended the game.")
				return
			}

			if guess < target {
				send(conn, "Too low! Try again.\n")
			} else if guess > target {
				send(conn, "Too high! Try again.\n")
			} else {
				send(conn, "Correct! You guessed the number.\nDo you want to play again? (yes/no): ")
				if strings.ToLower(read(conn)) != "yes" {
					send(conn, "Goodbye! Thanks for playing.")
					return
				} else {
					target = rand.Intn(100) + 1 // Random number between 1 and 100
					fmt.Println("New game started with target number:", target)
				}
			}
		}
	}
}

func send(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s + "\n"))
	if err != nil {
		panic(err)
	}
}

func read(conn net.Conn) string {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(buffer[:n]))
	return s
}

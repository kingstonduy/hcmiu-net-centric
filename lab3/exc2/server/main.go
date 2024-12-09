package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	HOST     = "localhost"
	PORT     = "8080"
	TYPE     = "tcp"
	USERNAME = "user"
	PASSWORD = "password"
)

var clientKeys sync.Map

func main() {
	fmt.Println("username:", USERNAME)
	fmt.Println("password:", PASSWORD)
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
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	if !authenticate(conn) {
		send(conn, "ERROR:Invalid credentials. Disconnecting.")
		return
	}

	clientKey := generateUniqueKey()
	clientKeys.Store(conn, clientKey)

	send(conn, fmt.Sprintf("%s_AUTH:OK. Welcome! Your unique key is: %s", clientKey, clientKey))
	playGame(conn, clientKey)
}

func authenticate(conn net.Conn) bool {
	send(conn, "AUTH:Enter username:")
	username := read(conn, "")

	send(conn, "AUTH:Enter password:")
	password := read(conn, "")

	return username == USERNAME && password == PASSWORD
}

func playGame(conn net.Conn, clientKey string) {
	for {
		target := rand.Intn(100) + 1
		fmt.Println("New game started with target number:", target)

		for {
			s := read(conn, clientKey)
			guess, err := strconv.Atoi(s)
			if err != nil {
				send(conn, fmt.Sprintf("%s_ERROR:Invalid input. Enter a number between 1 and 100, or -1 to quit.", clientKey))
				continue
			}

			if guess == -1 {
				send(conn, fmt.Sprintf("%s_MSG:Game stopped. Goodbye!", clientKey))
				fmt.Println("Client ended the game.")
				return
			}

			if guess < target {
				send(conn, fmt.Sprintf("%s_RESULT:Too low! Try again.", clientKey))
			} else if guess > target {
				send(conn, fmt.Sprintf("%s_RESULT:Too high! Try again.", clientKey))
			} else {
				send(conn, fmt.Sprintf("%s_RESULT:Correct! Play again? (yes/no):", clientKey))
				if strings.ToLower(read(conn, clientKey)) != "yes" {
					send(conn, fmt.Sprintf("%s_MSG:Thanks for playing! Goodbye.", clientKey))
					return
				}
				break
			}
		}
	}
}

func generateUniqueKey() string {
	return strconv.Itoa(rand.Intn(1000) + 1)
}

func send(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s + "\n"))
	if err != nil {
		panic(err)
	}
}

func read(conn net.Conn, clientKey string) string {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	msg := strings.TrimSpace(string(buffer[:n]))

	if clientKey != "" && !strings.HasPrefix(msg, clientKey+"_") {
		send(conn, fmt.Sprintf("%s_ERROR:Invalid key prefix. Disconnecting.", clientKey))
		panic("Invalid key prefix")
	}

	return strings.TrimPrefix(msg, clientKey+"_")
}

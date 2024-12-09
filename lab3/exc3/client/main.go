package main

import (
	"bufio"
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

	downloadFile(conn)
	fmt.Println("File downloaded successfully.")
}

func authenticate(conn *net.TCPConn) bool {
	fmt.Print(read(conn))
	var username string
	fmt.Scanln(&username)
	send(conn, username)

	fmt.Print(read(conn))
	var password string
	fmt.Scanln(&password)
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

func downloadFile(conn *net.TCPConn) {
	send(conn, "DOWNLOAD:words.txt")

	// Create a new file to store the downloaded content
	file, err := os.Create("downloaded_words.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for {
		response := read(conn)
		fmt.Println(response) // Display server messages for visibility

		if strings.Contains(response, "completed") {
			break
		} else if strings.Contains(response, "ERROR") {
			fmt.Println("Error:", response)
			return
		}

		// Write each received line to the file (strip the prefix)
		line := strings.TrimPrefix(response, clientKey+"_FILE:")
		writer.WriteString(line + "\n")
	}

	// Ensure all buffered data is written to the file
	writer.Flush()
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

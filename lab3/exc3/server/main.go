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

var (
	clientKeys sync.Map
	files      string
)

func main() {
	fmt.Println("username:", USERNAME)
	fmt.Println("password:", PASSWORD)
	generateWordFile()

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
		handleClient(conn)
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
	serveRequests(conn, clientKey)
}

func authenticate(conn net.Conn) bool {
	send(conn, "AUTH:Enter username:")
	username := read(conn, "")

	send(conn, "AUTH:Enter password:")
	password := read(conn, "")

	return username == USERNAME && password == PASSWORD
}

func serveRequests(conn net.Conn, clientKey string) {
	request := read(conn, clientKey)
	if strings.HasPrefix(request, "DOWNLOAD:") {
		filename := strings.TrimPrefix(request, "DOWNLOAD:")
		if filename == "words.txt" {
			handleFileDownload(conn, clientKey)
		} else {
			send(conn, fmt.Sprintf("%s_ERROR:File not found.", clientKey))
		}
	} else {
		send(conn, fmt.Sprintf("%s_ERROR:Unknown command.", clientKey))
	}
}

func handleFileDownload(conn net.Conn, clientKey string) {
	send(conn, fmt.Sprintf("%s_MSG:Starting transfer for words.txt", clientKey))

	send(conn, files)

	send(conn, fmt.Sprintf("%s_MSG:Transfer completed.", clientKey))
}

func generateWordFile() {
	wordList := []string{
		"apple", "banana", "cherry", "dog", "elephant", "fish", "giraffe", "house", "ice", "jungle",
		"kangaroo", "lemon", "mountain", "night", "ocean", "piano", "queen", "river", "sun", "tree",
		"umbrella", "violin", "whale", "xylophone", "yacht", "zebra", "ant", "bridge", "cloud", "door",
		"earth", "forest", "gold", "horse", "island", "jacket", "kite", "light", "moon", "nest",
		"orange", "penguin", "quiet", "road", "star", "table", "unicorn", "vase", "wolf", "yard",
		"zone", "beach", "camera", "dragon", "engine", "flower", "garden", "hill", "insect", "jewel",
		"knife", "lake", "music", "novel", "oyster", "planet", "quartz", "rope", "sword", "tiger",
		"universe", "volcano", "wind", "year", "zoo", "alien", "balloon", "castle", "dolphin", "engineer",
		"fire", "ghost", "hammer", "iceberg", "jungle", "keyboard", "lamp", "mountain", "nurse", "octopus",
		"parrot", "queen", "robot", "spaceship", "telescope", "umbrella", "vampire", "wizard", "yogurt", "zeppelin",
	}

	var builder strings.Builder
	for i := 0; i < 500; i++ {
		word := wordList[rand.Intn(len(wordList))]
		builder.WriteString(word + " ")
	}
	files = builder.String()
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
		log.Println("Error reading from connection:", err)
	}
	msg := strings.TrimSpace(string(buffer[:n]))

	if clientKey != "" && !strings.HasPrefix(msg, clientKey+"_") {
		send(conn, fmt.Sprintf("%s_ERROR:Invalid key prefix. Disconnecting.", clientKey))
		panic("Invalid key prefix")
	}

	return strings.TrimPrefix(msg, clientKey+"_")
}

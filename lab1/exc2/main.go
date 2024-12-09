package main

import (
	"fmt"

	"math/rand"
)

func main() {
	for i := 0; i < 1000; i++ {
		len := rand.Intn(100)

		str := GenerateRandomString(len)

		res := ScabbleScore(str)

		fmt.Printf("Word=%s\nScrabble Score=%d\n\n", str, res)
	}
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ScabbleScore(word string) int {
	score := 0
	for _, letter := range word {
		score += letterScore(letter)
	}
	return score
}

func letterScore(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

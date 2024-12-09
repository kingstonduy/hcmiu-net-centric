package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gammazero/workerpool"
	cmap "github.com/orcaman/concurrent-map/v2"
)

func countChars(segment string, freqMap cmap.ConcurrentMap[string, int]) {
	for _, char := range segment {
		charStr := string(char)
		if val, ok := freqMap.Get(charStr); ok {
			freqMap.Set(charStr, val+1)
		} else {
			freqMap.Set(charStr, 1)
		}
	}
}

func main() {
	// Prompt user for input
	fmt.Println("Enter a line of text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()

	freqMap := cmap.New[int]()

	numWorkers := 4
	partSize := len(input) / numWorkers

	pool := workerpool.New(numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * partSize
		end := start + partSize
		if i == numWorkers-1 {
			end = len(input)
		}

		segment := input[start:end]

		pool.Submit(func() {
			countChars(segment, freqMap)
		})
	}

	pool.StopWait()

	for item := range freqMap.IterBuffered() {
		fmt.Printf("%q: %d\n", item.Key, item.Val)
	}
}

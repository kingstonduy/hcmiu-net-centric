package main

import (
	"context"
	"fmt"
	logger "lab1/util"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()

	for i := 0; i < 1000; i++ {
		len := rand.Intn(100)
		dna1 := GenerateRandomDNA(len)
		dna2 := GenerateRandomDNA(len)

		res, err := HammingDistance(dna1, dna2)
		if err != nil {
			logger.Error(ctx, err)
		}

		fmt.Printf("DNA1=%s\nDNA2=%s\nHamming Distance=%d\n\n", dna1, dna2, res)
	}
}

func HammingDistance(dna1, dna2 string) (int, error) {
	if len(dna1) != len(dna2) {
		return -1, fmt.Errorf("DNA strings are not equal in length")
	}

	distance := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] != dna2[i] {
			distance++
		}
	}
	return distance, nil
}

func GenerateRandomDNA(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const dnaBases = "ACGT"
	dna := make([]byte, length)
	for i := range dna {
		dna[i] = dnaBases[r.Intn(len(dnaBases))]
	}
	return string(dna)
}

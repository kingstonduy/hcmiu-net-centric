package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Student struct {
	ID       int
	Duration int
}

var (
	start int
)

func visitLibrary(student Student, wg *sync.WaitGroup, library chan struct{}) {
	defer wg.Done()

	library <- struct{}{}

	log.Printf("Time %d: Student %d started reading at the library for %d hours\n", time.Now().Second()-start, student.ID, student.Duration)
	time.Sleep(time.Duration(student.Duration) * time.Second)

	<-library
	log.Printf("Time %d: Student %d finished reading. Spent %d hours\n", time.Now().Second()-start, student.ID, student.Duration)
}

func generateStudents(numStudents int) []Student {
	students := make([]Student, numStudents)
	for i := 0; i < numStudents; i++ {
		students[i] = Student{
			ID:       i + 1,
			Duration: rand.Intn(4) + 1,
		}
	}

	rand.Shuffle(numStudents, func(i, j int) {
		students[i], students[j] = students[j], students[i]
	})

	return students
}

func main() {
	start = time.Now().Second()
	libraryCapacity := 30
	totalStudents := 100
	library := make(chan struct{}, libraryCapacity)

	students := generateStudents(totalStudents)

	var wg sync.WaitGroup
	for i := 0; i < len(students); i++ {
		wg.Add(1)
		go visitLibrary(students[i], &wg, library)
	}

	wg.Wait()

	log.Printf("Time %d: No more students. Let's call it a day\n", time.Now().Second()-start)
}

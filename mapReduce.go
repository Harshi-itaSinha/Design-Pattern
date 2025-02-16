package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Map function: Splits text into words and counts occurrences
func Map(text string, out chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	wordCounts := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		wordCounts[word]++
	}
	out <- wordCounts
	fmt.Println(wordCounts)
}

// Reduce function: Combines multiple maps into a single result
func Reduce(in <-chan map[string]int) map[string]int {
	finalResult := make(map[string]int)
	for partialResult := range in { // many maps here []
		for word, count := range partialResult { // interaing throught sech map 
			finalResult[word] += count
		}
	}
	return finalResult
}



func main() {
	// Sample input texts
	texts := []string{
		"hello world hello",
		"go is fun go is powerful",
		"map reduce is useful",
	}

	startTime := time.Now() // Start execution time measurement

	out := make(chan map[string]int, len(texts)) // buffered with capaacity of 3 
	var wg sync.WaitGroup

	// Launch Map tasks
	for _, text := range texts {
		wg.Add(1)
		go Map(text, out, &wg)
	}

	// Wait for all goroutines to finish then close channel
	//go func() {
		wg.Wait()
		close(out)
	//}()

	// Reduce step
	finalResult := Reduce(out)

	// Measure execution time
	elapsedTime := time.Since(startTime)

	// Print the word count result and execution time
	fmt.Println("Word Count:", finalResult)
	fmt.Println("Execution Time:", elapsedTime)
}

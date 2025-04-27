// cmd/kafka-performance/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aleksandr/kafka-consumer-demo/internal/consumer"
	"github.com/segmentio/kafka-go"
)


func benchmarkForLoop(reader *kafka.Reader, maxMessages int) time.Duration {
	startTime := time.Now()
	count := 0

	for count < maxMessages {
		_, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}
		
		count++
	}

	return time.Since(startTime)
}

func benchmarkWithGoroutines(reader *kafka.Reader, maxMessages int) time.Duration {
	startTime := time.Now()
	count := 0
	sem := make(chan struct{}, 300)

	for count < maxMessages {
		_, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}
		count++

		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
		}()
	}

	// Wait for all goroutines to finish
	for range cap(sem) {
		sem <- struct{}{}
	}

	return time.Since(startTime)
}

func RunPerformanceTest() {
	const messageCount = 1000

	fmt.Println("Benchmarking with Goroutines (Concurrency)...")
	reader1 := consumer.NewKafkaReader()
	defer reader1.Close()
	timeWithGoroutines := benchmarkWithGoroutines(reader1, messageCount)
	fmt.Printf("Time taken with goroutines: %v\n", timeWithGoroutines)

	fmt.Println("Benchmarking with For Loop (No concurrency)...")
	reader2 := consumer.NewKafkaReader()
	defer reader2.Close()
	timeWithForLoop := benchmarkForLoop(reader2, messageCount)
	fmt.Printf("Time taken with for loop: %v\n", timeWithForLoop)
}

func main() {
	RunPerformanceTest()
}

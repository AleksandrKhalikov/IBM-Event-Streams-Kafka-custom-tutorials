package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aleksandr/kafka-consumer-demo/internal/consumer"
	"github.com/aleksandr/kafka-consumer-demo/internal/util"
)

func main() {
	reader := consumer.NewKafkaReader()
	defer reader.Close()

	fmt.Printf("Consuming from topic: %s\n", reader.Config().Topic)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("Message received:\n")
		helper_functions.PrettyPrintJSON(m.Value)
	}
}

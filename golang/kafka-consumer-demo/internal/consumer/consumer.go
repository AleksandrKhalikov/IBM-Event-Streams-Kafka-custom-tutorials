package consumer

import (
    "log"
    "time"

    "github.com/segmentio/kafka-go"
    "github.com/segmentio/kafka-go/sasl/scram"
    "github.com/aleksandr/kafka-consumer-demo/internal/config"
)

func NewKafkaReader() *kafka.Reader {
    cfg := config.GetKafkaConfig()

	// SCRAM mechanism
	mechanism, err := scram.Mechanism(scram.SHA512, cfg.Username, cfg.Password)
	if err != nil {
		log.Fatalf("failed to create SCRAM mechanism: %v", err)
	}

    dialer := &kafka.Dialer{
        Timeout:       10 * time.Second,
        DualStack:     true,
        SASLMechanism: mechanism,
        TLS:           config.GetSSLCertificate(),
    }

    return kafka.NewReader(kafka.ReaderConfig{
        Brokers:     cfg.Brokers,
        Topic:       cfg.Topic,
        Dialer:      dialer,
        StartOffset: kafka.FirstOffset,
    })
}

package config

import (
	"log"
	"crypto/x509"
	"os"
	"crypto/tls"
)

// Kafka Struct
type KafkaConfig struct {
	Brokers   []string
	Topic     string
	Username  string
	Password  string
}

// Kafka Config with credentials
func GetKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Brokers:  []string{"es-demo-kafka-bootstrap-tools.apps.67eb6a82dbbf3678605ecc92.ap1.techzone.ibm.com:443"},
		Topic:    "DOOR.BADGEIN",
		Username: "python-consumer-1",
		Password: "FDSrACfYignZqQPTJApM842jYJ8UlzWR",
	}
}

// SSL Config
func GetSSLCertificate() *tls.Config {
	
	caCert, err := os.ReadFile("../../certs/es-cert.pem")
	if err != nil {
		log.Fatalf("unable to read CA cert: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caPool,
	}

	return tlsConfig
}

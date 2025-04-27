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
		Brokers:  []string{"Your Kafka bootstrap server url"},
		Topic:    "Your Topic Name",
		Username: "Your Username",
		Password: "Your Password",
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

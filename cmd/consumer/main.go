package main

import (
	"log"

	"github.com/EgorikA4/golang-message-broker-lab/config"
	"github.com/EgorikA4/golang-message-broker-lab/internal/services/consumer"
	"github.com/EgorikA4/golang-message-broker-lab/internal/storage"
)

func main() {
    if err := config.LoadConfig(); err != nil {
        log.Fatalf("can't load .env file: %v", err)
        return
    }

    if err := storage.InitDB(); err != nil {
        log.Fatalf("could initialize Cassandra: %v", err)
        return
    }

    if err := consumer.Listen(); err != nil {
        log.Fatalf("consumer couldn't listen: %v", err)
        return
    }
}

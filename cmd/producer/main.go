package main

import (
	"github.com/EgorikA4/golang-message-broker-lab/config"
	"github.com/EgorikA4/golang-message-broker-lab/internal/services/producer"
	"log"
)

func main() {
    if err := config.LoadConfig(); err != nil {
        log.Fatalf("can't load .env file: %v", err)
        return
    }
    if err := producer.ProcessFile("logs.txt"); err != nil {
        log.Fatalf("can't process file: %v", err)
        return
    }
}

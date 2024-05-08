package main

import (
	"log"

	"github.com/egerges/GoCognitoManager/internal/config"
)

func main() {
	// Load the config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err.Description)
	}

	log.Println("Configs loaded successfully.")

	// If config loaded successfully, print Cognito values
	log.Printf("Cognito Region: %s", cfg.Cognito.Region)
	log.Printf("Cognito ClientID: %s", cfg.Cognito.ClientID)
	log.Printf("Cognito ClientSecret: %s", cfg.Cognito.ClientSecret)
	log.Printf("Cognito UserPoolID: %s", cfg.Cognito.UserPoolID)

	log.Println("Starting service:")
}

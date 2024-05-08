package config

import (
	"log"
	"os"

	"github.com/egerges/GoCognitoManager/internal/error"
	"github.com/joho/godotenv"
)

func LoadConfig() (*ConfigManager, *error.GoCognitoManagerError) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, &error.GoCognitoManagerError{Status: true, Description: "Failed to load .env file"}
	}

	config := &ConfigManager{}
	if err := OverwriteConfigWithEnv(config); err != nil {
		return nil, err
	}

	// Check if any configuration is missing
	if config.Cognito.Region == "" || config.Cognito.ClientID == "" || config.Cognito.ClientSecret == "" || config.Cognito.UserPoolID == "" {
		log.Fatalf("Failed to load configuration, no configuration exists.")
		return nil, &error.GoCognitoManagerError{Status: true, Description: "Failed to load configuration, no configuration exists."}
	}

	return config, nil
}

func CheckConfigFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func OverwriteConfigWithEnv(config *ConfigManager) *error.GoCognitoManagerError {
	// Define a map of environment variables and corresponding fields in the config struct
	envVars := map[string]*string{
		"COGNITO_REGION":        &config.Cognito.Region,
		"COGNITO_CLIENT_ID":     &config.Cognito.ClientID,
		"COGNITO_CLIENT_SECRET": &config.Cognito.ClientSecret,
		"COGNITO_USER_POOL_ID":  &config.Cognito.UserPoolID,
	}

	// Iterate over the map and update config fields if environment variables are set
	for envVar, field := range envVars {
		if value, exists := os.LookupEnv(envVar); exists {
			*field = value
		} else {
			log.Fatalf("Environment variable %s not found", envVar)
			return &error.GoCognitoManagerError{Status: true, Description: "Environment variable " + envVar + " not found"}
		}
	}

	return nil // No errors
}

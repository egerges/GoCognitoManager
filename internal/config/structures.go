package config

import "github.com/egerges/GoCognitoManager/internal/auth"

type ConfigError struct {
	Status      bool
	Description string
}

type Config struct {
	Cognito auth.Config
}

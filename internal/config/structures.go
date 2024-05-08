package config

type ConfigError struct {
	Status      bool
	Description string
}

type Config struct {
	Region       string
	ClientID     string
	UserPoolID   string
	ClientSecret string
}

type ConfigManager struct {
	Cognito Config
}

package auth

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/egerges/GoCognitoManager/internal/error"
	"github.com/egerges/GoCognitoManager/internal/flconfig"
)

func NewService() (Service, *error.GoCognitoManagerError) {
	// Load AWS config
	cfg, err := flconfig.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading AWS config from env:", err)
		return nil, &error.GoCognitoManagerError{Status: true, Description: "Failed to load .env file"}
	}

	// Log when configurations are successfully loaded
	log.Println("Configurations loaded from .env file successfully.")
	log.Println("--------------------------------------------------")
	log.Println("Values:")
	// Print Cognito values
	log.Printf("Cognito Region: %s", cfg.Cognito.Region)
	log.Printf("Cognito ClientID: %s", cfg.Cognito.ClientID)
	log.Printf("Cognito ClientSecret: %s", cfg.Cognito.ClientSecret)
	log.Printf("Cognito UserPoolID: %s", cfg.Cognito.UserPoolID)

	log.Printf("--------------------------------------------------\n\n")

	// Load Cognito config
	cognitoCfg, awsError := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Cognito.Region))
	if awsError != nil {
		log.Fatalln("Error loading AWS config:", awsError)
		return nil, &error.GoCognitoManagerError{Status: true, Description: "Error loading AWS config: " + awsError.Error()}
	}

	// Initialize Cognito client
	client := cognitoidentityprovider.NewFromConfig(cognitoCfg)

	return &service{
		cognitoClient:  client,
		cognitoManager: cfg,
	}, nil
}

// SignUp implements Service.
func (s *service) SignUp(params *SignUpParam) (*SignUpRes, *error.GoCognitoManagerError) {
	log.Printf("SignUp called with params: %+v", params)
	secretHash := calculateSecretHash(s.cognitoManager.Cognito.ClientID, s.cognitoManager.Cognito.ClientSecret, params.Email)
	input := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(s.cognitoManager.Cognito.ClientID),
		Username:   aws.String(params.Email),
		Password:   aws.String(params.Password),
		SecretHash: aws.String(secretHash),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(params.Email),
			},
			{
				Name:  aws.String("phone"),
				Value: aws.String(params.Password),
			},
		},
	}

	output, err := s.cognitoClient.SignUp(context.TODO(), input)
	if err != nil {
		log.Fatalln("Error signing up user:", err)
		return nil, &error.GoCognitoManagerError{Status: true, Description: "Failed to sign up user: " + err.Error()}
	}

	// Extract Cognito User ID
	cognitoUserID := *output.UserSub

	return &SignUpRes{
		Success:       true,
		Error:         nil,
		CognitoUserID: cognitoUserID,
	}, nil
}

// ConfirmEmail implements Service.
func (s *service) ConfirmEmail(params *ConfirmEmailParam) (*ConfirmEmailRes, *error.GoCognitoManagerError) {
	log.Printf("ConfirmEmail called with params: %+v", params)
	return nil, nil
}

// ConfirmOTP implements Service.
func (s *service) ConfirmOTP(params *ConfirmOTPParam) (*ConfirmOTPRes, *error.GoCognitoManagerError) {
	log.Printf("ConfirmOTP called with params: %+v", params)
	return nil, nil
}

// SignIn implements Service.
func (s *service) SignIn(params *SignInParam) (*SignInRes, *error.GoCognitoManagerError) {
	log.Printf("SignIn called with params: %+v", params)
	return nil, nil
}

// SignOut implements Service.
func (s *service) SignOut(params *SignOutParam) (*SignOutRes, *error.GoCognitoManagerError) {
	log.Printf("SignOut called with params: %+v", params)
	return nil, nil
}

// ValidateToken implements Service.
func (s *service) ValidateToken(params *ValidateTokenParam) (*ValidateTokenRes, *error.GoCognitoManagerError) {
	log.Printf("ValidateToken called with params: %+v", params)
	return nil, nil
}

// ChangePassword implements Service.
func (s *service) ChangePassword(params *ChangePasswordParam) (*ChangePasswordRes, *error.GoCognitoManagerError) {
	log.Printf("ChangePassword called with params: %+v", params)
	return nil, nil
}

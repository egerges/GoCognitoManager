package main

import (
	"log"

	"github.com/egerges/GoCognitoManager/internal/auth"
)

func main() {
	// Initialize the service
	service, err := auth.NewService()
	if err != nil {
		log.Fatalf("Error initializing service: %v", err)
	}

	// Test each function
	testSignUp(service)
	testSignIn(service)
	testSignOut(service)
	testConfirmOTP(service)
	testConfirmEmail(service)
	testChangePassword(service)
	testValidateToken(service)
}

func testSignUp(service auth.Service) {
	params := &auth.SignUpParam{
		Email:       "test@example.com",
		PhoneNumber: "1234567890",
		Password:    "password123",
	}

	res, err := service.SignUp(params)
	if err != nil {
		log.Printf("Error signing up: %v", err)
		return
	}

	log.Printf("SignUp successful. Response: %+v", res)
}

func testSignIn(service auth.Service) {
	params := &auth.SignInParam{
		Username: "test@example.com",
		Password: "password123",
	}

	res, err := service.SignIn(params)
	if err != nil {
		log.Printf("Error signing in: %v", err)
		return
	}

	log.Printf("SignIn successful. Response: %+v", res)
}

func testSignOut(service auth.Service) {
	params := &auth.SignOutParam{
		Username: "test@example.com",
	}

	res, err := service.SignOut(params)
	if err != nil {
		log.Printf("Error signing out: %v", err)
		return
	}

	log.Printf("SignOut successful. Response: %+v", res)
}

func testConfirmOTP(service auth.Service) {
	params := &auth.ConfirmOTPParam{
		Username: "test@example.com",
		Code:     "123456",
	}

	res, err := service.ConfirmOTP(params)
	if err != nil {
		log.Printf("Error confirming OTP: %v", err)
		return
	}

	log.Printf("ConfirmOTP successful. Response: %+v", res)
}

func testConfirmEmail(service auth.Service) {
	params := &auth.ConfirmEmailParam{
		Username: "test@example.com",
		Code:     "123456",
	}

	res, err := service.ConfirmEmail(params)
	if err != nil {
		log.Printf("Error confirming email: %v", err)
		return
	}

	log.Printf("ConfirmEmail successful. Response: %+v", res)
}

func testChangePassword(service auth.Service) {
	params := &auth.ChangePasswordParam{
		Username:    "test@example.com",
		OldPassword: "password123",
		NewPassword: "newpassword456",
	}

	res, err := service.ChangePassword(params)
	if err != nil {
		log.Printf("Error changing password: %v", err)
		return
	}

	log.Printf("ChangePassword successful. Response: %+v", res)
}

func testValidateToken(service auth.Service) {
	params := &auth.ValidateTokenParam{
		Token: "sometoken123",
	}

	res, err := service.ValidateToken(params)
	if err != nil {
		log.Printf("Error validating token: %v", err)
		return
	}

	log.Printf("ValidateToken successful. Response: %+v", res)
}

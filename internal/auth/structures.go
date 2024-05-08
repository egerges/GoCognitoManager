package auth

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/egerges/GoCognitoManager/internal/error"
	"github.com/egerges/GoCognitoManager/internal/flconfig"
)

type service struct {
	cognitoClient  *cognitoidentityprovider.Client
	cognitoManager *flconfig.ConfigManager
}

type Service interface {
	SignUp(params *SignUpParam) (*SignUpRes, *error.GoCognitoManagerError)
	SignIn(params *SignInParam) (*SignInRes, *error.GoCognitoManagerError)
	SignOut(params *SignOutParam) (*SignOutRes, *error.GoCognitoManagerError)
	ConfirmEmail(params *ConfirmEmailParam) (*ConfirmEmailRes, *error.GoCognitoManagerError)
	ConfirmOTP(params *ConfirmOTPParam) (*ConfirmOTPRes, *error.GoCognitoManagerError)
	ChangePassword(params *ChangePasswordParam) (*ChangePasswordRes, *error.GoCognitoManagerError)
	ValidateToken(params *ValidateTokenParam) (*ValidateTokenRes, *error.GoCognitoManagerError)
}

type SignUpParam struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type SignUpRes struct {
	Error         *error.GoCognitoManagerError `json:"error"`
	Success       bool                         `json:"success"`
	CognitoUserID string                       `json:"cognitoUserID"`
}

type SignInParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRes struct {
	Error   *error.GoCognitoManagerError `json:"error"`
	Success bool                         `json:"success"`
	Token   string                       `json:"token"`
}

type SignOutParam struct {
	Username string `json:"username"`
}

type SignOutRes struct {
	Error   *error.GoCognitoManagerError `json:"error"`
	Success bool                         `json:"success"`
}

type ConfirmOTPParam struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}

type ConfirmOTPRes struct {
	Error   *error.GoCognitoManagerError `json:"error"`
	Success bool                         `json:"success"`
}

type ConfirmEmailParam struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}

type ConfirmEmailRes struct {
	Error   *error.GoCognitoManagerError `json:"error"`
	Success bool                         `json:"success"`
}

type ChangePasswordParam struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordRes struct {
	Error   *error.GoCognitoManagerError `json:"error"`
	Success bool                         `json:"success"`
}

type ValidateTokenParam struct {
	Token string `json:"token"`
}

type ValidateTokenRes struct {
	Error *error.GoCognitoManagerError `json:"error"`
	Valid bool                         `json:"valid"`
}

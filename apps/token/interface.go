package token

import "context"

const (
	AppName = "token"
)

type Service interface {
	Login(context.Context, *LoginRequest) (*Token, error) // return the token
	Logout(context.Context, *LogoutRequest) error         // delete the token

	ValidateToken(context.Context, *ValidateToken) (*Token, error) // validate the token
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ValidateToken struct {
	AccessToken string `json:"access_token"`
}

func NewValidateToken(aToken string) *ValidateToken {
	return &ValidateToken{
		AccessToken: aToken,
	}
}

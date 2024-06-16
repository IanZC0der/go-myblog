package token

import "context"

type Service interface {
	Login(context.Context, *LoginRequest) (*Token, error) // return the token
	Logout(context.Context, *LogoutRequest) error         // delete the token

	ValidateToken(context.Context, *ValidateToken) error // validate the token
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ValidateToken struct {
	AccessToken string `json:"access_token"`
}

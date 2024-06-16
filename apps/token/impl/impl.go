package impl

import (
	"context"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/exception"

	"gorm.io/gorm"
)

// var AuthFailed = exception.NewAuthFailed("Incorrect Username/Password")

type TokenServiceImpl struct {
	db   *gorm.DB
	user user.Service
}

func NewTokenServiceImpl(userServiceImpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		db:   conf.C().MySQL.GetConn(),
		user: userServiceImpl,
	}
}

func (t *TokenServiceImpl) Login(ctx context.Context, req *token.LoginRequest) (*token.Token, error) {
	// get user
	queryRequest := user.NewQueryUserRequestByUsername(req.Username)
	userQueried, err := t.user.QueryUser(ctx, queryRequest)
	if err != nil {
		if exception.IsNotFound(err) {
			return nil, token.AuthFailed
		}

		return nil, err
	}
	// validate the user
	if err := userQueried.ValidatePassword(req.Password); err != nil {
		return nil, token.AuthFailed
	}
	// generate token

	tk := token.NewToken()
	tk.UserId = userQueried.Id
	tk.UserName = userQueried.Username

	if err := t.db.WithContext(ctx).Create(tk).Error; err != nil {
		return nil, err
	}
	// delete the old token after generating the new token, to be impled

	return tk, nil
}

func (t *TokenServiceImpl) Logout(ctx context.Context, req *token.LogoutRequest) error {
	return nil
}

func (t *TokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateToken) (*token.Token, error) {

	// query the token from the db
	tk := token.NewToken()
	if err := t.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(tk).Error; err != nil {
		return nil, err
	}
	// validate the token
	if err := tk.IsExpired(); err != nil {
		return nil, err
	}

	return tk, nil
}

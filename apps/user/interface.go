package user

import (
	"context"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
}

type CreateUserRequest struct {
	Username string            `json:"username" gorm:"column:username"`
	Password string            `json:"password" gorm:"column:password"`
	Role     Role              `json:"role"`
	Label    map[string]string `json:"label" gorm:"serializer:json"`
	isHashed bool
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  ROLE_MEMBER,
		Label: map[string]string{},
	}
}

func (req *CreateUserRequest) Validate() error {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("Empty username/password")
	}
	return nil
}

func (req *CreateUserRequest) PasswordHash() {
	if req.isHashed {
		return
	}
	b, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	req.Password = base64.StdEncoding.EncodeToString(b)

	req.isHashed = true
}

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

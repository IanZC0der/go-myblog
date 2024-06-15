package user

import (
	"context"
	"fmt"
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

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

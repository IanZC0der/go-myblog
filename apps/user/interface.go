package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	QueryUser(context.Context, *QueryUserRequest) (*User, error)
}

type QueryUserRequest struct {
	Queryby    QueryBy `json:"query_by"`
	QueryValue string  `json:"query_value"`
}

func NewQueryUserRequestById(id string) *QueryUserRequest {
	return &QueryUserRequest{
		Queryby:    QUERY_BY_ID,
		QueryValue: id,
	}
}

func NewQueryUserRequestByUsername(username string) *QueryUserRequest {
	return &QueryUserRequest{
		Queryby:    QUERY_BY_USERNAME,
		QueryValue: username,
	}
}

type CreateUserRequest struct {
	Username string            `json:"username" gorm:"column:username"`
	Password string            `json:"password" gorm:"column:password"`
	Role     Role              `json:"role"`
	Label    map[string]string `json:"label" gorm:"serializer:json"`
	IsHashed bool              `json: "is_hashed", gorm:"column:is_hashed"`
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:     ROLE_MEMBER,
		Label:    map[string]string{},
		IsHashed: false,
	}
}

func (req *CreateUserRequest) Validate() error {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("Empty username/password")
	}
	return nil
}

func (req *CreateUserRequest) PasswordHash() {
	if req.IsHashed {
		return
	}
	b, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	req.Password = string(b)

	req.IsHashed = true
}

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

package user

import "context"

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

type Role int

const (
	ROLE_MEMBER Role = iota
	ROLE_ADMIN
)

package impl

import (
	"context"

	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/conf"

	"gorm.io/gorm"
)

var _ user.Service = &UserServiceImpl{} // enforce interface implementation
// var _ user.Service = (*UserServiceImpl)(nil) another way to enforce interface implementation

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: conf.C().MySQL.GetConn(),
	}
}

func (u *UserServiceImpl) CreateUser(
	context.Context,
	*user.CreateUserRequest) (*user.User, error) {

	// validate the params

	// create the user entity
	// save to the db
	// return the user entity
	return nil, nil
}

func (u *UserServiceImpl) DeleteUser(
	context.Context,
	*user.DeleteUserRequest) error {
	return nil
}

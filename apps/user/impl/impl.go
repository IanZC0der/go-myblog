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
	ctx context.Context,
	req *user.CreateUserRequest) (*user.User, error) {

	// validate the params

	if err := req.Validate(); err != nil {
		return nil, err
	}

	// create the user entity

	newUser := user.NewUser(req)
	// save to the db
	if err := u.db.WithContext(ctx).Create(newUser).Error; err != nil {
		return nil, err
	}
	// return the user entity
	return newUser, nil
}

func (u *UserServiceImpl) DeleteUser(
	ctx context.Context,
	req *user.DeleteUserRequest) error {

	err := u.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&user.User{}).Error
	return err
}

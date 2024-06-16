package impl

import (
	"context"

	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/exception"
	"github.com/IanZC0der/go-myblog/ioc"

	"gorm.io/gorm"
)

func init() {
	ioc.DefaultControllerContainer().Register(&UserServiceImpl{})
}

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

func (u *UserServiceImpl) Init() error {
	u.db = conf.C().MySQL.GetConn()
	return nil
}

func (u *UserServiceImpl) Name() string {
	return user.AppName
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

func (u *UserServiceImpl) QueryUser(
	ctx context.Context,
	req *user.QueryUserRequest) (*user.User, error) {
	// query the user from the db

	query := u.db.WithContext(ctx)
	switch req.Queryby {
	case user.QUERY_BY_ID:
		query = query.Where("id = ?", req.QueryValue)
	case user.QUERY_BY_USERNAME:
		query = query.Where("username = ?", req.QueryValue)
	}
	newUser := user.NewUser(user.NewCreateUserRequest())
	if err := query.First(newUser).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("user %s not found", req.QueryValue)
		}
		return nil, err
	}
	return newUser, nil
}

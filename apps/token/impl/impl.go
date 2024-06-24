package impl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/exception"
	"github.com/IanZC0der/go-myblog/ioc"

	"gorm.io/gorm"
)

// var AuthFailed = exception.NewAuthFailed("Incorrect Username/Password")

type TokenServiceImpl struct {
	db   *gorm.DB
	user user.Service
}

func init() {
	ioc.DefaultControllerContainer().Register(&TokenServiceImpl{})
}

func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{
		db:   conf.C().MySQL.GetConn(),
		user: ioc.DefaultControllerContainer().Get(user.AppName).(user.Service),
	}
}

func (t *TokenServiceImpl) Init() error {
	t.db = conf.C().MySQL.GetConn()
	t.user = ioc.DefaultControllerContainer().Get(user.AppName).(user.Service)
	return nil
}

func (t *TokenServiceImpl) Name() string {
	return token.AppName
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
	tk.Role = userQueried.Role

	// query the token by id
	oldToken, err := t.QueryTokenBy(ctx, token.NewQueryTokenRequestById(strconv.Itoa(int(tk.UserId))))

	if err != nil {
		if !exception.IsNotFound(err) {
			return nil, token.AuthFailed
		}
	}
	// update the old token after generating the new token, to be impled

	if oldToken != nil {
		if err = t.db.WithContext(ctx).Model(&oldToken).Where("user_name = ?", oldToken.UserName).Updates(tk).Error; err != nil {
			return nil, err
		}
		return oldToken, nil
	} else {
		if err := t.db.WithContext(ctx).Create(tk).Error; err != nil {
			return nil, err
		}

		return tk, nil

	}

}

func (t *TokenServiceImpl) Logout(ctx context.Context, req *token.LogoutRequest) error {
	tk, err := t.QueryTokenBy(ctx, token.NewQueryTokenRequestByToken(req.AccessToken))

	if err != nil {
		return err
	}

	return t.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).Delete(&tk).Error
}

func (t *TokenServiceImpl) QueryTokenBy(ctx context.Context, req *token.QueryTokenRequest) (*token.Token, error) {
	query := t.db.WithContext(ctx)
	switch req.Queryby {
	case token.QUERY_BY_ID:
		query = query.Where("user_id = ?", req.QueryValue)
	case token.QUERY_BY_ACCESS_TOKEN:
		query = query.Where("access_token = ?", req.QueryValue)
	}
	// fmt.Println(a ...any)
	var oldToken *token.Token
	if err := query.First(&oldToken).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("token %s not found", req.QueryValue)
		}
		return nil, err
	}
	return oldToken, nil
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

	// get the user role, query the db
	queryUser := user.NewQueryUserRequestById(fmt.Sprintf("%d", tk.UserId))

	theUser, err := t.user.QueryUser(ctx, queryUser)

	if err != nil {
		return nil, err
	}

	tk.Role = theUser.Role

	return tk, nil
}

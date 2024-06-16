package impl_test

import (
	"context"
	"testing"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/apps/token/impl"
	userImpl "github.com/IanZC0der/go-myblog/apps/user/impl"
	"github.com/IanZC0der/go-myblog/test"
)

var (
	tokenSvc *impl.TokenServiceImpl
	ctx      = context.Background()
)

// func init() {
// 	userSvc = impl.NewUserServiceImpl()
// }

func init() {
	test.DevelopmentSetup()
	tokenSvc = impl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
}

func TestLogin(t *testing.T) {
	req := &token.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	tk, err := tokenSvc.Login(ctx, req)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)

}

func TestValidateToken(t *testing.T) {
	req := &token.ValidateToken{
		AccessToken: "cpn3jbiclaaljs6rmt50",
	}

	tk, err := tokenSvc.ValidateToken(ctx, req)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(tk)

}

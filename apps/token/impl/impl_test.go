package impl_test

import (
	"context"
	"testing"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/test"
)

var (
	tokenSvc token.Service
	ctx      = context.Background()
)

// func init() {
// 	userSvc = impl.NewUserServiceImpl()
// }

func init() {
	test.DevelopmentSetup()
	tokenSvc = ioc.DefaultControllerContainer().Get(token.AppName).(token.Service)
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

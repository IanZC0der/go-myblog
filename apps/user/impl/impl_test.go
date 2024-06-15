package impl_test

import (
	"context"
	"testing"

	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/apps/user/impl"
	"github.com/IanZC0der/go-myblog/test"
)

var (
	userSvc *impl.UserServiceImpl
	ctx     = context.Background()
)

// func init() {
// 	userSvc = impl.NewUserServiceImpl()
// }

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "testuser"
	req.Password = "testpassword"
	u, err := userSvc.CreateUser(ctx, req)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)

}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: 2,
	})

	if err != nil {
		t.Fatal(err)
	}

	// t.Log(u)
}

func TestQueryUser(t *testing.T) {
	u, err := userSvc.QueryUser(ctx, user.NewQueryUserRequestById("3"))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)

}

func TestQueryUserByUsername(t *testing.T) {

	u, err := userSvc.QueryUser(ctx, user.NewQueryUserRequestByUsername("testuser"))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)

}

func init() {
	test.DevelopmentSetup()
	userSvc = impl.NewUserServiceImpl()
}

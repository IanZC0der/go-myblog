package impl_test

import (
	"context"
	"testing"

	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/apps/user/impl"
)

var (
	userSvc *impl.UserServiceImpl
	ctx     = context.Background()
)

func init() {
	userSvc = &impl.UserServiceImpl{}
}

func TestCreateUser(t *testing.T) {
	u, err := userSvc.CreateUser(ctx, &user.CreateUserRequest{})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)

}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{})

	if err != nil {
		t.Fatal(err)
	}

	// t.Log(u)
}

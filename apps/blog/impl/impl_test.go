package impl_test

import (
	"context"
	"testing"

	"os"

	"github.com/IanZC0der/go-myblog/apps/blog"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/test"
)

var (
	blogSvc blog.Service
	ctx     = context.Background()
)

func init() {
	test.DevelopmentSetup()
	blogSvc = ioc.DefaultControllerContainer().Get(blog.AppName).(blog.Service)
}

func TestCreateBlog(t *testing.T) {

	bg := blog.NewCreateBlogRequest()
	bg.Title = "firstblog"
	bg.Author = "testuser"
	bg.Content = "this is my first blog"
	bg.Abstract = "firstblog"

	b, err := blogSvc.CreateBlog(ctx, bg)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(b)
}

func TestGetWorkingPwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dir)
}

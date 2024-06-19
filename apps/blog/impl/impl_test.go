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

func TestQueryBlogs(t *testing.T) {
	newReq := blog.NewQueryBlogRequest()
	newReq.SetStatus(blog.PUBLISHED)
	blogs, err := blogSvc.QueryBlog(ctx, newReq)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(blogs)
}

func TestQuerySingleBlog(t *testing.T) {
	// blog.NewCreateBlogRequest()
	newReq := blog.NewQuerySingleBlogRequest("1")

	oneBlog, err := blogSvc.QuerySingleBlog(ctx, newReq)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(oneBlog)
}

func TestUpdateBlog(t *testing.T) {
	newReq := blog.NewUpdateBlogRequest("1")

	newReq.Abstract = "New Abstact 3"
	newReq.Content = "New Content 3"
	newReq.Title = "New Title 3"
	newReq.Author = "testuser"
	newReq.UpdateMode = blog.UPDATE_MODE_PATCH

	updatedBlog, err := blogSvc.UpdateBlog(ctx, newReq)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(updatedBlog)
}

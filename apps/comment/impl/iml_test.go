package impl_test

import (
	"context"
	"testing"

	"github.com/IanZC0der/go-myblog/apps/comment"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/test"
)

var (
	commentSvc comment.Service
	ctx        = context.Background()
)

func init() {
	test.DevelopmentSetup()
	commentSvc = ioc.DefaultControllerContainer().Get(comment.AppName).(comment.Service)
}

func TestAddComment(t *testing.T) {
	req := comment.NewAddCommentRequest()
	req.BlogId = 1
	req.Content = "This is a test comment"
	req.CreatedBy = "testuser"

	c, err := commentSvc.AddComment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c)

}

func TestGetAllCommentsByBlogId(t *testing.T) {
	req := comment.NewGetAllCommentRequest()
	req.BlogId = 1

	c, err := commentSvc.GetAllCommentsByBlogId(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)

}

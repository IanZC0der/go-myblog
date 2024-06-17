package blog

import "context"

const (
	AppName = "blog"
)

// blog service interface
type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)

	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	DeleteBlog(context.Context, *DeleteBlogRequest) error
}

type UpdateBlogStatusRequest struct {
	BlogId int64 `json:"blog_id"`

	Status Status `json:"status"`
}

type UpdateBlogRequest struct {
	BlogId     int64      `json:"blog_id"`
	UpdateMode UpdateMode `json:"update_mode"`
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	BlogId int64 `json:"blog_id"`
}

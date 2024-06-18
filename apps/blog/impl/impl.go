package impl

import (
	"context"

	"github.com/IanZC0der/go-myblog/apps/blog"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.DefaultControllerContainer().Register(&blogServiceImpl{})
}

type blogServiceImpl struct {
	db *gorm.DB
}

func (b *blogServiceImpl) Init() error {
	b.db = conf.C().MySQL.GetConn()
	return nil
}

func (b *blogServiceImpl) Name() string {
	return blog.AppName
}

func (b *blogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	// validate the params

	if err := req.Validate(); err != nil {
		return nil, err
	}

	// create the blog entity

	newBlog := blog.NewBlog(req)
	// save to the db
	if err := b.db.WithContext(ctx).Create(newBlog).Error; err != nil {
		return nil, err
	}
	// return the blog entity
	return newBlog, nil
}

func (b *blogServiceImpl) UpdateBlogStatus(ctx context.Context, req *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

func (b *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

func (b *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) error {
	return nil
}

package impl

import (
	"context"
	"fmt"
	"time"

	"dario.cat/mergo"
	"github.com/IanZC0der/go-myblog/apps/blog"

	// mqimpl "github.com/IanZC0der/go-myblog/apps/mq/impl"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/exception"
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
	newReq := blog.NewQuerySingleBlogRequest(req.BlogId)
	blogToBeUpdated, err := b.QuerySingleBlog(ctx, newReq)

	if err != nil {
		return nil, err
	}

	blogToBeUpdated.Status = req.Status

	if blogToBeUpdated.Status == blog.DRAFT {
		blogToBeUpdated.PublishedAt = 0
	} else if blogToBeUpdated.Status == blog.PUBLISHED {
		blogToBeUpdated.PublishedAt = time.Now().Unix()
	}

	err = b.db.WithContext(ctx).Model(&blogToBeUpdated).Updates(map[string]interface{}{"status": req.Status, "published_at": blogToBeUpdated.PublishedAt}).Error

	if err != nil {
		return nil, err
	}

	return blogToBeUpdated, nil
}

func (b *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	newReq := blog.NewQuerySingleBlogRequest(req.BlogId)
	blogToBeUpdated, err := b.QuerySingleBlog(ctx, newReq)

	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		blogToBeUpdated.CreateBlogRequest = req.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		err := mergo.Merge(blogToBeUpdated.CreateBlogRequest, req.CreateBlogRequest,
			mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid update mode: %d", req.UpdateMode)
	}

	blogToBeUpdated.UpdatedAt = time.Now().Unix()
	err = b.db.WithContext(ctx).Where("id = ?", req.BlogId).Updates(blogToBeUpdated).Error

	if err != nil {
		return nil, err
	}
	return blogToBeUpdated, nil
}

func (b *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) error {
	err := b.db.WithContext(ctx).Where("id = ?", req.BlogId).Delete(&blog.Blog{}).Error
	return err
}

func (b *blogServiceImpl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogList, error) {

	query := b.db.WithContext(ctx).Model(&blog.Blog{})

	blogList := blog.NewBlogList()

	//query based on the status
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	if req.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+req.Keywords+"%")
	}
	//get the count of the blogs
	err := query.Count(&blogList.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(req.Offset()).Limit(req.PageSize).Find(&blogList.Items).Error
	if err != nil {
		return nil, err
	}
	return blogList, nil
}

func (b *blogServiceImpl) QuerySingleBlog(ctx context.Context, req *blog.QuerySingleBlogRequest) (*blog.Blog, error) {
	oneBlog := blog.NewBlog(blog.NewCreateBlogRequest())
	query := b.db.WithContext(ctx).Model(&blog.Blog{})

	if err := query.Where("id = ?", req.BlogId).Find(oneBlog).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("blog %s not found", req.BlogId)
		}
		return nil, err
	}
	return oneBlog, nil

}

// func ConsumeMessages()

func (b *blogServiceImpl) AuditBlog(ctx context.Context, req *blog.AuditBlogRequest) (*blog.Blog, error) {
	theBlog, err := b.QuerySingleBlog(ctx, blog.NewQuerySingleBlogRequest(req.BlogId))

	if err != nil {
		return nil, err
	}

	theBlog.AuditPassed = true
	theBlog.AuditAt = time.Now().Unix()

	err = b.db.WithContext(ctx).Where("id = ?", req.BlogId).Updates(theBlog).Error

	if err != nil {
		return nil, err
	}

	return theBlog, nil

}

package impl

import (
	"context"

	"github.com/IanZC0der/go-myblog/apps/comment"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/ioc"

	"gorm.io/gorm"
)

func init() {
	ioc.DefaultControllerContainer().Register(&CommentServiceImpl{})
}

var _ comment.Service = &CommentServiceImpl{} // enforce interface implementation
// var _ user.Service = (*UserServiceImpl)(nil) another way to enforce interface implementation

type CommentServiceImpl struct {
	db *gorm.DB
}

func NewCommentServiceImpl() *CommentServiceImpl {
	return &CommentServiceImpl{
		db: conf.C().MySQL.GetConn(),
	}
}

func (c *CommentServiceImpl) Init() error {
	c.db = conf.C().MySQL.GetConn()
	return nil
}

func (c *CommentServiceImpl) Name() string {
	return comment.AppName
}

func (c *CommentServiceImpl) AddComment(ctx context.Context, req *comment.AddCommentRequest) (*comment.Comment, error) {

	// validate the params
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// create the comment
	newComment := comment.NewComment(req)
	// save to the db
	if err := c.db.WithContext(ctx).Create(newComment).Error; err != nil {
		return nil, err
	}
	return newComment, nil
}
func (c *CommentServiceImpl) GetAllCommentsByBlogId(ctx context.Context, req *comment.GetAllCommentRequest) (*comment.CommentList, error) {

	commentList := comment.NewCommentList()
	//query based on the the blog id
	query := c.db.WithContext(ctx).Model(&comment.Comment{}).Where("blog_id = ?", req.BlogId)

	//get the count
	err := query.Count(&commentList.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Find(&commentList.Items).Error
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

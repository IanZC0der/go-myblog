package comment

import (
	"context"
	"fmt"
	"strconv"
)

type Service interface {
	AddComment(context.Context, *AddCommentRequest) (*Comment, error)
	GetAllCommentsByBlogId(context.Context, *GetAllCommentRequest) (*CommentList, error)
}

type AddCommentRequest struct {
	CreatedBy string `json:"created_by" gorm:"created_by"`
	Content   string `json:"content" gorm:"content"`
	BlogId    int    `json:"blog_id" gorm:"blog_id"`
}

func (req *AddCommentRequest) Validate() error {
	if req.Content == "" {
		return fmt.Errorf("comment cannot be empty")
	}
	return nil
}

func NewAddCommentRequest() *AddCommentRequest {
	return &AddCommentRequest{}
}

func (req *GetAllCommentRequest) SetId(BlogId string) error {
	if BlogId == "" {
		return fmt.Errorf("blog id cannot be empty")
	}
	id, err := strconv.Atoi(BlogId)
	if err != nil {
		return fmt.Errorf("invalid blog id")
	}
	req.BlogId = id
	return nil
}

type GetAllCommentRequest struct {
	BlogId int `json:"blog_id" gorm:"blog_id"`
}

func NewGetAllCommentRequest() *GetAllCommentRequest {
	return &GetAllCommentRequest{}
}

type CommentList struct {
	Items []*Comment `json:"items"`
	Total int64      `json:"total"`
}

func NewCommentList() *CommentList {
	return &CommentList{
		Items: []*Comment{},
	}
}

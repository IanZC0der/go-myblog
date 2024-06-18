package blog

import (
	"fmt"
	"time"
)

type CreateBlogRequest struct {
	Title    string            `json:"title" gorm:"title"`
	Author   string            `json:"author" gorm:"author"`
	Content  string            `json:"content" gorm:"content"`
	Tags     map[string]string `json:"tags" gorm:"serializer:json"`
	Abstract string            `json:"abstract"`
}

func (req *CreateBlogRequest) Validate() error {
	if req.Title == "" || req.Author == "" || req.Content == "" || req.Abstract == "" {
		return fmt.Errorf("cannot be empty")
	}

	return nil
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type Blog struct {
	Id        int64  `json:"id" gorm:"id"`
	CreatedBy string `json:"created_by" gorm:"created_by"`

	CreatedAt int64 `json:"created_at" gorm:"created_at"`

	UpdatedAt int64 `json:"updated_at" gorm:"updated_at"`

	PublishedAt int64 `json:"published_at" gorm:"published_at"`

	Status Status `json:"status" gorm:"status"`

	*CreateBlogRequest
}

func (b *Blog) TableName() string {
	return "blogs"
}

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{

		CreatedAt: time.Now().Unix(),
		// UpdatedAt:         time.Now().Unix(),
		Status:            DRAFT,
		CreateBlogRequest: req,
	}
}

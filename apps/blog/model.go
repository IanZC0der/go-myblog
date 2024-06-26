package blog

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateBlogRequest struct {
	Title     string            `json:"title" gorm:"title"`
	CreatedBy string            `json:"created_by" gorm:"created_by"`
	Author    string            `json:"author" gorm:"author"`
	Content   string            `json:"content" gorm:"content"`
	Tags      map[string]string `json:"tags" gorm:"serializer:json"`
	Abstract  string            `json:"abstract"`
}

func (req *CreateBlogRequest) Validate() error {
	if req.Title == "" || req.Author == "" || req.Abstract == "" {
		return fmt.Errorf("title/abstract/cannot be empty")
	}

	return nil
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type Blog struct {
	Id int64 `json:"id" gorm:"id"`

	CreatedAt int64 `json:"created_at" gorm:"created_at"`

	UpdatedAt int64 `json:"updated_at" gorm:"updated_at"`

	PublishedAt int64 `json:"published_at" gorm:"published_at"`

	Status Status `json:"status" gorm:"status"`

	AuditPassed bool  `json:"audit_passed" gorm:"audit_passed"`
	AuditAt     int64 `json:"audit_at" gorm:"audit_at"`

	*CreateBlogRequest
}

func (b *Blog) TableName() string {
	return "blogs"
}

func (b *Blog) String() string {
	jsonUser, _ := json.Marshal(b)
	return string(jsonUser)
}

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{

		CreatedAt: time.Now().Unix(),
		// UpdatedAt:         time.Now().Unix(),
		Status:            DRAFT,
		CreateBlogRequest: req,
	}
}

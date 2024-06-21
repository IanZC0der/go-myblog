package blog

import (
	"context"
	"strconv"
)

const (
	AppName = "blog"
)

// blog service interface
type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)

	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	DeleteBlog(context.Context, *DeleteBlogRequest) error

	QueryBlog(context.Context, *QueryBlogRequest) (*BlogList, error)

	QuerySingleBlog(context.Context, *QuerySingleBlogRequest) (*Blog, error)

	AuditBlog(context.Context, *AuditBlogRequest) (*Blog, error)
}

type BlogList struct {
	Items []*Blog `json:"items"`
	Total int64   `json:"total"`
}

func NewBlogList() *BlogList {
	return &BlogList{
		Items: []*Blog{},
	}
}

func (bl *BlogList) Add(items ...*Blog) {
	bl.Items = append(bl.Items, items...)
}

type QueryBlogRequest struct {
	PageSize int `json:"page_size"`

	PageNumber int `json:"page_number"`
	// nil: no filter condition
	// 0: filter condition is DRAFT
	// 1: filter condition is PUBLISHED
	Status *Status `json:""`
}

type QuerySingleBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewQuerySingleBlogRequest(id string) *QuerySingleBlogRequest {
	return &QuerySingleBlogRequest{
		BlogId: id,
	}
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   5,
		PageNumber: 1,
	}
}

func (req *QueryBlogRequest) SetStatus(s Status) {
	req.Status = &s
}

func (req *QueryBlogRequest) Offset() int {
	return int(req.PageSize * (req.PageNumber - 1))
}

func (req *QueryBlogRequest) ParsePageSize(pageSize string) error {
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)

	if err != nil {
		return err
	}

	req.PageSize = int(pageSizeInt)
	return nil
}

func (req *QueryBlogRequest) ParsePageNumber(pageNumber string) error {
	pageNumberInt, err := strconv.ParseInt(pageNumber, 10, 64)

	if err != nil {
		return err
	}

	req.PageNumber = int(pageNumberInt)
	return nil
}

type UpdateBlogStatusRequest struct {
	BlogId string `json:"blog_id"`

	Status Status `json:"status"`
}

func NewUpdateBlogStatusRequest(id string) *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{
		BlogId: id,
		Status: PUBLISHED,
	}
}

type UpdateBlogRequest struct {
	BlogId     string     `json:"blog_id"`
	UpdateMode UpdateMode `json:"update_mode"`
	*CreateBlogRequest
}

func (req *UpdateBlogRequest) SetUpdateBlogRequestUpdateMode(updateMode UpdateMode) {
	req.UpdateMode = updateMode
}

func NewUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:     id,
		UpdateMode: UPDATE_MODE_PUT,

		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type DeleteBlogRequest struct {
	BlogId int64 `json:"blog_id"`
}

func NewDeleteBlogRequest() *DeleteBlogRequest {
	return &DeleteBlogRequest{}
}

func (req *DeleteBlogRequest) SetBlogId(id string) error {
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return err
	}

	req.BlogId = idInt
	return nil
}

type AuditBlogRequest struct {
	BlogId      string `json:"blog_id"`
	AuditPassed bool   `json:"audit_passed"`
}

func NewAuditBlogRequest(BlogId string) *AuditBlogRequest {
	return &AuditBlogRequest{
		BlogId: BlogId,
	}
}

package impl

import (
	"context"

	"github.com/IanZC0der/go-myblog/apps/comment"
)

type CommentServiceImpl struct {
	comment.UnimplementedCommentServiceServer
}

func (i *CommentServiceImpl) AddComment(ctx context.Context, req *comment.AddCommentRequest) (*comment.Comment, error) {
	return nil, nil

}

func (i *CommentServiceImpl) GetAllCommentsByBlogId(ctx context.Context, req *comment.GetCommentRequest) (*comment.CommentList, error) {
	return nil, nil
}

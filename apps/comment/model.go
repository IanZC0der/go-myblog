package comment

import (
	"encoding/json"
	"time"
)

type Comment struct {
	*AddCommentRequest
	CreatedAt int64 `json:"created_at" gorm:"created_at"`
}

func NewComment(req *AddCommentRequest) *Comment {
	return &Comment{
		CreatedAt:         time.Now().Unix(),
		AddCommentRequest: req,
	}
}

func (cmt *Comment) String() string {
	jsonCmt, _ := json.Marshal(cmt)
	return string(jsonCmt)
}

func (cmt *Comment) TableName() string {
	return "comment"
}

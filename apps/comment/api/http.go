package api

import (
	"github.com/IanZC0der/go-myblog/apps/comment"
	"github.com/gin-gonic/gin"

	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/response"
)

type CommentApiHandler struct {
	svc comment.Service
}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&CommentApiHandler{})
}

func (c *CommentApiHandler) Init() error {
	c.svc = ioc.DefaultControllerContainer().Get(comment.AppName).(comment.Service)
	return nil
}

func (c *CommentApiHandler) Name() string {
	return comment.AppName
}

func (c *CommentApiHandler) Registry(router gin.IRouter) {

	// we need api for creating blog, updating blog, querying blog(s), u
	v1 := router.Group("v1").Group("comments")
	// /myblog/api/v1/comments
	v1.POST("/", c.AddComment)
	// myblog/api/v1/comments/:id
	v1.GET("/:id", c.GetAllComments)

}

func (c *CommentApiHandler) AddComment(ctx *gin.Context) {

	newReq := comment.NewAddCommentRequest()
	// fmt.Println(newReq)
	err := ctx.BindJSON(newReq)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	// fmt.Println(newReq)

	newComment, err := c.svc.AddComment(ctx, newReq)

	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, newComment)

}

func (c *CommentApiHandler) GetAllComments(ctx *gin.Context) {

	newReq := comment.NewGetAllCommentRequest()
	newReq.SetId(ctx.Param("id"))

	commentList, err := c.svc.GetAllCommentsByBlogId(ctx, newReq)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, commentList)

}

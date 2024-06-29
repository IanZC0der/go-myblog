package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IanZC0der/go-myblog/apps/blog"
	mq "github.com/IanZC0der/go-myblog/apps/mq"

	// "github.com/IanZC0der/go-myblog/apps/mq/impl"
	mqimpl "github.com/IanZC0der/go-myblog/apps/mq/impl"
	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/apps/user"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/middlewares"
	"github.com/IanZC0der/go-myblog/response"
	"github.com/gin-gonic/gin"
	// "encoding/json"
)

type BlogApiHandler struct {
	svc blog.Service
}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&BlogApiHandler{})
}

func (b *BlogApiHandler) Init() error {
	b.svc = ioc.DefaultControllerContainer().Get(blog.AppName).(blog.Service)
	b.ConsumeCreateBlog()
	return nil
}

func (b *BlogApiHandler) Name() string {
	return blog.AppName
}

func (b *BlogApiHandler) Registry(router gin.IRouter) {

	// we need api for creating blog, updating blog, querying blog(s), u
	v1 := router.Group("v1").Group("blogs")
	v1.Use(middlewares.NewAuthMiddleware().Authenticator)
	// /myblog/api/v1/blogs
	v1.GET("/", b.QueryBlogList)
	// myblog/api/v1/blogs/:id
	v1.GET("/:id", b.QueryOneBlog)

	// added middleware for authorization
	v1.POST("/", middlewares.AuthorizerWithRole(user.ROLE_AUTHOR), b.CreateBlogWithMQ)
	// v1.POST("/", b.CreateBlog)
	v1.PATCH("/:id/publish", middlewares.AuthorizerWithRole(user.ROLE_AUTHOR), b.UpdateBlogStatus)
	v1.DELETE("/:id", middlewares.AuthorizerWithRole(user.ROLE_AUTHOR), b.DeleteOneBlog)
	v1.PUT("/:id", middlewares.AuthorizerWithRole(user.ROLE_AUTHOR), b.UpdateBlogAll)
	v1.PATCH("/:id", middlewares.AuthorizerWithRole(user.ROLE_AUTHOR), b.UpdateBlogPartial)
	v1.POST("/:id/audit", middlewares.AuthorizerWithRole(user.ROLE_AUDITOR), b.AuditOneBlog)

}

func (b *BlogApiHandler) UpdateBlogStatus(c *gin.Context) {

	newReq := blog.NewUpdateBlogStatusRequest(c.Param("id"))
	err := c.BindJSON(newReq)
	if err != nil {
		response.Failed(c, err)
		return
	}

	newBlog, err := b.svc.UpdateBlogStatus(c.Request.Context(), newReq)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, newBlog)

}

func (b *BlogApiHandler) CreateBlog(c *gin.Context) {

	tokenObject := c.Keys[token.TOKEN_GIN_KEY_IN_CONTEXT]
	// fmt.Println(tokenObject.(*token.Token).UserId)
	theToken := tokenObject.(*token.Token)
	newReq := blog.NewCreateBlogRequest()
	err := c.BindJSON(newReq)

	if err != nil {
		// c.JSON(http.StatusBadRequest, err.Error())
		response.Failed(c, err)
		return
	}
	// get the author name
	newReq.CreatedBy = theToken.UserName

	newBlog, err := b.svc.CreateBlog(c.Request.Context(), newReq)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//return response
	response.Success(c, newBlog)
	// c.JSON(http.StatusOK, tk)

}

func (b *BlogApiHandler) CreateBlogWithMQ(c *gin.Context) {

	tokenObject := c.Keys[token.TOKEN_GIN_KEY_IN_CONTEXT]
	// fmt.Println(tokenObject.(*token.Token).UserId)
	theToken := tokenObject.(*token.Token)
	newReq := blog.NewCreateBlogRequest()
	err := c.BindJSON(newReq)
	// mqimpl.GetMQClient()

	if err != nil {
		// c.JSON(http.StatusBadRequest, err.Error())
		response.Failed(c, err)
		return
	}
	// get the author name
	newReq.CreatedBy = theToken.UserName

	resultChan := make(chan interface{}, 1)
	defer close(resultChan)

	err = mqimpl.GetMQClient().Publish(c, mq.CREATE_BLOG_QUEUE, newReq, resultChan)

	// newBlog, err := b.svc.CreateBlog(c.Request.Context(), newReq)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createdBlog := <-resultChan
	if createdBlog == nil {
		response.Failed(c, fmt.Errorf("failed to create blog"))
		return
	}
	response.Success(c, createdBlog)

}

func (b *BlogApiHandler) ConsumeCreateBlog() {
	msgs, err := mqimpl.GetMQClient().Consumer(mq.CREATE_BLOG_QUEUE)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			var req blog.CreateBlogRequest
			err := json.Unmarshal(d.Body, &req)
			if err != nil {
				log.Printf("Error decoding JSON: %v", err)
				continue
			}
			if mqimpl.GetMQClient().GetCtx() == nil {
				fmt.Println(mqimpl.GetMQClient().GetCtx())
				d.Ack(false)
				continue
			}

			createdBlog, err := b.svc.CreateBlog(mqimpl.GetMQClient().GetCtx().Request.Context(), &req)
			if err != nil {
				log.Printf("Failed to create blog: %v", err)
				d.Ack(false)
				continue
			}
			d.Ack(false)

			// Retrieve the result channel and send the created blog
			resultChan := mqimpl.GetMQClient().RetrieveResultChannel(mq.CREATE_BLOG_QUEUE)
			resultChan <- createdBlog
		}
	}()
}

func (b *BlogApiHandler) UpdateBlogAll(c *gin.Context) {

	newReq := blog.NewUpdateBlogRequest(c.Param("id"))

	err := c.BindJSON(newReq.CreateBlogRequest)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// err := c.BindJSON(newReq)
	newReq.SetUpdateBlogRequestUpdateMode(blog.UPDATE_MODE_PUT)

	newBlog, err := b.svc.UpdateBlog(c.Request.Context(), newReq)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//return response
	response.Success(c, newBlog)

}

func (b *BlogApiHandler) UpdateBlogPartial(c *gin.Context) {

	newReq := blog.NewUpdateBlogRequest(c.Param("id"))
	newReq.SetUpdateBlogRequestUpdateMode(blog.UPDATE_MODE_PATCH)
	// err := c.BindJSON(newReq)

	err := c.BindJSON(newReq.CreateBlogRequest)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newBlog, err := b.svc.UpdateBlog(c.Request.Context(), newReq)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//return response
	response.Success(c, newBlog)

}

func (b *BlogApiHandler) QueryBlogList(c *gin.Context) {

	// get the token from the context

	// tokenObject := c.Keys[token.TOKEN_GIN_KEY_IN_CONTEXT]
	// fmt.Println(tokenObject.(*token.Token).UserId)
	newReq := blog.NewQueryBlogRequest()
	// err := c.BindJSON(newReq)
	newReq.ParsePageSize(c.Query("page_size"))

	newReq.ParsePageNumber(c.Query("page_number"))
	newReq.Keywords = c.Query("keywords")
	newReq.Author = c.Query("author")

	switch c.Query("status") {
	case "draft":
		newReq.SetStatus(blog.DRAFT)
	case "published":
		newReq.SetStatus(blog.PUBLISHED)
	}

	blogsList, err := b.svc.QueryBlog(c.Request.Context(), newReq)
	if err != nil {

		response.Failed(c, err)
		return
	}
	//return response
	response.Success(c, blogsList)

}

func (b *BlogApiHandler) QueryOneBlog(c *gin.Context) {
	newReq := blog.NewQuerySingleBlogRequest(c.Param("id"))
	// err := c.BindJSON(newReq)

	newBlog, err := b.svc.QuerySingleBlog(c.Request.Context(), newReq)
	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//return response
	response.Success(c, newBlog)

}

func (b *BlogApiHandler) DeleteOneBlog(c *gin.Context) {
	newReq := blog.NewDeleteBlogRequest()

	err := newReq.SetBlogId(c.Param("id"))

	if err != nil {
		response.Failed(c, err)
		return
	}

	err = b.svc.DeleteBlog(c.Request.Context(), newReq)
	if err != nil {
		response.Failed(c, err)
		return
	}

}

// pass blog id in the path, POST .../{id}/audit
func (b *BlogApiHandler) AuditOneBlog(c *gin.Context) {
	newReq := blog.NewAuditBlogRequest(c.Param("id"))
	err := c.BindJSON(newReq)
	if err != nil {
		response.Failed(c, err)
		return
	}

	theBlog, err := b.svc.AuditBlog(c.Request.Context(), newReq)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, theBlog)

}

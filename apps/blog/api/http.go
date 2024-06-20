package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IanZC0der/go-myblog/apps/blog"
	mq "github.com/IanZC0der/go-myblog/apps/mq"
	mqimpl "github.com/IanZC0der/go-myblog/apps/mq/impl"
	"github.com/IanZC0der/go-myblog/ioc"
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
	v1.POST("/", b.CreateBlogWithMQ)
	// v1.POST("/", b.CreateBlog)
	v1.DELETE("/:id", b.DeleteOneBlog)
	v1.PUT("/:id", b.UpdateBlogAll)
	v1.PATCH("/:id", b.UpdateBlogPartial)
	// /myblog/api/v1/blogs
	v1.GET("/", b.QueryBlogList)
	// myblog/api/v1/blogs/:id
	v1.GET("/:id", b.QueryOneBlog)
}

func (b *BlogApiHandler) CreateBlog(c *gin.Context) {
	newReq := blog.NewCreateBlogRequest()
	err := c.BindJSON(newReq)

	if err != nil {
		// c.JSON(http.StatusBadRequest, err.Error())
		response.Failed(c, err)
		return
	}

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
	newReq := blog.NewCreateBlogRequest()
	err := c.BindJSON(newReq)

	if err != nil {
		// c.JSON(http.StatusBadRequest, err.Error())
		response.Failed(c, err)
		return
	}

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

			createdBlog, err := b.svc.CreateBlog(context.Background(), &req)
			if err != nil {
				log.Printf("Failed to create blog: %v", err)
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
	newReq := blog.NewQueryBlogRequest()
	// err := c.BindJSON(newReq)
	err := newReq.ParsePageSize(c.Query("page_size"))

	if err != nil {
		response.Failed(c, err)
		return
	}
	err = newReq.ParsePageNumber(c.Query("page_number"))

	if err != nil {
		response.Failed(c, err)
		return
	}

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

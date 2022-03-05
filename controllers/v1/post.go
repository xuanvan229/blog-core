package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xuanvan229/blog-core/constants"
	"github.com/xuanvan229/blog-core/datatransfers"
	"github.com/xuanvan229/blog-core/handlers"
)

func CreatePost(c *gin.Context) {
	var err error
	var post datatransfers.PostCreate

	if err = c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.CreatePost(c.MustGet(constants.UserIDKey).(uint), post); err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Error: "failed creating post"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Response{Data: "post created"})
}

func GetAll(c *gin.Context) {
	var err error
	var posts []datatransfers.PostInfor
	if posts, err = handlers.Handler.GetAllPost(c.MustGet(constants.UserIDKey).(uint)); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find user"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Response{Data: posts})
}

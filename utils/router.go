package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xuanvan229/blog-core/constants"
	"github.com/xuanvan229/blog-core/datatransfers"
)

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: "user not authenticated"})
	}
}

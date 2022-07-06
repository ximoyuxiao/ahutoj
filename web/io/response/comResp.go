package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"code"`
	StatusMsg  string `json:"msg"`
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		StatusCode: code,
		StatusMsg:  "error",
	})
}

func ResponseOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

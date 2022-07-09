package response

import (
	"ahutoj/web/io/constanct"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode constanct.ResCode `json:"code"`
	StatusMsg  string            `json:"msg"`
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code constanct.ResCode) {
	c.JSON(http.StatusOK, Response{
		StatusCode: code,
		StatusMsg:  code.Msg(),
	})
}

func ResponseOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func CreateResponse(code constanct.ResCode) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  code.Msg(),
	}
}

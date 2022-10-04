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
func ResponseServerError(c *gin.Context, code constanct.ResCode) {
	c.JSON(http.StatusBadGateway, Response{
		StatusCode: code,
		StatusMsg:  code.Msg(),
	})
}

func ResponseError(c *gin.Context, code constanct.ResCode) {
	c.JSON(code.HttpCode(), Response{
		StatusCode: code,
		StatusMsg:  code.Msg(),
	})
}

func ResponseErrorStr(c *gin.Context, code constanct.ResCode, str string) {
	c.JSON(code.HttpCode(), Response{
		StatusCode: code,
		StatusMsg:  str,
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

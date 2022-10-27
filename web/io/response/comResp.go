package response

import (
	"ahutoj/web/io/constanct"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode constanct.ResCode `json:"code"`
	StatusMsg  string            `json:"msg"`
}
type RetType string

const (
	SUCCESS RetType = "success"
	WARNING RetType = "warning"
	INFO    RetType = "info"
	ERROR   RetType = "error"
)

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
		StatusMsg:  "",
	})
}

func ResponseErrorStr(c *gin.Context, code constanct.ResCode, str string, retType RetType) {
	c.JSON(code.HttpCode(), Response{
		StatusCode: code,
		StatusMsg:  fmt.Sprintf("%s\\\\%s", str, retType),
	})
}
func ResponseOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func CreateResponse(code constanct.ResCode) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  "",
	}
}
func CreateResponseStr(code constanct.ResCode, str string, retType RetType) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  fmt.Sprintf("%s\\%s", str, retType),
	}
}

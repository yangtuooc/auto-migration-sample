package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Fail(httpStatus int, message string, c *gin.Context, data ...any) {
	c.JSON(httpStatus, Response{
		Data: data,
		Msg:  message,
	})
}

func Success(c *gin.Context, data ...any) {
	c.JSON(http.StatusOK, Response{
		Data: data,
		Msg:  "success",
	})
}

func SuccessWith(httpStatus int, message string, c *gin.Context, data ...any) {
	c.JSON(httpStatus, Response{
		Data: data,
		Msg:  message,
	})
}

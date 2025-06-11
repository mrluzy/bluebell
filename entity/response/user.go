package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

type nilData map[string]any

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "success", nilData{}, c)
}

func SuccessWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, msg, nilData{}, c)
}

func SuccessWithData(data any, c *gin.Context) {
	Result(SUCCESS, "success", data, c)
}

func SuccessWithDetail(msg string, data any, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

func Failure(c *gin.Context) {
	Result(ERROR, "success", nilData{}, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, msg, nilData{}, c)
}

func FailWithData(data any, c *gin.Context) {
	Result(ERROR, "success", data, c)
}

func FailWithDetail(msg string, data any, c *gin.Context) {
	Result(ERROR, msg, data, c)
}

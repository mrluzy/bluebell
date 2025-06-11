package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrluzy/blueball/entity/request"
	"github.com/mrluzy/blueball/entity/response"
	"github.com/mrluzy/blueball/global"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	// 校验参数
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Error("params of the register error", zap.Error(err))
		response.FailWithMsg(err.Error(), c)
	}

	// 业务处理
	_, err := userService.Register(&req)
	if err != nil {
		global.Logger.Error("Failed to register user:", zap.Error(err))
		response.FailWithMsg("Failed to register user", c)
		return
	}

	// 返回响应
	response.SuccessWithMsg("register success", c)
}

func LoginHandler(c *gin.Context) {
	// 1.检验参数
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Error("params of the login error", zap.Error(err))
		response.FailWithMsg(err.Error(), c)
	}
	// 2.业务处理
	_, err := userService.Login(&req)
	if err != nil {
		global.Logger.Error("failed to login", zap.Error(err))
		response.FailWithMsg(err.Error(), c)
		return
	}
	// 3.返回响应
	response.SuccessWithMsg("login successfully", c)
}

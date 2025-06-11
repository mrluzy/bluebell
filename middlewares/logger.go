package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mrluzy/blueball/global"
	"go.uber.org/zap"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 获取请求的路径和查询参数
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 继续执行后续的处理
		c.Next()

		// 计算请求处理的耗时
		cost := time.Since(start)

		// 使用 Zap 记录请求日志
		global.Logger.Info(path,
			// 记录响应状态码
			zap.Int("status", c.Writer.Status()),
			// 记录请求方法
			zap.String("method", c.Request.Method),
			// 记录请求路径
			zap.String("path", path),
			// 记录查询参数
			zap.String("query", query),
			// 记录客户端 IP
			zap.String("ip", c.ClientIP()),
			// 记录 User-Agent 信息
			zap.String("user-agent", c.Request.UserAgent()),
			// 记录错误信息（如果有）
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			// 记录请求耗时
			zap.Duration("cost", cost),
		)
	}
}

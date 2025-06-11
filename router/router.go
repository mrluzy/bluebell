package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mrluzy/blueball/controllers"
	"github.com/mrluzy/blueball/global"
	"github.com/mrluzy/blueball/middlewares"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func RunHTTPServer() {
	addr := viper.GetString("host")
	if addr == "" {
		global.Logger.Panic("empty http address")
	}
	runHTTPServerOnAddr(addr)
}

func runHTTPServerOnAddr(addr string) {
	apiRouter := gin.New()
	setMiddlewares(apiRouter)

	apiRouter.GET("./", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
	apiRouter.POST("./register", controllers.RegisterHandler)
	apiRouter.POST("./login", controllers.LoginHandler)
	if err := apiRouter.Run(addr); err != nil {
		global.Logger.Panic("failed to run http server", zap.Error(err))
	}
}

func setMiddlewares(r *gin.Engine) {
	r.Use(middlewares.GinLogger())
	r.Use(gin.Recovery())
}

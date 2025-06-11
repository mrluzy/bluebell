package main

import (
	_ "github.com/mrluzy/blueball/config"
	"github.com/mrluzy/blueball/initialize"
	"github.com/mrluzy/blueball/router"
)

func main() {
	// 初始化zap，mysql, redis等配置
	initialize.InitAllConfig()
	// 启动http服务
	router.RunHTTPServer()
}

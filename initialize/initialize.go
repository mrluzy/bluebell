package initialize

import (
	"github.com/mrluzy/blueball/global"
	"github.com/mrluzy/blueball/infrastructure/persistent"
	"github.com/mrluzy/blueball/logger"
	"github.com/mrluzy/blueball/utils/snowflake"
)

func InitAllConfig() {
	global.Logger = logger.Init()
	global.DB = persistent.InitMySQL()
	//global.RedisClient = redis.Init()
	snowflake.Init()
}

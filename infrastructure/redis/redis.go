package redis

import (
	"context"
	"github.com/mrluzy/blueball/global"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		global.Logger.Panic("connect to redis failed", zap.Error(err))
	}
	global.Logger.Info("redis init success")
	return redisClient
}

package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.Logger
	DB          *gorm.DB
	RedisClient *redis.Client
)

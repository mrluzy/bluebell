package persistent

import (
	"fmt"
	"github.com/mrluzy/blueball/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Fatal("connect to mysql failed", zap.Error(err))
	}
	mysqlDb, _ := db.DB()
	mysqlDb.SetMaxIdleConns(11)
	mysqlDb.SetMaxOpenConns(110)

	// 检查数据库连接是否通畅
	if err := mysqlDb.Ping(); err != nil {
		global.Logger.Fatal("ping mysql failed", zap.Error(err))
	}

	global.Logger.Info("mysql init success")

	return db
}

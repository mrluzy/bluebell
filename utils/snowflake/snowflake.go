package snowflake

import (
	"github.com/mrluzy/blueball/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init() {
	st, err := time.Parse("2006-01-02", viper.GetString("startTime"))
	if err != nil {
		global.Logger.Panic("parse time failed", zap.Error(err))
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(viper.GetInt64("machineID"))
	if err != nil {
		global.Logger.Panic("snowflake init failed", zap.Error(err))
	}
	global.Logger.Info("snowflake init success")
}

func GenID() int64 {
	return node.Generate().Int64()
}

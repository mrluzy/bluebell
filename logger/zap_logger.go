package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func Init() *zap.Logger {
	encoder := getEncoder()

	ws := getWriter(
		viper.GetInt("log.maxAge"),
		viper.GetInt("log.maxBackups"),
		viper.GetInt("log.maxSize"),
		viper.GetString("log.filename"),
	)
	if viper.GetBool("log.isConsolePrint") {
		ws = zapcore.NewMultiWriteSyncer(ws, zapcore.AddSync(os.Stdout))
	}

	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(viper.GetString("log.level"))); err != nil {
		fmt.Println("logger error", err)
		log.Fatal(err)
	}
	core := zapcore.NewCore(encoder, ws, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getWriter(maxAge, maxBackups, maxSize int, filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.TimeKey = "time"
	return zapcore.NewJSONEncoder(cfg)
}

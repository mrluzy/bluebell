package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var once sync.Once

func init() {
	if err := InitViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func InitViperConfig() (err error) {
	once.Do(
		func() {
			err = initViperConfig()
		})
	return
}

func initViperConfig() error {
	relPath, err := getRelativePathCaller()
	if err != nil {
		return err
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(relPath)

	return viper.ReadInConfig()
}

func getRelativePathCaller() (relPath string, err error) {
	// 获取当前进程的工作目录
	callerPwd, err := os.Getwd()
	if err != nil {
		return
	}
	// 	获取当前函数所在文件的绝对路径
	_, filename, _, _ := runtime.Caller(0)
	relPath, err = filepath.Rel(callerPwd, filepath.Dir(filename))
	fmt.Printf("caller from %s, here is %s, relpath is %s", callerPwd, filename, relPath)
	return
}

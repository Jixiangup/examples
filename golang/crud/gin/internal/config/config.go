package config

import (
	"github.com/joho/godotenv"
	"jixiangup.me/examples/gin/internal/constants"
	"jixiangup.me/examples/gin/pkg/logger"
	_ "jixiangup.me/examples/gin/pkg/logger"
	"os"
)

func setup() {
	// 设置
	dir, err := os.Getwd()
	if err != nil {
		logger.Log.Error("获取当前目录失败", true, err)
	}
	err = os.Setenv(constants.HOME, dir)
	if err != nil {
		logger.Log.Error(err.Error(), true, err)
	}
}

func init() {
	setup()
	err := godotenv.Load(os.Getenv(constants.HOME) + "/.env")
	if err != nil {
		logger.Log.Error("Error loading .env file: %v", true, err)
	}
}

package config

import (
	"github.com/joho/godotenv"
	"jixiangup.me/examples/gin/internal/constants"
	"log"
	"os"
)

func SetupBootstrap() {
	// 设置
	dir, err := os.Getwd()
	if err != nil {
		log.Println("获取当前目录失败", err)
	}
	err = os.Setenv(constants.HOME, dir)
	if err != nil {
		log.Println(err.Error(), true, err)
	}
	err = godotenv.Load(os.Getenv(constants.HOME) + "/.env")
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err.Error())
	}
}

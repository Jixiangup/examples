package datasource

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jixiangup.me/examples/gin/internal/constants"
	"jixiangup.me/examples/gin/pkg/logger"
	"os"
	"sync"
)

// MySQLInstance 是数据库连接的单例实例
var MySQLInstance *gorm.DB

// lock 用于确保数据库连接的线程安全
var lock = &sync.Mutex{}

func init() {
	if MySQLInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if MySQLInstance == nil {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", os.Getenv(constants.DbUser), os.Getenv(constants.DbPassword), os.Getenv(constants.DbHost), os.Getenv(constants.DbPort), os.Getenv(constants.DbName))
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				logger.Log.Fatal("初始化MySQL链接失败: %s", dsn, err)
			}
			MySQLInstance = db
		}
	}
}

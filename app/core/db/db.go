package db

import (
	"fmt"
	"gaoMall/app/core/conf"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var dbInstance = make(map[string]*gorm.DB)
var dbLocker sync.RWMutex

func Connect(conf conf.PGSqlConf, confLog conf.PGSqlLog, key string) *gorm.DB {
	dbLocker.RLock()
	db, ok := dbInstance[key]
	if ok {
		dbLocker.RUnlock()
		return db
	}
	dbLocker.RUnlock()

	// TODO: 优化锁的开销 原子操作 atomic

	dbLocker.Lock()
	defer dbLocker.Unlock()
	if _, exist := dbInstance[key]; exist {
		return dbInstance[key]
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host, conf.Username, conf.Password, conf.Database, conf.Port,
	)

	var l logger.Interface
	if confLog.Enable {
		var level logger.LogLevel
		switch confLog.Level {
		case "silent":
			level = logger.Silent
		case "error":
			level = logger.Error
		case "info":
			level = logger.Info
		default:
			level = logger.Warn
		}

		file, err := os.OpenFile(confLog.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed to open gorm.log: %v", err)
		}

		l = logger.New(
			log.New(file, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      level,
				Colorful:      true,
			},
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		panic(err)
	}

	// TODO: 数据库池配置

	dbInstance[key] = db
	return db
}

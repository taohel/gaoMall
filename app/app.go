package app

import (
	"fmt"
	"gaoMall/app/core/conf"
	"gaoMall/app/core/db"
	"gaoMall/app/core/rdb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Config conf.Model

func Init() error {
	return Config.SetUp()
}

func DBW(keys ...string) *gorm.DB {
	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}

	config, ok := Config.PGSql[k]
	if !ok {
		panic(fmt.Sprintf("db config %s not found", k))
	}

	cacheKey := fmt.Sprintf("%s_write", k)
	return db.Connect(config.Write, config.Log, cacheKey)
}

func DBR(keys ...string) *gorm.DB {
	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}

	config, ok := Config.PGSql[k]
	if !ok {
		panic(fmt.Sprintf("db config %s not found", k))
	}

	cacheKey := fmt.Sprintf("%s_read", k)
	return db.Connect(config.Read, config.Log, cacheKey)
}

func RedisR(key ...string) *redis.Client {
	return rdb.Connect("read", &Config, key...)
}

func RedisW(key ...string) *redis.Client {
	return rdb.Connect("write", &Config, key...)
}

package rdb

import (
	"fmt"
	"gaoMall/app/core/conf"
	"github.com/redis/go-redis/v9"
	"sync"
)

var redisClientSyncMap map[string]*redis.Client
var redisClientMutex sync.RWMutex

func Connect(flag string, config *conf.Model, key ...string) *redis.Client {
	var k = "default"
	if len(key) != 0 {
		k = key[0]
	}
	cacheKey := flag + k

	redisClientMutex.RLock()
	if c, ok := redisClientSyncMap[cacheKey]; ok {
		redisClientMutex.RUnlock()
		return c
	}
	redisClientMutex.RUnlock()

	redisClientMutex.Lock()
	defer redisClientMutex.Unlock()
	if c, ok := redisClientSyncMap[cacheKey]; ok {
		return c
	}

	rdsItem, ok := config.Redis[k]
	if !ok {
		panic(fmt.Sprintf("redis %s config not exist, please check your yaml config!", key[0]))
	}

	var rdsConf conf.RedisConf
	if flag == "write" {
		rdsConf = rdsItem.Write
	} else {
		rdsConf = rdsItem.Read
	}

	return redis.NewClient(&redis.Options{
		Addr:     rdsConf.Addr,
		Password: rdsConf.Password,
		DB:       rdsConf.DB,
	})
}

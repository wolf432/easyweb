package initialize

import (
	"easyweb/global"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     global.Cfg.Redis.Addr,
		Password: global.Cfg.Redis.Password,
		DB:       global.Cfg.Redis.DB,
	})
}

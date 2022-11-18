package global

import (
	"easyweb/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Cfg   config.ServerConfig
	DB    *gorm.DB
	REdis *redis.Client
	Log   *zap.Logger
)

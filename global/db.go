package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	RDB      *redis.Client
)

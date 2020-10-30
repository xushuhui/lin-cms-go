package cache

import (
	"github.com/go-redis/redis"
	"lin-cms-go/global"
	"lin-cms-go/pkg/setting"
)

// 初始化连接

func NewRedisClient(redisSet *setting.RedisSettings) (*redis.Client, error) {
	RDB := redis.NewClient(&redis.Options{
		Addr:     redisSet.Host,
		Password: redisSet.Password, // no password set
		DB:       0,                 // use default DB
	})
	return RDB, nil
}

func SelectDB(db int) {
	global.RDB.Do("SELECT", db)
}

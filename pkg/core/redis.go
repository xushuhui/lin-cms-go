package core

import (
	"lin-cms-go/global"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var RDB *redis.Client

// 初始化连接

func NewClient() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host,
		Password: global.RedisSetting.Password, // no password set
		DB:       0,                            // use default DB
	})
}
func SelectDB(db int) {
	RDB.Do("SELECT", db)
}

package initialize

import (
	"fmt"
	"gincrm/global"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	REDIS := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%v", global.Config.Redis.Host, global.Config.Redis.Port),
		Password:     global.Config.Redis.Password,
		DB:           global.Config.Redis.DB,
		PoolSize:     global.Config.Redis.PoolSize,
		MinIdleConns: global.Config.Redis.MinIdleConn,
	})

	global.REDIS = REDIS
}

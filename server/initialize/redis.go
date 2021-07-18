package initialize

import (
	"github.com/techidea8/restgo/redis"
	"turingapp/global"
)

//注册控制器
func InitRedis() *redis.RedisClient {
	global.Redis = redis.InitRedis(global.Conf.Redis)
	return global.Redis
}

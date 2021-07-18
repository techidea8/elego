package global

import (
	"github.com/spf13/viper"
	redis "github.com/techidea8/restgo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"turingapp/config"
)

var (
	DbEngin *gorm.DB
	Redis   *redis.RedisClient
	Log     *zap.Logger
	Viper   *viper.Viper
	Conf    config.Conf
)

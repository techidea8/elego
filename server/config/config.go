package config

import (
	"github.com/techidea8/restgo/database"
	"github.com/techidea8/restgo/redis"
)

type Conf struct {
	Redis redis.Conf `mapstructure:"redis" json:"redis" yaml:"redis"`

	System System `mapstructure:"system" json:"system" yaml:"system"`

	Log Log `mapstructure:"log" json:"log" yaml:"log"`

	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	Aliyun Aliyun `mapstructure:"aliyun" json:"aliyun" yaml:"aliyun"`

	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	Sms Sms `mapstructure:"sms" json:"sms" yaml:"sms"`
	// gorm
	Database database.Conf `mapstructure:"database" json:"database" yaml:"database"`
}

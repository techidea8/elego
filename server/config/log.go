package config

import "go.uber.org/zap/zapcore"

type Log struct {
	Filename    string        `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxSize     int           `mapstructure:"maxsize" json:"maxsize" yaml:"maxsize"`
	MaxBackups  int           `mapstructure:"maxbackups" json:"maxbackups" yaml:"maxbackups"`
	MaxAge      int           `mapstructure:"maxage" json:"maxage" yaml:"maxage"`
	Compress    bool          `mapstructure:"compress" json:"compress" yaml:"compress"`
	Level       zapcore.Level `mapstructure:"level" json:"level" yaml:"level"`
	Formate     string        `mapstructure:"formate" json:"formate" yaml:"formate"`
	EncodeLevel string        `mapstructure:"encodelevel" json:"encodelevel" yaml:"encodelevel"`
}

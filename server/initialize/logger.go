package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
	"turingapp/config"
	"turingapp/global"
)

func Logger(conf config.Log) (logger *zap.Logger) {

	hook := lumberjack.Logger{
		Filename:   conf.Filename,   // 日志文件路径
		MaxSize:    conf.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: conf.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     conf.MaxAge,     // 文件最多保存多少天
		Compress:   conf.Compress,   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     CustomTimeEncoder,              // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	switch {
	case conf.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	case conf.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case conf.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case conf.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(conf.Level)

	//设置日志格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if conf.Formate == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder, // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger = zap.New(core, caller, development)
	return logger

}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
func InitLogger() *zap.Logger {
	global.Log = Logger(global.Conf.Log)
	return global.Log
}

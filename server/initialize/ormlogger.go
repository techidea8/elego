package initialize

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type ormlogger struct {
	logger.Config
	log *zap.Logger
}
func New(conf logger.Config,log *zap.Logger) *ormlogger{
	return &ormlogger{
		conf,log,
	}
}
// LogMode log mode
func (l *ormlogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l ormlogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.log.Info(msg, zap.Any("",append([]interface{}{utils.FileWithLineNum()}, data...)))
	}
}

// Warn print warn messages
func (l ormlogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.log.Warn(msg, zap.Any("",append([]interface{}{utils.FileWithLineNum()}, data...)))
	}
}

// Error print error messages
func (l ormlogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.log.Error(msg, zap.Any("",append([]interface{}{utils.FileWithLineNum()}, data...)))
	}
}

// Trace print sql message
func (l ormlogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) ):
			sql, rows := fc()
			if rows == -1 {
				l.log.Error(utils.FileWithLineNum(),zap.Error(err),zap.Float64("time",float64(elapsed.Nanoseconds())/1e6),zap.String("rows","-"), zap.String("sql",sql))
			} else {
				l.log.Error(utils.FileWithLineNum(),zap.Error(err),zap.Float64("time",float64(elapsed.Nanoseconds())/1e6),zap.Int64("rows",rows), zap.String("sql",sql))
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >=  logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				l.log.Warn(utils.FileWithLineNum(),zap.String("slowsql",slowLog),zap.Float64("@",float64(elapsed.Nanoseconds())/1e6),zap.String("rows","-"), zap.String("sql",sql))
			} else {
				l.log.Warn(utils.FileWithLineNum(),zap.String("slowsql",slowLog),zap.Float64("@",float64(elapsed.Nanoseconds())/1e6),zap.Int64("rows",rows), zap.String("sql",sql))
			}
		case l.LogLevel ==  logger.Info:
			sql, rows := fc()
			if rows == -1 {
				l.log.Info(utils.FileWithLineNum(),zap.Float64("time",float64(elapsed.Nanoseconds())/1e6),zap.String("rows","-"), zap.String("sql",sql))
			} else {
				l.log.Info(utils.FileWithLineNum(),zap.Float64("time",float64(elapsed.Nanoseconds())/1e6),zap.Int64("rows",rows), zap.String("sql",sql))
			}
		}
	}
}
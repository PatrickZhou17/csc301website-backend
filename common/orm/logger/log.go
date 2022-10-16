package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"shopping-cart/common/global"
	"shopping-cart/common/log"
)

type XZLogBridge struct {
	xzlog         *log.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func NewXZLogBridge() logger.Interface {
	return &XZLogBridge{
		xzlog: log.StdLogger(),
	}
}

func (l *XZLogBridge) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.xzlog = log.StdLogger()
	newLogger.LogLevel = level
	newLogger.SlowThreshold = 200 * time.Millisecond
	return &newLogger
}

func (l *XZLogBridge) Info(ctx context.Context, msg string, data ...interface{}) {
	l.xzlog.Infof(msg, data)
}

func (l *XZLogBridge) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.xzlog.Warnf(msg, data)
}

func (l *XZLogBridge) Error(ctx context.Context, msg string, data ...interface{}) {
	l.xzlog.Errorf(msg, data)
}

func (l *XZLogBridge) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if os.Getenv(global.APPEnv) == global.DevEnv {
		if l.LogLevel > logger.Silent {
			elapsed := time.Since(begin)
			switch {
			case err != nil && l.LogLevel >= logger.Error:
				sql, rows := fc()
				if rows == -1 {
					l.xzlog.Debugf("[file:%s] [err:%s] [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				} else {
					l.xzlog.Debugf("[file:%s] [err:%s] [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				}
			case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
				sql, rows := fc()
				slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
				if rows == -1 {
					l.xzlog.Debugf("[file:%s] [err:%s] [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				} else {
					l.xzlog.Debugf("[file:%s] [err:%s] [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				}
			case l.LogLevel == logger.Info:
				sql, rows := fc()
				if rows == -1 {
					l.xzlog.Debugf("[file:%s] [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
				} else {
					l.xzlog.Debugf("[file:%s] [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
				}
			}
		}
	}
}

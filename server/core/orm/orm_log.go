package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/core/trace"
	"github.com/limeschool/easy-admin/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

// zap 适配gorm 日志
type sqlLog struct {
	logger        *zap.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func newOrmLog(conf config.Orm) logger.Interface {
	return &sqlLog{
		logger:        global.Logger.WithOptions(zap.AddCallerSkip(3)),
		LogLevel:      logger.LogLevel(conf.Level),
		SlowThreshold: conf.SlowThreshold,
	}
}

func (l *sqlLog) Log(ctx context.Context) *zap.Logger {
	traceId, _ := ctx.Value(trace.Key).(string)
	return l.logger.With(zap.Any(trace.Key, traceId))
}

// LogMode sqlLog mode
func (l *sqlLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

// Info print info
func (l sqlLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Log(ctx).Info("SQL信息", getSqlInfo("", fmt.Sprintf(msg, data...), 0, 0, false)...)
	}
}

// Warn print warn messages
func (l sqlLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Log(ctx).Info("SQL告警", getSqlInfo("", fmt.Sprintf(msg, data...), 0, 0, false)...)
	}
}

// Error print error messages
func (l sqlLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Log(ctx).Info("SQL错误", getSqlInfo("", fmt.Sprintf(msg, data...), 0, 0, false)...)
	}
}

// Trace print sql message
func (l sqlLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	costTime := float64(elapsed.Nanoseconds()) / 1e6
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		sql, rows := fc()
		l.Log(ctx).Info("SQL错误", getSqlInfo(err.Error(), sql, rows, costTime, false)...)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.Log(ctx).Info("SQL告警", getSqlInfo("", sql, rows, costTime, true)...)
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		l.Log(ctx).Info("SQL信息", getSqlInfo("", sql, rows, costTime, false)...)
	}
}

// 获取sql的执行信息
func getSqlInfo(err, sql string, rows int64, costTime float64, slow bool) []zap.Field {
	return []zap.Field{
		zap.String("err", err),
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.String("time", fmt.Sprintf("%vms", costTime)),
		zap.Bool("slow", slow),
	}
}

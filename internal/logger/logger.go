package logger

import (
	"github.com/costtinha/first-golang-rest-api/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(cfg *config.Config) *Logger {
	level := zapcore.InfoLevel
	if cfg.AppEnv == "dev" {
		level = zapcore.DebugLevel
	}
	cfgZap := zap.NewDevelopmentConfig()
	cfgZap.Level = zap.NewAtomicLevelAt(level)
	z, _ := cfgZap.Build()
	return &Logger{z.Sugar()}
}

func (l *Logger) Sync() { _ = l.SugaredLogger.Sync() }

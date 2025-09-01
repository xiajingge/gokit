package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestLoggerV1(t *testing.T) {
	cfg := zap.NewDevelopmentConfig()
	logger, err := cfg.Build()
	if err != nil {
		return
	}
	l1 := NewZapLogger(logger)
	l1.Info("这是一条info信息", String("info", ""))
}

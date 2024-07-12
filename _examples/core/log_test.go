package core

import (
	"errors"
	"github.com/xingcxb/goKit/core/logKit"
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {
	l := logKit.LogConfig{
		LogFilePath: "log/log.log",
		LogLevel:    "error,info,debug",
		ShowConsole: true,
		MaxSize:     10,
		MaxBackups:  50,
		MaxAge:      30,
		Compress:    true,
		LocalTime:   true,
	}
	l.InitGlobalLogger()
	// 添加测试日志
	zap.L().Info("这是一个测试信息日志")
	zap.L().Debug("这是一个测试调试日志")
	err := errors.New("测试")
	zap.L().Error("这是一个测试错误日志", zap.Error(err))
}

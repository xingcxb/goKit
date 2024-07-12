package logKit

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

var (
	once sync.Once
)

// LogConfig 日志配置文件
type LogConfig struct {
	LogFilePath string // 日志文件路径
	LogLevel    string // 日志级别
	ShowConsole bool   // 是否显示控制台日志
	MaxSize     int    // 日志文件大小，单位M
	MaxBackups  int    // 最大备份个数
	MaxAge      int    // 最大备份天数
	Compress    bool   // 是否压缩
	LocalTime   bool   // 是否使用本地时间
}

// InitGlobalLogger 初始化全局日志实例
func (l *LogConfig) InitGlobalLogger() {
	once.Do(func() {
		logger := l.initLogger()
		zap.ReplaceGlobals(logger)
	})
}

// InitLog 初始化日志
func (l *LogConfig) initLogger() *zap.Logger {
	// 配置日志编码器
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 解析配置的日志级别（仅用于文件输出）
	levels := parseLogLevels(l.LogLevel)
	// 创建文件输出核心
	fileSyncer := zapcore.AddSync(l.openLogFile())
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileSyncer,
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return levels[lvl]
		}),
	)
	// 创建核心列表
	cores := []zapcore.Core{fileCore}
	// 如果配置允许，添加控制台输出核心（输出所有级别的日志）
	if l.ShowConsole {
		consoleSyncer := zapcore.AddSync(os.Stdout)
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			consoleSyncer,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return true // 允许所有级别的日志
			}),
		)
		cores = append(cores, consoleCore)
	}
	// 创建多核心日志记录器
	core := zapcore.NewTee(cores...)
	// 构建日志记录器
	return zap.New(core, zap.AddCaller())
}

// 打开日志文件
func (l *LogConfig) openLogFile() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   l.LogFilePath,
		MaxSize:    l.MaxSize,
		MaxBackups: l.MaxBackups,
		MaxAge:     l.MaxAge,
		Compress:   l.Compress,
		LocalTime:  l.LocalTime,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func parseLogLevels(levelString string) map[zapcore.Level]bool {
	levels := make(map[zapcore.Level]bool)
	for _, level := range strings.Split(levelString, ",") {
		switch strings.ToLower(strings.TrimSpace(level)) {
		case "debug":
			levels[zapcore.DebugLevel] = true
		case "info":
			levels[zapcore.InfoLevel] = true
		case "warn":
			levels[zapcore.WarnLevel] = true
		case "error":
			levels[zapcore.ErrorLevel] = true
		}
	}
	return levels
}

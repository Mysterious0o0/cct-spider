package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	logConsole = true                 // 是否打印日志到终端
	JsonFormat = true                 // 日志保存格式是否采用json类型
	Filename   = "log/cct-spider.log" // 日志保存路径
	MaxSize    = 100                  // 每个日志文件保存几M，默认100M
	MaxBackups = 30                   // 保留多少个备份，默认不限
	MaxAge     = 0                    // 保留多少天，默认不限
	Compress   = false                // 是否压缩，默认不压缩
)

var (
	logger       *zap.Logger
	defaultLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
)

var levelPool = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func init() {
	writer := writerSyncer()
	var syncer zapcore.WriteSyncer
	if logConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer)
	} else {
		syncer = zapcore.AddSync(writer)
	}

	config := encoderConfig()
	var encoder zapcore.Encoder
	if JsonFormat {
		encoder = zapcore.NewJSONEncoder(config)
	} else {
		encoder = zapcore.NewConsoleEncoder(config)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		defaultLevel,
	)

	logger = zap.New(core).WithOptions(
		zap.AddCaller(),
		zap.Development(),
		zap.AddCallerSkip(1),
	)
}

func writerSyncer() zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename: Filename,
		Compress: Compress,
	}
	if MaxBackups > 0 {
		hook.MaxBackups = MaxBackups
	}
	if MaxSize > 0 {
		hook.MaxSize = MaxSize
	}
	if MaxAge > 0 {
		hook.MaxAge = MaxAge
	}
	return zapcore.AddSync(hook)
}

func encoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	return config
}

func SetLevel(level string) {
	if Level, ok := levelPool[level]; ok {
		defaultLevel.SetLevel(Level)
		return
	}
	defaultLevel.SetLevel(zapcore.InfoLevel)
}

func Field(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instLog   *Logger
	zapLogger *zap.Logger
	zapConfig zap.Config
)

type Logger struct {
	ILogger
}

func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05")))
}

func encodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func InitLog(workingDirPath string) {
	var err error
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          &zap.SamplingConfig{},
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "debug",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			TimeKey:       "timestamp",
			EncodeTime:    zapcore.EpochMillisTimeEncoder,
			StacktraceKey: "stack",
			LineEnding:    "\n",
		},
		OutputPaths:      []string{"stdout", path.Join(workingDirPath, "bin", "logs", "log.log")},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{},
	}

	cfg.EncoderConfig.EncodeTime = encodeTime
	cfg.EncoderConfig.EncodeLevel = encodeLevel
	zapConfig = cfg
	zapLogger, err = zapConfig.Build(zap.WrapCore(zapCore))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer zapLogger.Sync()
}

func zapCore(c zapcore.Core) zapcore.Core {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   zapConfig.OutputPaths[1],
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     100, //days
		Compress:   true,
	})

	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
		w,
		zap.DebugLevel,
	)

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = encodeTime
	pe.EncodeLevel = encodeLevel
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	cores := zapcore.NewTee(c, fileCore, consoleCore)

	return cores
}

func Inst() ILogger {
	if instLog == nil {
		instLog = &Logger{}
	}
	return instLog
}

func (l *Logger) Info(message string) {
	zapLogger.Info(message)
}

func (l *Logger) Error(message string) {
	zapLogger.Error(message)
}

package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"opspaas/pkg/config"
	"os"
	"path/filepath"
	"strconv"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("get current dir failed")
		return ""
	}
	return dir
}

func GetLogConfig() (logPath ,logName string ,logDebug bool) {
	logPath = config.GetString("LogPath")
	if logPath == ""{
		dir := getCurrentDirectory()
		logPath = dir
	}
	logName = config.GetString("LogName")
	Debug := config.GetString("LogDebug")
	logDebug, _ = strconv.ParseBool(Debug)
	return logPath,logName,logDebug
}

func InitLogger() *zap.Logger {
	logPath, logName, logDebug := GetLogConfig()
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    128,
		MaxAge:     7,
		MaxBackups: 30,
		Compress:   false,
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	//如果是开发环境，同时在控制台输出
	if logDebug {
		writes =append(writes,zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)
	//开启开发模式。堆栈跟踪
	caller := zap.AddCaller()
	//开启文件以及行号
	development := zap.Development()
	//设置初始化字段
	field := zap.Fields(zap.String("logName", logName))
	//构造日志
	logger := zap.New(core, caller, development, field)

	return logger
}

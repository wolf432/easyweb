package initialize

import (
	"easyweb/global"
	"easyweb/pkg/directory"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitLog() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	options = append(options, zap.AddCaller())

	// 初始化 zap
	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := directory.PathExists(global.Cfg.Zap.Director); !ok {
		_ = os.Mkdir(global.Cfg.Zap.Director, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.Cfg.Zap.Level {
	case "debug":
		level = zap.DebugLevel
		//options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.Cfg.Zap.Level + "." + l.String())
	}

	// 设置编码器
	if global.Cfg.Zap.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	//设置写入日志的文件
	logF, err := os.OpenFile(global.Cfg.Zap.Director+"/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("创建info.log日志失败。%v", err))
	}
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), global.Cfg.Zap.GetLevel())
	// 记录ERROR级别的日志
	errF, err := os.OpenFile(global.Cfg.Zap.Director+"/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("error.log日志失败。%v", err))
	}
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	// 如果是开发环境打印到控制台
	var core zapcore.Core
	if global.Cfg.App.Env == "dev" {
		conF := os.Stdout
		c3 := zapcore.NewCore(encoder, zapcore.AddSync(conF), zap.DebugLevel)
		core = zapcore.NewTee(c1, c2, c3)
	} else {
		core = zapcore.NewTee(c1, c2)
	}
	return core
}

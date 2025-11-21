package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() error {
	config := zap.NewProductionConfig()
	//设置日志级别
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	//设置日志输出文件

	config.OutputPaths = []string{
		"logs/app.log",
		"stdout", //
	}

	Logger, err := config.Build()
	if err != nil {
		return err
	}
	Logger.Info("success")
	return nil

}

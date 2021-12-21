package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger        *ZapOpaLogger
	defaultConfig = zap.NewDevelopmentConfig()
	// defaultConfig = zapdriver.NewDevelopmentConfig()
)

func init() {
	config := defaultConfig
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zapper, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = &ZapOpaLogger{
		config:        config,
		SugaredLogger: zapper.Sugar(),
	}
}

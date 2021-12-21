package pkg

import (
	"github.com/open-policy-agent/opa/logging"
	"go.uber.org/zap"
)

type ZapOpaLogger struct {
	config zap.Config
	*zap.SugaredLogger
}

func (z ZapOpaLogger) Debug(fmt string, a ...interface{}) {
	z.Debugw(fmt, a)
}

func (z ZapOpaLogger) Info(fmt string, a ...interface{}) {
	z.Infow(fmt, a)
}

func (z ZapOpaLogger) Error(fmt string, a ...interface{}) {
	z.Errorw(fmt, a)
}

func (z ZapOpaLogger) Warn(fmt string, a ...interface{}) {
	z.Warnw(fmt, a)
}

func (z ZapOpaLogger) WithFields(m map[string]interface{}) logging.Logger {
	return ZapOpaLogger{
		config:        z.config,
		SugaredLogger: z.With(m),
	}
}

func (z ZapOpaLogger) GetFields() map[string]interface{} {
	return map[string]interface{}{}
}

func (z ZapOpaLogger) GetLevel() logging.Level {
	switch defaultConfig.Level.Level() {
	case zap.DebugLevel:
		return logging.Debug
	case zap.InfoLevel:
		return logging.Info
	case zap.WarnLevel:
		return logging.Warn
	default:
		return logging.Error
	}
}

func (z ZapOpaLogger) SetLevel(level logging.Level) {
	config := defaultConfig

	switch level {
	case logging.Debug:
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case logging.Info:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case logging.Warn:
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	setGlobalLogger(config)
}

var (
	Logger        *ZapOpaLogger
	defaultConfig = zap.NewDevelopmentConfig()
)

func init() {
	setGlobalLogger(defaultConfig)
}

func setGlobalLogger(config zap.Config) {
	zapper, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = &ZapOpaLogger{
		config:        config,
		SugaredLogger: zapper.Sugar(),
	}
}

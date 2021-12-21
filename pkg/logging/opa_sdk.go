package logging

import (
	"github.com/open-policy-agent/opa/logging"
	"go.uber.org/zap"
)

func (z *ZapOpaLogger) Debug(format string, a ...interface{}) {
	z.initSkipped()
	z.withSkip.Debugf(format, a...)
}

func (z *ZapOpaLogger) Info(format string, a ...interface{}) {
	z.initSkipped()
	z.withSkip.Infof(format, a...)
}

func (z *ZapOpaLogger) Error(format string, a ...interface{}) {
	z.initSkipped()
	z.withSkip.Errorf(format, a...)
}

func (z *ZapOpaLogger) Warn(format string, a ...interface{}) {
	z.initSkipped()
	z.withSkip.Warnf(format, a...)
}

func (z *ZapOpaLogger) WithFields(m map[string]interface{}) logging.Logger {
	z.initSkipped()

	var newContext []interface{}
	for k, v := range m {
		newContext = append(newContext, k, v)
	}

	newZ := &ZapOpaLogger{
		config:        z.config,
		SugaredLogger: z.With(newContext...),
	}
	newZ.initSkipped()
	return newZ
}

func (z *ZapOpaLogger) GetFields() map[string]interface{} {
	return map[string]interface{}{}
}

func (z *ZapOpaLogger) GetLevel() logging.Level {
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

func (z *ZapOpaLogger) SetLevel(level logging.Level) {
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

	logger, err := config.Build()
	if err != nil {
		z.SugaredLogger.Panic(err)
	}
	*z = ZapOpaLogger{
		config:        config,
		SugaredLogger: logger.Sugar(),
	}
}

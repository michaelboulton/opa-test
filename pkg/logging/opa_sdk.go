package logging

import (
	"github.com/open-policy-agent/opa/logging"
	"go.uber.org/zap"
)

func (z *ZapOpaLogger) Debug(format string, a ...interface{}) {
	z.SugaredLogger.Debugf(format, a...)
}

func (z *ZapOpaLogger) Info(format string, a ...interface{}) {
	z.SugaredLogger.Infof(format, a...)
}

func (z *ZapOpaLogger) Error(format string, a ...interface{}) {
	z.SugaredLogger.Errorf(format, a...)
}

func (z *ZapOpaLogger) Warn(format string, a ...interface{}) {
	z.SugaredLogger.Warnf(format, a...)
}

func (z *ZapOpaLogger) WithFields(m map[string]interface{}) logging.Logger {
	var newContext []interface{}
	for k, v := range m {
		newContext = append(newContext, k, v)
	}
	newContext = append(z.context, newContext...)

	build, err := z.config.Build(zap.AddCallerSkip(1))
	if err != nil {
		z.Panic(err)
	}

	newZ := &ZapOpaLogger{
		config:        z.config,
		SugaredLogger: build.Sugar().With(newContext...),
		context:       newContext,
	}
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

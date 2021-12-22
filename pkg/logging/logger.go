package logging

import (
	"go.uber.org/zap"
)

type ZapOpaLogger struct {
	config zap.Config
	*zap.SugaredLogger
}

func (z *ZapOpaLogger) WithSkip(skip int) *ZapOpaLogger {
	newLogger, err := z.config.Build(zap.AddCallerSkip(1))
	if err != nil {
		z.Panic(err)
	}

	newZ := &ZapOpaLogger{
		config:        z.config,
		SugaredLogger: newLogger.Sugar(),
	}

	return newZ
}

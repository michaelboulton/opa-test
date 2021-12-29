package logging

import (
	"go.uber.org/zap"
)

// ZapOpaLogger is a wrapper for a Zap logger that implements the OPA logger interface
type ZapOpaLogger struct {
	config zap.Config
	*zap.SugaredLogger
	context []interface{}
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

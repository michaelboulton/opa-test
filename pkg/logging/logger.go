package logging

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
)

type ZapOpaLogger struct {
	config   zap.Config
	withSkip *zap.SugaredLogger
	*zap.SugaredLogger
	once sync.Once
}

func (z *ZapOpaLogger) initSkipped() {
	z.once.Do(func() {
		logger, err := z.config.Build(zap.AddCallerSkip(1))
		if err != nil {
			z.SugaredLogger.Panic(err)
		}

		fmt.Fprintln(os.Stderr, "reinintng")
		z.withSkip = logger.Sugar()
	})
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
	newZ.initSkipped()

	return newZ
}

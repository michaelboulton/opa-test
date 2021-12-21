package logging

import "go.uber.org/zap/zapcore"

// logFunc logs some message
type logFunc func(fmt string, a ...interface{})

type LoggerAtLevel struct {
	logFunc logFunc
}

func (l *LoggerAtLevel) Write(p []byte) (n int, err error) {
	l.logFunc(string(p))
	return len(p), nil
}

func (z *ZapOpaLogger) WriteAtLevel(level zapcore.Level) *LoggerAtLevel {
	var logFunc logFunc
	switch level {
	case zapcore.DebugLevel:
		logFunc = z.Debugw
	case zapcore.InfoLevel:
		logFunc = z.Infow
	case zapcore.ErrorLevel:
		logFunc = z.Errorw
	default:
		panic("bad level")
	}

	return &LoggerAtLevel{logFunc}
}

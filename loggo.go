package loggo

import log "github.com/inconshreveable/log15"

type (
	Logger interface {
		New(ctx ...interface{}) Logger
		Debug(msg string, ctx ...interface{})
		Info(msg string, ctx ...interface{})
		Warn(msg string, ctx ...interface{})
		Error(msg string, ctx ...interface{})
		Crit(msg string, ctx ...interface{})
	}

	logger struct {
		logger log.Logger
	}
)

func New(ctx ...interface{}) Logger {
	return &logger{
		logger: log.New(ctx...),
	}
}

func (l *logger) New(ctx ...interface{}) Logger {
	child := l.logger.New(ctx...)
	return &logger{logger: child}
}

func (l *logger) Info(msg string, ctx ...interface{}) {
	l.logger.Info(msg, ctx...)
}

func (l *logger) Debug(msg string, ctx ...interface{}) {
	l.logger.Debug(msg, ctx...)
}

func (l *logger) Warn(msg string, ctx ...interface{}) {
	l.logger.Warn(msg, ctx...)
}

func (l *logger) Error(msg string, ctx ...interface{}) {
	l.logger.Error(msg, ctx...)
}

func (l *logger) Crit(msg string, ctx ...interface{}) {
	l.logger.Crit(msg, ctx...)
}

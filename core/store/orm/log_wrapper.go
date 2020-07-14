package orm

import (
	logpkg "github.com/smartcontractkit/chainlink/core/logger"
	"go.uber.org/zap"
)

var logger *ormLogWrapper

func init() {
	logger = newOrmLogWrapper(logpkg.Default)
}

type ormLogWrapper struct {
	*logpkg.Logger
}

func newOrmLogWrapper(logger *logpkg.Logger) *ormLogWrapper {
	newLogger := logger.
		SugaredLogger.
		Desugar().
		WithOptions(zap.AddCaller(), zap.AddCallerSkip(6)).
		Sugar()
	return &ormLogWrapper{
		Logger: &logpkg.Logger{
			SugaredLogger: newLogger,
		},
	}
}

func (l *ormLogWrapper) Print(args ...interface{}) {
	switch args[0] {
	case "error":
		l.Error(args[2])
	case "log":
		l.Warn(args[2])
	case "sql":
		l.Debugw(args[3].(string), "time", args[2], "rows_affected", args[5])
	default:
		// Don't log these, only seems to be the callback logs which aren't super useful
	}
}

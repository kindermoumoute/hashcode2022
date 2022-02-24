package logger

import (
	"go.uber.org/zap"
)

var L = createLogger()

func createLogger() *zap.SugaredLogger {
	devLog, _ := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}.Build()
	return devLog.Sugar()
}

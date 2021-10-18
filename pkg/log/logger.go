package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Defaultlogger *zap.Logger

func init() {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.00")

	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig

	Defaultlogger, _ = config.Build()

}

package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(isProduction bool) {
	var err error
	if isProduction {
		Log, err = zap.NewProduction()
	} else {
		Log, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("failed to initialize logger" + err.Error())
	}
}

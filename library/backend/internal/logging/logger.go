package logging

import (
	"github.com/LouisBrunner/lemmy/library/backend/internal/common"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

// TODO: FINISH HERE
func NewLogger(dir string) *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	// log.SetFormatter()

	err := common.MkdirIfDoNotExist(dir)
	if err != nil {
		log.WithError(err).Error("failed to create logging directory")
	} else {
		// log.SetOutput()
	}
	return log
}

func NewAdapter(log *logrus.Logger) logger.Logger {
	return &logrusAdapter{logger: log}
}

type logrusAdapter struct {
	logger *logrus.Logger
}

func (me *logrusAdapter) Print(message string) {
	me.logger.Print(message)
}

func (me *logrusAdapter) Trace(message string) {
	me.logger.Trace(message)
}

func (me *logrusAdapter) Debug(message string) {
	me.logger.Debug(message)
}

func (me *logrusAdapter) Info(message string) {
	me.logger.Info(message)
}
func (me *logrusAdapter) Warning(message string) {
	me.logger.Warning(message)
}

func (me *logrusAdapter) Error(message string) {
	me.logger.Error(message)
}

func (me *logrusAdapter) Fatal(message string) {
	me.logger.Fatal(message)
}

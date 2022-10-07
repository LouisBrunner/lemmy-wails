package logging

import (
	"io"
	"os"
	"path/filepath"

	"github.com/LouisBrunner/lemmy-wails/backend/internal/common"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

func NewLogger(dir string) (*logrus.Logger, func()) {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		CallerPrettyfier: prettyCaller,
	})

	log.SetLevel(logrus.DebugLevel)

	closer := func() {}
	err := common.MkdirIfDoNotExist(dir)
	if err != nil {
		log.WithError(err).Errorf("failed to create logging directory: %s", dir)
	} else {
		path := filepath.Join(dir, "app.log")
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.WithError(err).Errorf("failed to open log for writing: %s", path)
		} else {
			closer = func() {
				err = file.Close()
				if err != nil {
					log.WithError(err).Error("failed to close log file")
				}
			}
			log.SetOutput(&writerMux{
				writers: []io.Writer{
					os.Stderr,
					file,
				},
			})
		}
	}
	return log, closer
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

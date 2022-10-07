package internal

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func Init(logger *logrus.Logger, folder string, force bool) error {
	if !force {
		_, err := os.Stat(filepath.Join(folder, "wails.json"))
		if err == nil {
			logger.Debug("already init'd")
			return nil
		}
	}

	panic("nope")
}

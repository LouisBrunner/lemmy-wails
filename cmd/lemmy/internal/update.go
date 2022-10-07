package internal

import "github.com/sirupsen/logrus"

func Update(logger *logrus.Logger, version, folder string) error {
	err := updateGo(logger, folder, version)
	if err != nil {
		return err
	}
	err = updateNPM(logger, folder, version)
	if err != nil {
		return err
	}
	err = updateCI(logger, folder, version)
	if err != nil {
		return err
	}
	return makeInstall(logger, folder)
}

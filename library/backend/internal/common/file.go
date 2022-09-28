package common

import "os"

func MkdirIfDoNotExist(dir string) error {
	err := os.Mkdir(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

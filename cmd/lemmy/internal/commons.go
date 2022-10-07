package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func mkdirAll(folder string) error {
	return os.MkdirAll(folder, 0755)
}

func execIn(logger *logrus.Logger, folder string, exe string, args ...string) error {
	logger.Debugf("running %s %s in %q", exe, strings.Join(args, " "), folder)
	cmd := exec.Command(exe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = folder
	return cmd.Run()
}

func makeInstall(logger *logrus.Logger, folder string) error {
	return execIn(logger, folder, "make", "install")
}

func updateGo(logger *logrus.Logger, folder, version string) error {
	return execIn(logger, folder, "go", "get", fmt.Sprintf("%s@%s", packageURL, version))
}

func updateNPM(logger *logrus.Logger, folder, version string) error {
	if version == "latest" {
		version = "main"
	}
	return execIn(logger, folder, "npm", "up", fmt.Sprintf("%s?%s", npmPackage, version))
}

func updateCI(logger *logrus.Logger, folder, version string) error {
	if version == "latest" {
		version = "main"
	}

	ciFile := filepath.Join(folder, pathCI)

	input, err := os.ReadFile(ciFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, ciLine) {
			newCILine := fmt.Sprintf("%s@%s", ciLine, version)
			lines[i] = findCILine.ReplaceAllString(line, newCILine)
			break
		}
	}
	output := strings.Join(lines, "\n")

	return os.WriteFile(ciFile, []byte(output), 0644)
}

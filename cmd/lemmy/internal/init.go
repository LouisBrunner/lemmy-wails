package internal

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

//go:embed all:boilerplate
var boilerplate embed.FS

type InitData struct {
	Name               string // e.g. My Super App
	Repo               string // e.g. github.com/user/repo
	BoilerplateVersion string
	Author             InitDataAuthor
	Boilerplate        initDataBoilerplate
}

type InitDataAuthor struct {
	Name  string
	Email string
}

type initDataBoilerplate struct {
	Title      string
	Name       string
	URL        string
	NPMPackage string
	CIURL      string
}

func Init(logger *logrus.Logger, data InitData, folder string, force bool) error {
	if !force && exists(filepath.Join(folder, "wails.json")) {
		logger.Debug("already init'd")
		return nil
	}

	err := mkdirAll(folder)
	if err != nil {
		return err
	}

	data.Boilerplate = initDataBoilerplate{
		Title:      packageTitle,
		Name:       packageName,
		URL:        packageURL,
		NPMPackage: npmPackage,
		CIURL:      ciURL,
	}

	templateMe := func(content string) (string, error) {
		tmpl, err := template.New("file").Parse(content)
		if err != nil {
			return "", err
		}

		w := &strings.Builder{}
		err = tmpl.Execute(w, data)
		if err != nil {
			return "", err
		}

		return w.String(), nil
	}

	err = copyFrom(logger, templateMe, "boilerplate", folder)
	if err != nil {
		return err
	}

	err = updateGo(logger, folder, data.BoilerplateVersion)
	if err != nil {
		return err
	}

	return makeInstall(logger, folder)
}

type templator func(content string) (string, error)

const templateSuffix = ".tmpl"

func copyFrom(logger *logrus.Logger, templateMe templator, from, to string) error {
	logger.Debugf("copying: %q to %q", from, to)

	files, err := boilerplate.ReadDir(from)
	if err != nil {
		return fmt.Errorf("failed to open %q: %w", from, err)
	}

	for _, file := range files {
		src := filepath.Join(from, file.Name())
		dest := filepath.Join(to, file.Name())
		if file.IsDir() {
			logger.Debugf("copying directory: %q to %q", src, dest)

			err = mkdirAll(dest)
			if err != nil {
				return err
			}
			err = copyFrom(logger, templateMe, src, dest)
			if err != nil {
				return err
			}
		} else {
			logger.Debugf("copying file: %q to %q", src, dest)

			contentRaw, err := boilerplate.ReadFile(src)
			if err != nil {
				return fmt.Errorf("could not read %q: %w", src, err)
			}
			content := string(contentRaw)

			if strings.HasSuffix(src, templateSuffix) {
				dest = strings.TrimSuffix(dest, templateSuffix)
				logger.Debugf("templating file: %q to %q", src, dest)
				content, err = templateMe(content)
				if err != nil {
					return fmt.Errorf("could not template %q: %w", src, err)
				}
			}

			err = os.WriteFile(dest, []byte(content), 0644)
			if err != nil {
				return fmt.Errorf("could not write %q: %w", dest, err)
			}
		}
	}
	return nil
}

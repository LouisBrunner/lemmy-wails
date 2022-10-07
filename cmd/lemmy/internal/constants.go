package internal

import (
	"fmt"
	"path/filepath"
	"regexp"
)

var (
	packageTitle     = "Lemmy Wails"
	packageName      = "lemmy-wails"
	packageNamespace = fmt.Sprintf("LouisBrunner/%s", packageName)
	packageURL       = fmt.Sprintf("github.com/%s", packageNamespace)
	packageCommand   = "lemmy"
	pathFrontend     = "frontend"
	pathCI           = filepath.Join(".github", "workflows", "build.yml")
	ciURL            = fmt.Sprintf("%s/.github/workflows/build-app.yml", packageNamespace)
	ciLine           = fmt.Sprintf("uses: %s", ciURL)
	npmPackage       = fmt.Sprintf("https://gitpkg.now.sh/%s/%s", packageNamespace, pathFrontend)
	findCILine       = regexp.MustCompile(fmt.Sprintf("%s@\\S+", ciLine))
)

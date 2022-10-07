package logging

import (
	"fmt"
	"regexp"
	"runtime"
)

const maxCompLen = 20

func trunkSize(in string, maxLen int) string {
	if len(in) > maxLen {
		return in[len(in)-maxLen:]
	}
	return in
}

// e.g. with /Users/me/go/repo/package/file.go, it will match package/file.go
// e.g. with file.go, it will match file.go
var filePathMatcher = regexp.MustCompile(`(([^/]+/)?[^/]+)$`)

// e.g. with /Users/me/go/repo/package/file.func, it will match file.func
// e.g. with func, it will match func
var funcPackageMatcher = regexp.MustCompile(`([^/]+)$`)

// e.g. with myfunc.func1.2, it will match myfunc
// e.g. with myfunc.func3, it will match myfunc
// e.g. with myfunc, it will match myfunc
var anonFuncMatcher = regexp.MustCompile(`\.func[0-9]+(\.[0-9]+)?$`)

func prettyCaller(f *runtime.Frame) (function string, file string) {
	newFile := f.File
	// remove everything but the filename and parent directory (if available)
	if loc := filePathMatcher.FindStringIndex(newFile); loc != nil {
		newFile = newFile[loc[0]:]
	}
	newFile = trunkSize(newFile, maxCompLen)

	newFunction := f.Function
	// strip anonymous functions like `func1` or `func2`
	if loc := anonFuncMatcher.FindStringIndex(newFunction); loc != nil {
		newFunction = newFunction[:loc[0]]
	}
	// only keep the last part of the package path (e.g. main.main), not the whole path
	if loc := funcPackageMatcher.FindStringIndex(newFunction); loc != nil {
		newFunction = newFunction[loc[0]:]
	}
	newFunction = trunkSize(newFunction, maxCompLen)

	return newFunction, fmt.Sprintf("%s:%d", newFile, f.Line)
}

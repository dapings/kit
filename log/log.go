package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

var (
	hostName        string
	loggerPrefixLen int
)

const (
	rotateFileNameLayout     = "20060102"
	fileNameAndLineNumFormat = "[%s:%d]"
)

func setupPrefixLen() {
	_, filePath, _, _ := runtime.Caller(2)
	loggerPrefixLen = len(path.Dir(filePath) + "/")
}

// DefaultFormat returns a %+v args format string.
func DefaultFormat(args ...any) string {
	return strings.Repeat("%+v", len(args))
}

// CallerFile reports file name and line number info using defined format.
func CallerFile(skip int) string {
	// report file and line number info
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return fmt.Sprintf(fileNameAndLineNumFormat, "???", 0)
	}

	var fileName string
	if loggerPrefixLen > 0 && len(file) > loggerPrefixLen {
		fileName = file[loggerPrefixLen:]
	} else {
		_, fileName = path.Split(file)
	}

	return fmt.Sprintf(fileNameAndLineNumFormat, fileName, line)
}

func Error(args ...interface{}) {}

func ErrorExt(skip int, args ...interface{}) {}

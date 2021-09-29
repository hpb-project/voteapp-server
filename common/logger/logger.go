package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

const (
	// LogLevelInfo info level
	LogLevelInfo = "info"
	// LogLevelDebug debug level
	LogLevelDebug = "debug"
	// LogLevelWarn warn level
	LogLevelWarn = "warn"
	// LogLevelError error level
	LogLevelError = "error"

	LogFilename = "hpbvote.log"
)

// Init logrus configuration
func Init(fileDir string, level string) {
	l := strings.ToLower(level)
	switch l {
	case LogLevelInfo:
		log.SetLevel(log.InfoLevel)
	case LogLevelDebug:
		log.SetLevel(log.DebugLevel)
	case LogLevelWarn:
		log.SetLevel(log.WarnLevel)
	case LogLevelError:
		log.SetLevel(log.ErrorLevel)
	}

	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		TimestampFormat:        "2006-01-02 15:04:05", // the "time" field configuration
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmed everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)

	if fileDir == "" {
		log.Infoln("log file path is empty, only output log to os.Stdout")
		log.SetOutput(os.Stdout)
		return
	}

	fp := filepath.Join(fileDir, LogFilename)

	writer, err := rotatelogs.New(
		fp+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fp),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.WithError(err).Error("fail to new rotate logs")
		log.SetOutput(os.Stdout)
		return
	}

	// output log to the STDOUT and File at the same time
	mw := io.MultiWriter(os.Stdout, writer)
	log.SetOutput(mw)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) > 1 {
		return arr[len(arr)-2] + "/" + arr[len(arr)-1]
	}

	return path
}

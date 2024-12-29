package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

var DEFAULT_LOG_DIR = filepath.Join(os.TempDir(), "tipJar.log")

func InitializeFileLogger(logPath string, level log.Level) {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		stdLog.Fatal("failed to open log file", "error", err)
	}

	logger := log.NewWithOptions(file, log.Options{
		Level:           level,
		ReportCaller:    true,
		ReportTimestamp: false,
		Formatter:       log.JSONFormatter,
	})
	// set the charm global default logger
	log.SetDefault(logger)
}

func InitializeNullLogger() {
	log.SetDefault(log.New(io.Discard))
}

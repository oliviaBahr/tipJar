package logger

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func InitializeFileLogger(logDir string, level log.Level) (err error) {
	f, err := os.OpenFile(logDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("failed to open log file: %v\n", err)
		return err
	}
	logger := log.NewWithOptions(f, log.Options{
		Level:           level,
		ReportCaller:    true,
		ReportTimestamp: false,
		Formatter:       log.JSONFormatter,
	})
	// set the charm global default logger
	log.SetDefault(logger)
	return nil
}

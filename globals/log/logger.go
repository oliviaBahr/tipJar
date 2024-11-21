package log

import (
	"fmt"
	"os"
	"path/filepath"
	"tipJar/globals/utils"

	charmLog "github.com/charmbracelet/log"
)

var logger = newLogger()

func newLogger() *charmLog.Logger {
	var writer *os.File
	dir, err := utils.GetRepoDir()
	if err == nil {
		writer, err = os.OpenFile(filepath.Join(dir, ".tmp", "log.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			writer = os.NewFile(0, os.DevNull)
		}
	} else {
		writer = os.NewFile(0, os.DevNull)
	}

	lgr := charmLog.NewWithOptions(writer, charmLog.Options{
		Level:           charmLog.DebugLevel,
		ReportCaller:    true,
		ReportTimestamp: false,
	})
	styles := charmLog.DefaultStyles()
	lgr.SetStyles(styles)
	return lgr
}

func InitializeLogger(logDir string, level charmLog.Level) (err error) {
	if logDir == "stderr" {
		logger.SetOutput(os.Stderr)
	} else {
		f, err := os.OpenFile(logDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("failed to open log file: %v\n", err)
			return err
		}
		logger.SetOutput(f)
	}
	logger.SetLevel(level)
	return nil
}

func Debug(msg string, args ...any) {
	logger.Helper() // ignore this caller function to log the actual caller
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Helper()
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Helper()
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Helper()
	logger.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	logger.Helper()
	logger.Fatal(msg, args...)
}

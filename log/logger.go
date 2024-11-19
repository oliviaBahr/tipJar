package log

import (
	"os"
	fp "path/filepath"
	"tipJar/utils"

	"github.com/charmbracelet/lipgloss"
	charmLog "github.com/charmbracelet/log"
	"github.com/muesli/termenv"
)

var Log = newLogger()

func newLogger() *charmLog.Logger {
	var writer *os.File
	logDir, err := utils.GetRepoDir()
	if err != nil { // log to null
		writer, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
	} else {
		writer, err = os.OpenFile(fp.Join(logDir, "log.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	}

	Logger := charmLog.NewWithOptions(writer, charmLog.Options{
		Level:           charmLog.DebugLevel,
		ReportCaller:    true,
		ReportTimestamp: false,
	})
	lipgloss.SetColorProfile(termenv.ANSI256)
	styles := charmLog.DefaultStyles()
	Logger.SetStyles(styles)
	return Logger
}

func Debug(msg string, args ...any) {
	Log.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	Log.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	Log.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	Log.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	Log.Fatal(msg, args...)
}

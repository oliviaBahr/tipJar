package main

import (
	"fmt"
	"os"
	"tipJar/globals/logger"
	"tipJar/ui"

	"github.com/charmbracelet/log"

	"github.com/spf13/cobra"
)

var stdLog = log.New(os.Stdout)
var logFile string
var logLevel string

func validateLogLevel() {
	if _, ok := logger.LOG_LEVELS[logLevel]; !ok {
		stdLog.Fatal("invalid log level", "e", fmt.Errorf("invalid log level: %s", logLevel))
	}
}

func validateLogFile() {
	logPath := ""
	if logTo := logFile; logTo != "" {
		logPath = logTo
	} else if logLevel == "true" {
		logPath = logger.DEFAULT_LOG_DIR
	} else {
		logPath = os.DevNull
	}

	// open log file
	_, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		stdLog.Fatal("failed to open log file", "e", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "tipJar",
	Short: "tipJar is a CLI tool for managing notes",
	Run: func(cmd *cobra.Command, args []string) {
		stdLog.Info("running UI", "logFile", logFile, "logLevel", logLevel)
		err := ui.RunUI()
		if err != nil {
			stdLog.Fatal("Main UI failed", "e", err)
		}
	},
}

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "listen to the log file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateLogFile()
		logger.Listen(logFile)
	},
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Run tipJar with logging enabled",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validateLogLevel()
		validateLogFile()
		logger.InitializeFileLogger(logFile, logger.LOG_LEVELS[logLevel])
		rootCmd.Run(cmd, args)
	},
}

func init() {
	logCmd.Flags().StringVarP(&logLevel, "level", "l", "debug", "set log level (debug, info, warn, error, fatal)\ndefault: debug")
	logCmd.Flags().StringVarP(&logFile, "log-file", "f", logger.DEFAULT_LOG_DIR, "specify the file to log to")
	listenCmd.Flags().StringVarP(&logFile, "log-file", "f", logger.DEFAULT_LOG_DIR, "specify the file to log to")

	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(listenCmd)

	logger.InitializeNullLogger()
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		stdLog.Fatal("fatal error", "e", err)
	}
}

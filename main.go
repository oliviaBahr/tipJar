package main

import (
	"os"
	"tipJar/globals/logger"
	"tipJar/ui"

	"github.com/charmbracelet/log"

	"github.com/spf13/cobra"
)

var stdLog = log.New(os.Stdout)

var rootCmd = &cobra.Command{
	Use:   "tipJar",
	Short: "tipJar is a CLI tool for managing notes",
	Run: func(cmd *cobra.Command, args []string) {
		logLevel := log.DebugLevel // default level
		switch {
		case cmd.Flag("debug").Changed:
			logLevel = log.DebugLevel
		case cmd.Flag("info").Changed:
			logLevel = log.InfoLevel
		case cmd.Flag("warn").Changed:
			logLevel = log.WarnLevel
		case cmd.Flag("error").Changed:
			logLevel = log.ErrorLevel
		}
		// init global logger
		err := logger.InitializeFileLogger(cmd.Flag("log-to").Value.String(), logLevel)
		if err != nil {
			stdLog.Fatal("failed to initialize logger", "e", err)
		}
		// enter ui
		err = ui.RunUI()
		if err != nil {
			stdLog.Fatal("Main UI failed", "e", err)
		}
	},
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		stdLog.Error("cobra Root command failed", "e", err)
		return err
	}
	return nil
}

func init() {
	rootCmd.Flags().StringP("log-to", "l", "tmp/log.log", "stream or file to log to")

	// Add boolean flags for each log level
	rootCmd.Flags().BoolP("debug", "d", false, "set log level to debug")
	rootCmd.Flags().BoolP("info", "i", false, "set log level to info")
	rootCmd.Flags().BoolP("warn", "w", false, "set log level to warn")
	rootCmd.Flags().BoolP("error", "e", false, "set log level to error")

	// Make the log level flags mutually exclusive
	rootCmd.MarkFlagsMutuallyExclusive("debug", "info", "warn", "error")
}

func main() {
	err := Execute()
	if err != nil {
		stdLog.Fatal("fatal error", "e", err)
	}
}

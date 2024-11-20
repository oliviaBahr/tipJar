package main

import (
	"tipJar/globals/log"
	"tipJar/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tipJar",
	Short: "tipJar is a CLI tool for managing notes",
	Run: func(cmd *cobra.Command, args []string) {
		err := ui.RunUI()
		if err != nil {
			log.Fatal("Main UI failed", "e", err)
		}
	},
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		log.Error("cobra Root command failed", "e", err)
		return err
	}
	return nil
}

func init() {}

func main() {
	log.Info("===== starting tipJar =====")
	err := Execute()
	if err != nil {
		log.Fatal("fatal error", "e", err)
	}
}

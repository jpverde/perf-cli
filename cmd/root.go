package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "perf-cli",
	Short: "perf-cli is a personal, perfect CLI for my day to day work",
	Long:  "perf-cli is a personal, perfect CLI for my day to day work which is full of multiple utilities used to debug, monitor and operate inside of servers",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. an error occured while executing perf-cli: '%s'\n", err)
		os.Exit(1)
	}
}

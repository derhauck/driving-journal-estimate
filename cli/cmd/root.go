package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Driving Journal Estimate",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	_:
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(CalendarCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Driving Journal Estimate ${VERSION}",
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

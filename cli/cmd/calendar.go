package cmd

import (
	"derhauck/driving-journal-estimate/cmd/calendar"

	"github.com/spf13/cobra"
)

var CalendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Estimate journal based on days.",
	Long:  `Estimate journal based on days. Either randomly or with detailed config`,
}

func init() {
	CalendarCmd.AddCommand(calendar.RandomCmd)
	CalendarCmd.AddCommand(calendar.ConfigCmd)
	CalendarCmd.AddCommand(calendar.ServerCmd)
	CalendarCmd.PersistentFlags().Bool("out", false, "Output to file")
}

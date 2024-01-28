package cmd

import (
	"driving-journal-estimate/cmd/calendar"
	"github.com/spf13/cobra"
)

var CalendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
}

func init() {
	CalendarCmd.AddCommand(calendar.RandomCmd)
}

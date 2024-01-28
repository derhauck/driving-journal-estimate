package calendar

import (
	"driving-journal-estimate/public/calendar"
	"github.com/spf13/cobra"
)

var totalFlag = "total"
var daysFlag = "days"

var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		total, err := cmd.Flags().GetFloat32(totalFlag)
		if err != nil {
			return err
		}
		days, err := cmd.Flags().GetInt(daysFlag)
		if err != nil {
			return err
		}
		month := calendar.NewRandomMonth(days)
		err = month.Calculate(total)
		month.Print()
		return err
	},
}

func init() {
	RandomCmd.Flags().Float32(totalFlag, 10000, "Total amount of kilometers driven")
	RandomCmd.Flags().Int(daysFlag, 30, "Number of days driven")
}

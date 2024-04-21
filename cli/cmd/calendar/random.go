package calendar

import (
	"derhauck/driving-journal-estimate/factory"
	"derhauck/driving-journal-estimate/public/logger"

	"github.com/spf13/cobra"
)

var totalFlag = "total"
var daysFlag = "days"

var logLevelFlag = "log"

var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generates estimates randomly.",
	Long:  `Generates estimates randomly. Useful for a quick approximation without specific requirements.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		total, err := cmd.Flags().GetFloat64(totalFlag)
		if err != nil {
			return err
		}
		days, err := cmd.Flags().GetUint(daysFlag)
		if err != nil {
			return err
		}
		level, err := cmd.Flags().GetString(logLevelFlag)
		if err != nil {
			return err
		}
		out, err := cmd.Flags().GetBool("out")
		if err != nil {
			return err
		}
		if level != logger.DEFAULT.String() {
			factory.SetLogLevel(level)
		}

		month := factory.NewRandomMonth(days)
		month.Calculate(total)
		month.Print()

		if out {
			month.WriteOut("output.txt")
		}
		return err
	},
}

func init() {
	RandomCmd.Flags().Float64(totalFlag, 10000, "Total amount of kilometers driven")
	RandomCmd.Flags().Int(daysFlag, 30, "Number of days driven")
	RandomCmd.Flags().String(logLevelFlag, logger.DEFAULT.String(), "Log level => DEBUG,INFO,WARNING,ERROR")
}

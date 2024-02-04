package calendar

import (
	"driving-journal-estimate/factory"
	"driving-journal-estimate/public/logger"
	"github.com/spf13/cobra"
)

var fileFlag = "file"
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		total, err := cmd.Flags().GetFloat32(totalFlag)
		if err != nil {
			return err
		}
		level, err := cmd.Flags().GetString(logLevelFlag)
		if err != nil {
			return err
		}
		path, err := cmd.Flags().GetString(fileFlag)
		if err != nil {
			return err
		}
		if level != logger.DEFAULT.String() {
			factory.SetLogLevel(level)
		}

		config := factory.LoadConfigFromFile(path)

		month := factory.NewMonth()
		month.Days = config.DayConfig()
		err = month.Calculate(total)
		month.Print()
		return err
	},
}

func init() {
	ConfigCmd.Flags().String(logLevelFlag, logger.DEFAULT.String(), "Log level => DEBUG,INFO,WARNING,ERROR")
	ConfigCmd.Flags().String("file", "config.yaml", "config yaml with days")
	ConfigCmd.Flags().Float32(totalFlag, 10000, "Total amount of kilometers driven")
}

package calendar

import (
	"derhauck/driving-journal-estimate/factory"
	"derhauck/driving-journal-estimate/public/logger"

	"github.com/spf13/cobra"
)

var fileFlag = "file"
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Generates estimates based on config file.",
	Long:  `Generates estimates based on a detailed YAML configuration for more accurate distribution reflecting actual driving patterns.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		level, err := cmd.Flags().GetString(logLevelFlag)
		if err != nil {
			return err
		}
		path, err := cmd.Flags().GetString(fileFlag)
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

		config := factory.LoadConfigFromFile(path)

		month := factory.NewMonth()
		month.Days = config.DayConfig()
		month.Calculate(config.Total)
		month.Print()
		if out {
			month.WriteOut("output.txt")
		}
		return err
	},
}

func init() {
	ConfigCmd.Flags().String(logLevelFlag, logger.DEFAULT.String(), "Log level => DEBUG,INFO,WARNING,ERROR")
	ConfigCmd.Flags().String("file", "config.yaml", "config yaml with days")
}

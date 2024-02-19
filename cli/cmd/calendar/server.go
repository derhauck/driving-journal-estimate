package calendar

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use: "server", Short: "Starts the calendar server",
	Long: `Starts the calendar server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		router := gin.Default()
		initRoutes(router)
		err := router.Run(":8080")
		if err != nil {
			return err
		}
		return nil
	},
}

package calendar

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the calendar server",
	Long:  `Starts the calendar server. Will offer the same as the other commands but available via REST API.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		router := gin.New()

		router.Use(
			gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
			gin.Recovery(),
		)
		initRoutes(router)
		err := router.Run(":8080")
		if err != nil {
			return err
		}
		return nil
	},
}

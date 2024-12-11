package calendar

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/Cyprinus12138/otelgin"
	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the calendar server",
	Long:  `Starts the calendar server. Will offer the same as the other commands but available via REST API.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()
		router, err := graceful.New(gin.New())
		router.Use(otelgin.Middleware("driving-journal-estimate-server"))
		if err != nil {
			return err
		}
		router.Use(
			gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
			gin.Recovery(),
		)
		initRoutes(router)
		defer router.Close()
		router.StaticFS("/assets", http.Dir("./schema"))
		if err := router.RunWithContext(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Printf("Error: %s", err)
		}
		return nil
	},
}

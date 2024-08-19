package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/nikhildev/gofy/internal/routes"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "api",
	Short: "API",
	Long:  `This command starts the API server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := StartApiServer()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		quit := make(chan os.Signal, 1)

		// Handle interrupt signal from the terminal
		signal.Notify(quit, os.Interrupt)
		// Handle interrupt signal from the Kubernetes
		signal.Notify(quit, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10000)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not gracefully shutdown the server: %v", err)
		}

		fmt.Println("shutting down the server gracefully")
		return nil
	},
}

func StartApiServer() *echo.Echo {
	fmt.Println("Starting API server")

	//This creates a simple http echo server and starts it
	server := echo.New()

	go func() {
		routes.RegisterRoutes(server)

		s := &http.Server{
			Addr:    ":3000",
			Handler: server,
		}
		// For simplicity, we are registering all the routes in the main function

		// Start the server
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("failed to start http server: %s", err)
		}
	}()

	return server

}

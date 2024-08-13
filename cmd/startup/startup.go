package startup

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nikhildev/gofy/cmd/api"
	"github.com/nikhildev/gofy/internal/db"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
 * This is the startup command. This command is used to start the application.
 * This command is a subcommand of the root command.
 * The root command is the parent command and this command is the child command.
 */
func NewStartupCommand(store *db.Store) *cobra.Command {
	startupCommand := &cobra.Command{
		Use:   "startup",
		Short: "Startup",
		Long:  `This command starts the application and introduces dependencies`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting the application ...")

			// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
			// Use a buffered channel to avoid missing signals as recommended for signal.Notify
			quit := make(chan os.Signal, 1)

			// Handle interrupt signal from the terminal
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

			// This defer would catch the interrupts as well and disconnect the mongo client.
			// Because defer functions are executed outside the context of the main function,
			// they need to have the client and context passed to them as dependencies.
			defer func(client *mongo.Client, ctx context.Context) {
				<-quit
				cleanup(client, ctx)
				os.Exit(1)
			}(store.Client, nil)

			api.StartApiServer()
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Preparing to start the application ...")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Application startup completed!")
		},
	}

	return startupCommand
}

func cleanup(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from MongoDB")
}

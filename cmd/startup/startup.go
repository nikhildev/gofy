package startup

import (
	"context"
	"fmt"
	"github.com/nikhildev/gofy/cmd/api"
	"github.com/nikhildev/gofy/internal/db"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
)

func NewStartupCommand(store *db.Store) *cobra.Command {
	// This is where we call the mongodb.NewStore() function to create a new instance of the Store struct. This function returns a pointer to the Store struct and an error. We are ignoring the error for now, but in a real-world application, you should handle it properly.

	startupCommand := &cobra.Command{
		Use:   "startup",
		Short: "Startup",
		Long:  `This command starts the application and introduces dependencies`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting the application")

			// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
			// Use a buffered channel to avoid missing signals as recommended for signal.Notify
			quit := make(chan os.Signal)

			// Handle interrupt signal from the terminal
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

			//This defer would catch the interrupts as well and disconnect the mongo client
			defer func(client *mongo.Client, ctx context.Context) {
				<-quit
				cleanup(client, ctx)
				os.Exit(1)
			}(store.Client, nil)

			api.StartApiServer()
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Preparing to start the application")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Application startup completed")
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

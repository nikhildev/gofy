package greet

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

/*
 * This is the greet command. This command is used to start the application.
 * This command is a subcommand of the root command.
 * The root command is the parent command and this command is the child command.
 * In this context, we can use the startup command for the following
 * - Gracefully shutting down the application when interrupted
 */

var Command = &cobra.Command{
	Use:   "greet",
	Short: "Greet",
	Long:  `This command greets the user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Greeting the user ...")

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		quit := make(chan os.Signal, 1)

		// Handle interrupt signal from the terminal
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		// This defer would catch the interrupts as well and disconnect the mongo client.
		// Because defer functions are executed outside the context of the main function,
		// they need to have the client and context passed to them as dependencies.
		defer func() {
			<-quit
			fmt.Println("Shutting down ...")
			os.Exit(0)
		}()

		greet()

		return nil
	},
}

func greet() {
	fmt.Println("Hello, World!")
}

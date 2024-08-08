package main

import (
	"fmt"

	"github.com/nikhildev/gofy/cmd/api"
	"github.com/nikhildev/gofy/cmd/startup"
	"github.com/nikhildev/gofy/internal/db"
	"github.com/spf13/cobra"
)

func newRootCommand(store *db.Store) *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:   "gofy",
		Short: "Go Funk Yourself",
		Long:  `Just my playground to try go`,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Root command")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Root command started")

		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Root command completed")
		},
	}
	rootCommand.AddCommand(startup.NewStartupCommand(store))
	rootCommand.AddCommand(api.NewApiServerCommand())
	return rootCommand
}

func main() {
	// We get a new store here so as to pass it down to all commands that will need it.
	// This also ensures that all downstream commands and subprocesses will not have to recreate new clients
	store, err := db.NewStore(nil)
	if err != nil {
		panic(err)
	}

	rootCommand := newRootCommand(store)

	if err := rootCommand.Execute(); err != nil {
		fmt.Println("rootCommand encountered fatal exception %s", err)
	}

}

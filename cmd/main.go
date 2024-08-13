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
		Use:   "gofy", // This is the name of the command. This is what will be used to run the command. ex: 'gofy startup' or 'gofy api'
		Short: "Go Funk Yourself",
		Long:  `Just my playground to try go`,
		PreRun: func(cmd *cobra.Command, args []string) { // This is a pre-run hook. This will be executed before the command is run
			fmt.Println("Starting Root command ...")
		},
		Run: func(cmd *cobra.Command, args []string) { // This is the main function that will be executed when the command is run
			fmt.Println("Root command started")
		},
		PostRun: func(cmd *cobra.Command, args []string) { // This is a post-run hook. This will be executed after the command is run
			fmt.Println("Root command completed")
		},
	}

	// We add all the subcommands here. Consider these are nested commands. These are the commands that will be registered under the root command.
	// The root command is the parent command and the subcommands are the child commands.
	// The root command can have multiple subcommands. Each subcommand can have its own subcommands as well.
	// This is how we can create a tree of commands.
	// The root command is the root of the tree and all other commands are the nodes of the tree.
	rootCommand.AddCommand(startup.NewStartupCommand(store))

	// ex: 'gofy startup' is a subcommand of 'gofy' and 'gofy api' is another subcommand of 'gofy'
	rootCommand.AddCommand(api.NewApiServerCommand())
	return rootCommand
}

func main() {
	/*
	 * The main function when using cobra is very simple. We just create a new root command and execute it.
	 * The root command will take care of executing all the subcommands.
	 * Its probably a good idea to initialize the store here and pass it down to all the subcommands.
	 */
	store, err := db.NewStore(nil)
	if err != nil {
		panic(err)
	}

	rootCommand := newRootCommand(store)

	if err := rootCommand.Execute(); err != nil {
		fmt.Printf("rootCommand encountered fatal exception %s", err)
	}

}

package cmd

import (
	"fmt"
	"os"

	"github.com/gambit-interactive/gambit-cli/cmd/create"
	"github.com/gambit-interactive/gambit-cli/cmd/docker"
	"github.com/gambit-interactive/gambit-cli/internal/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gambit",
	Short: "Gambit CLI is a tool for managing your Gambit Framework projects.",
	Long: `A complete and efficient CLI for creating, managing, and interacting with
applications built on the Gambit Framework.`,
	// This function runs before any subcommand, except for calls like 'gambit --help'
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Only print the logo for action commands, not for help or version,
		// to keep the output clean.
		if cmd.Name() != "help" && cmd.Name() != "version" && cmd.Name() != "gambit" {
			utils.PrintLogo()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI: '%s'", err)
		os.Exit(1)
	}
}

// init is called by Go when the package is initialized.
func init() {
	// Add child commands to the root command here.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(create.NewCreateCmd())
	rootCmd.AddCommand(docker.NewDockerCmd())
}

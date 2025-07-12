package create

import (
	"github.com/spf13/cobra"
)

// NewCreateCmd creates and returns the 'create' command.
// This command serves as an entry point for other creation-related subcommands.
func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates new Gambit projects or files.",
		Long:  `The 'create' command allows you to quickly scaffold a new Gambit project or add specific files to an existing project.`,
	}

	// Add subcommands to the 'create' command
	cmd.AddCommand(NewCreateProjectCmd())
	// In the future, you could add: cmd.AddCommand(NewCreateControllerCmd())

	return cmd
}

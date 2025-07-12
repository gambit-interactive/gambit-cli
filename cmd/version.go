package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// We define the version here. In a real project, this would be injected at build time.
const cliVersion = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of the Gambit CLI.",
	Long:  `Every respectable CLI needs a version command. This is ours.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Gambit Framework CLI Version %s\n", cliVersion)
	},
}

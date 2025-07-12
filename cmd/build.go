package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the Gambit application binary.",
	Long:  `Compiles the application and creates an executable binary in the project root.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building application binary...")
		buildApp("go", "build", "-o", "gambit-app", "./cmd/api")
		fmt.Println("✅ Build successful! Executable 'gambit-app' created.")
	},
}

func buildApp(command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Build failed: %v\n%s", err, string(output))
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

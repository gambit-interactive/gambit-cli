package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the Gambit application.",
	Long:  `Compiles and runs the main application file. This is suitable for quick tests but does not provide hot-reloading.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Gambit application...")
		runApp("go", "run", "./cmd/api")
	},
}

// devCmd represents the dev command with hot-reloading
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Runs the application in development mode with hot-reloading.",
	Long:  `Starts the application using Air for live-reloading upon file changes. Requires Air to be installed (https://github.com/cosmtrek/air).`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := exec.LookPath("air"); err != nil {
			log.Fatal("Error: 'air' is not installed or not in your PATH. Please install it to use 'gambit dev': go install github.com/cosmtrek/air@latest")
		}
		fmt.Println("Starting development server with hot-reloading (Air)...")
		runApp("air")
	},
}

// Helper function to run external commands
func runApp(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("Error running the application: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(devCmd)
}

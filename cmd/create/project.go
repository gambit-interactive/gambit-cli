package create

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

const frameworkRepoURL = "https://github.com/gambit-interactive/gambit-framework.git"

// NewCreateProjectCmd creates the 'gambit create project <name>' command.
func NewCreateProjectCmd() *cobra.Command {
	var outputDir string
	var branch string

	cmd := &cobra.Command{
		Use:   "project [name]",
		Short: "Creates a new Gambit project by cloning the official framework.",
		Long: `Creates a new project by cloning the Gambit Framework repository.
By default, it uses the 'stable' branch for production-ready projects.

Example:
  gambit create project my-api
  gambit create project my-api -o ./services/api
  gambit create project my-dev-api --branch master`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			targetDir := projectName

			if outputDir != "" {
				targetDir = outputDir
			}

			fmt.Printf("Cloning Gambit Framework from branch '%s' into '%s'...\n", branch, targetDir)

			// Execute the 'git clone' command
			gitCmd := exec.Command("git", "clone", "--branch", branch, frameworkRepoURL, targetDir)
			gitCmd.Stdout = os.Stdout
			gitCmd.Stderr = os.Stderr

			if err := gitCmd.Run(); err != nil {
				log.Fatalf("Error cloning repository: %v. Make sure Git is installed and you have network access.", err)
			}

			// Remove the .git directory to start a fresh repository for the user
			if err := os.RemoveAll(filepath.Join(targetDir, ".git")); err != nil {
				log.Printf("Warning: could not remove .git directory: %v", err)
			}

			// Initialize a new git repository for the user's project
			fmt.Println("Initializing a new git repository...")
			initCmd := exec.Command("git", "init")
			initCmd.Dir = targetDir // Run the command in the new project's directory
			if err := initCmd.Run(); err != nil {
				log.Printf("Warning: could not initialize a new git repository: %v", err)
			}

			fmt.Printf("\n✅ Project '%s' created successfully!\n", projectName)
			fmt.Printf("\nNext steps:\n")
			fmt.Printf("  1. cd %s\n", targetDir)
			fmt.Printf("  2. Run 'go mod tidy' to sync dependencies.\n")
			fmt.Printf("  3. Run 'gambit dev' to start the development server.\n")
		},
	}

	cmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for the new project")
	cmd.Flags().StringVarP(&branch, "branch", "b", "stable", "The repository branch to clone (e.g., stable, master)")

	return cmd
}

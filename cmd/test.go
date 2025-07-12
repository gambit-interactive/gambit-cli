package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs tests for the project.",
	Long:  `Executes 'go test' for all packages within the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running tests...")
		testCmd := exec.Command("go", "test", "./...")
		output, err := testCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Tests failed: %v\n%s", err, string(output))
		}
		fmt.Println(string(output))
		fmt.Println("✅ All tests passed!")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

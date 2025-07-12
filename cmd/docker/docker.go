package docker

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// NewDockerCmd creates the root 'docker' command
func NewDockerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker",
		Short: "Manage Docker resources for your project.",
		Long:  `A collection of commands to build Docker images and run services using Docker Compose. Assumes docker-compose.yml exists.`,
	}

	cmd.AddCommand(newDockerBuildCmd())
	cmd.AddCommand(newDockerUpCmd())

	return cmd
}

func newDockerBuildCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Builds the Docker image for the application.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Building Docker image...")
			runDocker("docker", "build", "-t", "gambit-app", ".")
			fmt.Println("✅ Docker image built successfully!")
		},
	}
}

func newDockerUpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Starts the application services using Docker Compose.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting services with Docker Compose...")
			runDocker("docker-compose", "up", "-d")
			fmt.Println("✅ Services started! Run 'docker-compose logs -f' to see the logs.")
		},
	}
}

func runDocker(command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Docker command failed: %v\n%s", err, string(output))
	}
	fmt.Println(string(output))
}

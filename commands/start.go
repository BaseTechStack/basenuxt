package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Start the development server",
	Long:    `Start the development server using Bun`,
	Example: "basenuxt start     # Run development server\nbasenuxt s        # Run development server (shorthand)",
	Run:     startApplication,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startApplication(cmd *cobra.Command, args []string) {
	// If arguments are provided, show usage info
	if len(args) > 0 {
		fmt.Println("The start command doesn't accept arguments. Use it without arguments to start the development server.")
		fmt.Println("For other commands, use 'basenuxt run <command>' or the shortcuts 'b' for build and 'g' for generate.")
		return
	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Check if bun is installed
	if _, err := exec.LookPath("bun"); err != nil {
		fmt.Println("Bun is not installed. Please install Bun to run the application.")
		fmt.Println("Visit https://bun.sh/ for installation instructions.")
		return
	}

	// Always run dev command
	bunScript := "dev"

	fmt.Printf("Running: bun run %s\n", bunScript)
	bunCmd := exec.Command("bun", "run", bunScript)
	bunCmd.Stdout = os.Stdout
	bunCmd.Stderr = os.Stderr
	bunCmd.Dir = cwd

	if err := bunCmd.Run(); err != nil {
		fmt.Printf("Error running 'bun run %s': %v\n", bunScript, err)
		return
	}
}

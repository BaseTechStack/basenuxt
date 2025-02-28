package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start [command]",
	Short:   "Start the application",
	Long:    `Start the application with various commands using Bun`,
	Example: "basenuxt start dev     # Run development server\nbasenuxt start generate # Generate static site\nbasenuxt start build    # Build the application",
	Run:     startApplication,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startApplication(cmd *cobra.Command, args []string) {
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

	// Default command is "dev" if no subcommand is provided
	bunScript := "dev"

	// If args are provided, use the first arg as the bun script to run
	if len(args) > 0 {
		bunScript = args[0]
	}

	// Map of supported script commands
	supportedScripts := map[string]bool{
		"dev":         true,
		"build":       true,
		"generate":    true,
		"preview":     true,
		"postinstall": true,
		"lint":        true,
		"lint:fix":    true,
	}

	// Check if the script is supported
	if _, exists := supportedScripts[bunScript]; !exists {
		fmt.Printf("Unsupported command: %s\n", bunScript)
		fmt.Println("Supported commands: dev, build, generate, preview, postinstall, lint, lint:fix")
		return
	}

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

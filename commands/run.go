package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [command]",
	Short:   "Run application commands",
	Long:    `Run application commands with Bun`,
	Example: "basenuxt run build    # Build the application\nbasenuxt run generate # Generate static site",
	Run:     runApplication,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runApplication(cmd *cobra.Command, args []string) {
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

	// Require a subcommand
	if len(args) == 0 {
		fmt.Println("Please specify a command to run. Examples:")
		fmt.Println("  basenuxt run build    # Build the application")
		fmt.Println("  basenuxt run generate # Generate static site")
		fmt.Println("Or use the shortcuts:")
		fmt.Println("  basenuxt b            # Build the application")
		fmt.Println("  basenuxt g            # Generate static site")
		return
	}

	bunScript := args[0]

	// We allow any subcommand to be passed directly to bun run

	// Pass any remaining args to the bun command
	bunArgs := append([]string{"run", bunScript}, args[1:]...)
	
	fmt.Printf("Running: bun %s\n", formatBunArgs(bunArgs))
	bunCmd := exec.Command("bun", bunArgs...)
	bunCmd.Stdout = os.Stdout
	bunCmd.Stderr = os.Stderr
	bunCmd.Dir = cwd

	if err := bunCmd.Run(); err != nil {
		fmt.Printf("Error running 'bun %s': %v\n", formatBunArgs(bunArgs), err)
		return
	}
}

// Helper function to format bun args for display
func formatBunArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += " "
		}
		result += arg
	}
	return result
}

package cmd

import (
	"fmt"
	"os"

	"github.com/BaseTechStack/basenuxt/commands/generators"
	"github.com/spf13/cobra"
)

var (
	// Debug flag for verbose output
	debug bool
)

func init() {
	rootCmd.AddCommand(generateCmd)

	// Add debug flag to help with debugging
	generateCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate [name] [field:type...]",
	Aliases: []string{"g"},
	Short:   "Generate a new crud module",
	Long:    `Generate a new crud module with the specified name and fields.`,
	Args:    cobra.MinimumNArgs(1),
	Run:     run,
}

// run executes the generate command with step-by-step output
func run(cmd *cobra.Command, args []string) {
	// Get current working directory
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Parse the entity name and fields
	fmt.Println("🔍 Parsing entity structure...")
	
	entityName := args[0]
	var fieldStrs []string
	
	if len(args) > 1 {
		fieldStrs = args[1:]
	}
	
	// Initialize the generator
	generator := generators.NewGenerator(baseDir, entityName, fieldStrs)
	
	// Generate the entity
	if err := generator.Generate(); err != nil {
		fmt.Printf("❌ Error generating entity: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Generated CRUD module for %s successfully!\n", entityName)
}

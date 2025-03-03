package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BaseTechStack/basenuxt/commands/generators"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "d [name]",
	Short: "Destroy an existing structure",
	Long:  `Destroy an existing structure with the specified name.`,
	Args:  cobra.ExactArgs(1),
	Run:   destroyModule,
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}

func destroyModule(cmd *cobra.Command, args []string) {
	singularName := args[0]
	pluralName := generators.ToKebabCase(generators.ToPlural(singularName))

	// Check if the module exists
	moduleDir := filepath.Join("./structures/", pluralName)
	if _, err := os.Stat(moduleDir); os.IsNotExist(err) {
		fmt.Printf("Structure '%s' does not exist.\n", singularName)
		return
	}

	// Prompt for confirmation with Y preselected
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure you want to destroy the '%s' structure? [Y/n] ", singularName)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response != "" && response != "y" {
		fmt.Println("Operation cancelled.")
		return
	}

	// Delete module directory
	if err := os.RemoveAll(moduleDir); err != nil {
		fmt.Printf("Error removing directory %s: %v\n", moduleDir, err)
		return
	}

	// Update nuxt.config.ts to unregister the module
	if err := updateInitFileForDestroy(pluralName); err != nil {
		fmt.Printf("Error updating nuxt.config.ts: %v\n", err)
		return
	}

	// Remove the entity from the sidebar
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
	}

	if err := generators.RemoveFromSidebar(baseDir, pluralName); err != nil {
		// Just log the error but don't fail the destroy process
		fmt.Printf("Warning: could not update sidebar: %v\n", err)
	}

	fmt.Printf("Successfully destroyed module '%s'.\n", singularName)
}

// updateInitFileForDestroy removes a module entry from nuxt.config.ts
func updateInitFileForDestroy(moduleKebabName string) error {
	// Path to nuxt.config.ts
	configPath := "nuxt.config.ts"

	// Check if the file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("nuxt.config.ts not found: %v", err)
	}

	// Read the config file
	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading nuxt.config.ts: %v", err)
	}

	contentStr := string(content)

	// Create a regex to match the entry with surrounding whitespace and optional comma
	re := regexp.MustCompile(fmt.Sprintf(`(\s*)'\./structures/%s',(\s*)`, moduleKebabName))
	// Replace with just a newline to maintain formatting
	contentStr = re.ReplaceAllString(contentStr, "$1$2")

	// Clean up any double commas that might have been created
	contentStr = strings.Replace(contentStr, ",,", ",", -1)

	// Write the updated content back to the file
	if err := os.WriteFile(configPath, []byte(contentStr), 0644); err != nil {
		return fmt.Errorf("error writing to nuxt.config.ts: %v", err)
	}

	return nil
}

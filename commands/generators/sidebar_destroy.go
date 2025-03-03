package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// RemoveFromSidebar removes a navigation item from TheSidebar.vue for the destroyed entity
func RemoveFromSidebar(baseDir, pluralName string) error {
	// Path to the sidebar component
	sidebarPath := filepath.Join(baseDir, "app", "components", "app", "TheSidebar.vue")

	// Check if the sidebar file exists
	if _, err := os.Stat(sidebarPath); os.IsNotExist(err) {
		return fmt.Errorf("sidebar component not found at %s", sidebarPath)
	}

	// Read the sidebar file
	content, err := os.ReadFile(sidebarPath)
	if err != nil {
		return fmt.Errorf("error reading sidebar component: %v", err)
	}

	// Convert content to string
	sidebarContent := string(content)

	// Create the kebab case plural name for the route path
	kebabCasePluralName := ToKebabCase(pluralName)

	// Check if the entity exists in the sidebar
	if !strings.Contains(sidebarContent, fmt.Sprintf("to: '/%s'", kebabCasePluralName)) {
		fmt.Printf("Navigation item for %s not found in sidebar\n", pluralName)
		return nil
	}

	// Create a regex pattern to match the entire navigation item block
	pattern := fmt.Sprintf(`\s*\{\s*label:\s*['"]%s['"],\s*icon:\s*['"][-\w]+['"],\s*to:\s*['"]/%s['"]\s*\},?`, 
		pluralName, kebabCasePluralName)
	
	re := regexp.MustCompile(pattern)
	
	// Replace the navigation item with an empty string
	updatedContent := re.ReplaceAllString(sidebarContent, "")

	// Clean up any double commas that might have been created
	updatedContent = strings.Replace(updatedContent, ",,", ",", -1)

	// Write the updated content back to the file
	if err := os.WriteFile(sidebarPath, []byte(updatedContent), 0644); err != nil {
		return fmt.Errorf("error writing updated sidebar component: %v", err)
	}

	fmt.Printf("Removed navigation item for %s from sidebar\n", pluralName)
	return nil
}

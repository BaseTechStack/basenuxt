package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// UpdateSidebar adds a new navigation item to TheSidebar.vue for the generated entity
func UpdateSidebar(baseDir, entityName, pluralName string) error {
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

	// Create the new navigation item
	kebabCasePluralName := ToKebabCase(pluralName)

	// Define the new navigation item
	newNavItem := fmt.Sprintf(`    {
        label: '%s',
        icon: 'i-heroicons-document-text',
        to: '/%s'
    }`, pluralName, kebabCasePluralName)

	// Find the position to insert the new item (after the last item in mainNavItems)
	// First check if the entity already exists in the sidebar
	if strings.Contains(sidebarContent, fmt.Sprintf("to: '/%s'", kebabCasePluralName)) {
		fmt.Printf("Navigation item for %s already exists in sidebar\n", pluralName)
		return nil
	}

	// Look for the mainNavItems array
	var updatedContent string

	// Find the position to insert the new item
	mainNavItemsStart := strings.Index(sidebarContent, "const mainNavItems = ref([")
	if mainNavItemsStart == -1 {
		return fmt.Errorf("could not find mainNavItems in sidebar component")
	}

	// Find the end of the mainNavItems array
	mainNavItemsEnd := strings.Index(sidebarContent[mainNavItemsStart:], "])")
	if mainNavItemsEnd == -1 {
		return fmt.Errorf("could not find closing bracket of mainNavItems")
	}

	// Calculate the absolute position of the end of the array
	insertPosition := mainNavItemsStart + mainNavItemsEnd

	// Insert the new item before the closing bracket
	updatedContent = sidebarContent[:insertPosition] + ",\n" + newNavItem + sidebarContent[insertPosition:]

	// Write the updated content back to the file
	if err := os.WriteFile(sidebarPath, []byte(updatedContent), 0644); err != nil {
		return fmt.Errorf("error writing updated sidebar component: %v", err)
	}

	fmt.Printf("Updated sidebar to include navigation for %s\n", pluralName)
	return nil
}

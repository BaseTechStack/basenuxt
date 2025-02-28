package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// UpdateMainNuxtConfig updates the main nuxt.config.ts file to include the new entity
func UpdateMainNuxtConfig(baseDir string, entityDirName string) error {
	// Path to the main nuxt.config.ts file
	nuxtConfigPath := filepath.Join(baseDir, "nuxt.config.ts")

	// Read the current nuxt.config.ts file
	content, err := os.ReadFile(nuxtConfigPath)
	if err != nil {
		return fmt.Errorf("error reading main nuxt.config.ts: %v", err)
	}

	// Convert content to string for easier manipulation
	nuxtConfigContent := string(content)

	// Create the path to add in the extends array
	entityPath := fmt.Sprintf("./structures/%s", entityDirName)

	// Check if the entity is already in the extends array
	if strings.Contains(nuxtConfigContent, entityPath) {
		fmt.Printf("Entity %s is already in the extends array of nuxt.config.ts\n", entityPath)
		return nil
	}

	// Find the extends array
	extendsRegex := regexp.MustCompile(`extends:\s*\[([\s\S]*?)\]`)
	matches := extendsRegex.FindStringSubmatchIndex(nuxtConfigContent)

	if len(matches) < 4 {
		return fmt.Errorf("could not find extends array in nuxt.config.ts")
	}

	// Get the content of the extends array
	extendsContent := nuxtConfigContent[matches[2]:matches[3]]
	
	// Check if the extends array is empty or if the last entry doesn't have a comma
	var newExtendsContent string
	trimmedContent := strings.TrimSpace(extendsContent)
	
	if trimmedContent == "" {
		// If extends is empty
		newExtendsContent = fmt.Sprintf("    '%s'", entityPath)
	} else if !strings.HasSuffix(trimmedContent, ",") {
		// If the last entry doesn't have a comma, add one before adding the new entry
		newExtendsContent = fmt.Sprintf("%s,\n    '%s'", extendsContent, entityPath)
	} else {
		// If there's already a comma after the last entry
		newExtendsContent = fmt.Sprintf("%s\n    '%s'", extendsContent, entityPath)
	}

	// Create the new content for the nuxt.config.ts file
	newNuxtConfigContent := nuxtConfigContent[:matches[2]] + newExtendsContent + nuxtConfigContent[matches[3]:]

	// Write the updated content back to the file
	if err := os.WriteFile(nuxtConfigPath, []byte(newNuxtConfigContent), 0644); err != nil {
		return fmt.Errorf("error writing updated nuxt.config.ts: %v", err)
	}

	fmt.Printf("Updated main nuxt.config.ts to include %s in extends array\n", entityPath)
	return nil
}

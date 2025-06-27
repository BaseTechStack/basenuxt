package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateSinglePage generates the page for an entity with dynamic routing
func GenerateSinglePage(baseDir, pagesDir, entityName, pluralName string, fields []Field) error {
	// Ensure proper casing for paths and names
	kebabPlural := ToKebabCase(ToPascalCase(pluralName))
	
	// Create directory structure for dynamic routing
	// The pages structure should be: structures/[entity]/pages/[entity]/[id]/index.vue
	entityPageDir := filepath.Join(pagesDir, kebabPlural, "[id]")
	outputPath := filepath.Join(entityPageDir, "index.vue")
	
	fmt.Printf("Creating single page at %s\n", outputPath)

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("page_id.vue.tmpl")
	if err != nil {
		return err
	}

	// Create the pages directory if it doesn't exist
	if err := os.MkdirAll(entityPageDir, 0755); err != nil {
		return fmt.Errorf("error creating pages directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("page", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating page file: %v", err)
	}
	defer file.Close()

	// Execute the template with the data
	data := struct {
		StructName string
		PluralName string
		Fields     []Field
		// Add helper method to access field JSON name
		JSONName   func(f Field) string
	}{
		StructName: entityName,
		PluralName: pluralName,
		Fields:     fields,
		// Function to access JSONName from Field
		JSONName: func(f Field) string {
			return f.JSONName
		},
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing page template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

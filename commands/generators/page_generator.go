package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GeneratePage generates the page for an entity
func GeneratePage(baseDir, pagesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the page
	entityPageDir := filepath.Join(pagesDir, ToKebabCase(pluralName))
	outputPath := filepath.Join(entityPageDir, "index.vue")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("page_index.vue.tmpl")
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
	}{
		StructName: entityName,
		PluralName: pluralName,
		Fields:     fields,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing page template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

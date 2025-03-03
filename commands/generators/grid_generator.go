package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateGrid generates the Grid component for an entity
func GenerateGrid(baseDir, componentsDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the component
	outputPath := filepath.Join(componentsDir, "Grid.vue")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("grid.vue.tmpl")
	if err != nil {
		return err
	}

	// Create the component directory if it doesn't exist
	if err := os.MkdirAll(componentsDir, 0755); err != nil {
		return fmt.Errorf("error creating components directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("grid", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating Grid file: %v", err)
	}
	defer file.Close()

	// Execute the template with the data
	data := struct {
		EntityName string
		PluralName string
		Fields     []Field
	}{
		EntityName: entityName,
		PluralName: pluralName,
		Fields:     fields,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing Grid template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

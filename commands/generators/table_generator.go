package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateTable generates the {EntityName}Table component for an entity
func GenerateTable(baseDir, componentsDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the component
	outputPath := filepath.Join(componentsDir, entityName+"Table.vue")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("table.vue.tmpl")
	if err != nil {
		return err
	}

	// Create the component directory if it doesn't exist
	if err := os.MkdirAll(componentsDir, 0755); err != nil {
		return fmt.Errorf("error creating components directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("table", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating %sTable file: %v", entityName, err)
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
		return fmt.Errorf("error executing %sTable template: %v", entityName, err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

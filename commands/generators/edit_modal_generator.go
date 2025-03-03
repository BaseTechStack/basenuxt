package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateEditModal generates the EditModal component for an entity
func GenerateEditModal(baseDir, componentsDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the component
	outputPath := filepath.Join(componentsDir, "EditModal.vue")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("edit_modal.vue.tmpl")
	if err != nil {
		return err
	}

	// Create the component directory if it doesn't exist
	if err := os.MkdirAll(componentsDir, 0755); err != nil {
		return fmt.Errorf("error creating components directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("editModal", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating EditModal file: %v", err)
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
		return fmt.Errorf("error executing EditModal template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

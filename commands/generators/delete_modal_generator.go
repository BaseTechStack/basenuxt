package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateDeleteModal generates the Delete{EntityName}Modal component for an entity
func GenerateDeleteModal(baseDir, componentsDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the component
	outputPath := filepath.Join(componentsDir, "Delete"+entityName+"Modal.vue")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("delete_modal.vue.tmpl")
	if err != nil {
		return err
	}

	// Create the component directory if it doesn't exist
	if err := os.MkdirAll(componentsDir, 0755); err != nil {
		return fmt.Errorf("error creating components directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("deleteModal", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating Delete%sModal file: %v", entityName, err)
	}
	defer file.Close()

	// Execute the template with the data
	data := struct {
		StructName string
		PluralName string
		VarName    string
		Fields     []Field
	}{
		StructName: entityName,
		PluralName: pluralName,
		VarName:    toLowerFirst(entityName),
		Fields:     fields,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing Delete%sModal template: %v", entityName, err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

func toLowerFirst(s string) string {
	return strings.ToLower(strings.TrimPrefix(s, ""))
}

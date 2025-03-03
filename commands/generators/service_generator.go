package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateService generates the service for an entity
func GenerateService(baseDir, servicesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the service
	outputPath := filepath.Join(servicesDir, ToCamelCase(entityName)+"Service.ts")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("entity_service.ts.tmpl")
	if err != nil {
		return err
	}

	// Create the services directory if it doesn't exist
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return fmt.Errorf("error creating services directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("service", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating service file: %v", err)
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
		return fmt.Errorf("error executing service template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

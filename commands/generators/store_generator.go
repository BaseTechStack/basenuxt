package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateEntityStore generates the entityStore.ts file for an entity
func GenerateEntityStore(baseDir, storesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the entity store file
	outputPath := filepath.Join(storesDir, strings.ToLower(pluralName)+"Store.ts")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("entities_store.ts.tmpl")
	if err != nil {
		return err
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(storesDir, 0755); err != nil {
		return fmt.Errorf("error creating stores directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("entityStore", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating entity store file: %v", err)
	}
	defer file.Close()

	// Define the data for the template
	data := struct {
		StructName string
		PluralName string
		Fields     []Field
	}{
		StructName: entityName,
		PluralName: pluralName,
		Fields:     fields,
	}

	// Execute the template with the data
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing entity store template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

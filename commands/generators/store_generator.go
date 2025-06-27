package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateEntityStore generates the entityStore.ts file for an entity
func GenerateEntityStore(baseDir, storesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the entity store file - ensure proper camelCase with capital preservation
	// First ensure pluralName is properly capitalized in PascalCase
	pascalPlural := ToPascalCase(pluralName)
	// Then convert to camelCase while preserving internal capitalization
	camelPlural := ToCamelCase(pascalPlural)
	
	// Log the store filename to debug
	filename := camelPlural + "Store.ts"
	fmt.Printf("Creating store file: %s\n", filename)
	
	outputPath := filepath.Join(storesDir, filename)

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

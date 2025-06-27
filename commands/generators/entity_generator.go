package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateEntityType generates the entity.ts file with the TypeScript interface
func GenerateEntityType(baseDir, storesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the entity type file
	// First ensure entityName is properly capitalized in PascalCase
	pascalEntity := ToPascalCase(entityName)
	// Then convert to camelCase while preserving internal capitalization
	camelEntity := ToCamelCase(pascalEntity)
	
	// Log the entity filename to debug
	filename := camelEntity + ".ts"
	fmt.Printf("Creating entity file: %s\n", filename)
	
	outputPath := filepath.Join(storesDir, filename)

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("entity.ts.tmpl")
	if err != nil {
		return err
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(storesDir, 0755); err != nil {
		return fmt.Errorf("error creating stores directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("entityType", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating entity.ts file: %v", err)
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
		return fmt.Errorf("error executing entity.ts template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

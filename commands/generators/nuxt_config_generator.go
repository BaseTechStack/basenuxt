package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateNuxtConfig generates the nuxt.config.ts file
func GenerateNuxtConfig(baseDir, rootDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the nuxt config
	outputPath := filepath.Join(rootDir, "nuxt.config.ts")

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("nuxt.config.ts.tmpl")
	if err != nil {
		return err
	}

	// Create the root directory if it doesn't exist
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		return fmt.Errorf("error creating root directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("nuxt_config", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating nuxt.config.ts file: %v", err)
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
		return fmt.Errorf("error executing nuxt.config.ts template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

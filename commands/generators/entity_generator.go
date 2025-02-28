package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateEntityType generates the entity.ts file with the TypeScript interface
func GenerateEntityType(baseDir, storesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the entity type file
	outputPath := filepath.Join(storesDir, strings.ToLower(entityName)+".ts")

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "entity.ts.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading entity.ts template: %v", err)
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(storesDir, 0755); err != nil {
		return fmt.Errorf("error creating stores directory: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("entityType").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing entity.ts template: %v", err)
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

package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateService generates the service file for an entity
func GenerateService(baseDir, servicesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the service
	outputPath := filepath.Join(servicesDir, fmt.Sprintf("%sService.ts", entityName))

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "entity_service.ts.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading Service template: %v", err)
	}

	// Create the services directory if it doesn't exist
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return fmt.Errorf("error creating services directory: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("service").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing Service template: %v", err)
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating Service file: %v", err)
	}
	defer file.Close()

	// Execute the template with data passed as a struct
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
		return fmt.Errorf("error executing Service template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

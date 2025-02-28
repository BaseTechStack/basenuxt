package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateComposable generates the composable for an entity
func GenerateComposable(baseDir, composablesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the composable
	outputPath := filepath.Join(composablesDir, fmt.Sprintf("use%s.ts", pluralName))

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "use_entities.ts.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading Composable template: %v", err)
	}

	// Create the composables directory if it doesn't exist
	if err := os.MkdirAll(composablesDir, 0755); err != nil {
		return fmt.Errorf("error creating composables directory: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("composable").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing Composable template: %v", err)
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating Composable file: %v", err)
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
		return fmt.Errorf("error executing Composable template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

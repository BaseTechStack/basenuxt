package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateGridCard generates the GridCard.vue component for an entity
func GenerateGridCard(baseDir, componentsDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the component
	outputPath := filepath.Join(componentsDir, "GridCard.vue")

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "grid_card.vue.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading GridCard template: %v", err)
	}

	// Create the component directory if it doesn't exist
	if err := os.MkdirAll(componentsDir, 0755); err != nil {
		return fmt.Errorf("error creating components directory: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("gridCard").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
		"computedTitle": func() string { return "{{ computedTitle }}" },
		"computedSubtitle": func() string { return "{{ computedSubtitle }}" },
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing GridCard template: %v", err)
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating GridCard file: %v", err)
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
		return fmt.Errorf("error executing GridCard template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

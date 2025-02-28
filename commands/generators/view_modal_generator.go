package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateViewModal generates the ViewModal.vue component for an entity
func GenerateViewModal(baseDir, componentsDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the component
	outputPath := filepath.Join(componentsDir, "ViewModal.vue")

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "view_modal.vue.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading ViewModal template: %v", err)
	}

	// Create the component directory if it doesn't exist
	if err := os.MkdirAll(componentsDir, 0755); err != nil {
		return fmt.Errorf("error creating components directory: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("viewModal").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
		"index": func(arr []Field, i int) Field {
			if i >= 0 && i < len(arr) {
				return arr[i]
			}
			return Field{}
		},
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing ViewModal template: %v", err)
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating ViewModal file: %v", err)
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
		return fmt.Errorf("error executing ViewModal template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

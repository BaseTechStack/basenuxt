package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GeneratePage generates the main page for an entity
func GeneratePage(baseDir, pagesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the page file
	outputPath := filepath.Join(pagesDir, "index.vue")

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "page_index.vue.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading page template: %v", err)
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(pagesDir, 0755); err != nil {
		return fmt.Errorf("error creating pages directory: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("page").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
		"formatDate": func(date string) string {
			return date
		},
		"index": func(arr []Field, i int) Field {
			if i >= 0 && i < len(arr) {
				return arr[i]
			}
			return Field{}
		},
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing page template: %v", err)
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating page file: %v", err)
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
		return fmt.Errorf("error executing page template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

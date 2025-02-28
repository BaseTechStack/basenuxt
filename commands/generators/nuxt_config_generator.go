package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateNuxtConfig generates the nuxt.config.ts file for an entity
func GenerateNuxtConfig(baseDir, entityDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the nuxt config
	outputPath := filepath.Join(entityDir, "nuxt.config.ts")

	// Path to the Go template file
	templatePath := filepath.Join(baseDir, "utils", "templates", "entity_templates", "nuxt.config.ts.tmpl")

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading nuxt.config.ts template: %v", err)
	}

	// Process the template with the Go templating engine
	tmpl, err := template.New("nuxtconfig").Funcs(template.FuncMap{
		"toLower":  strings.ToLower,
		"toUpper":  strings.ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"contains": strings.Contains,
	}).Parse(string(templateContent))

	if err != nil {
		return fmt.Errorf("error parsing nuxt.config.ts template: %v", err)
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating nuxt.config.ts file: %v", err)
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
		return fmt.Errorf("error executing nuxt.config.ts template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

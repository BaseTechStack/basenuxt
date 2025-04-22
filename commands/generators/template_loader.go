package generators

import (
	"fmt"
	"io/fs"
	"strings"
	"text/template"

	"github.com/BaseTechStack/basenuxt/templates"
)

// loadTemplate loads a template from the embedded filesystem
func loadTemplate(templateName string) ([]byte, error) {
	templateFS := templates.GetEntityTemplates()
	templateContent, err := fs.ReadFile(templateFS, templateName)
	if err != nil {
		return nil, fmt.Errorf("error reading template %s: %v", templateName, err)
	}
	return templateContent, nil
}

// createTemplate creates a template with common functions
func createTemplate(name string, content []byte) (*template.Template, error) {
	tmpl, err := template.New(name).Funcs(template.FuncMap{
		"toLower":  ToLower,
		"toUpper":  ToUpper,
		"toPascal": ToPascalCase,
		"toKebab":  ToKebabCase,
		"toSnake":  ToSnakeCase,
		"ToPascal": ToPascalCase,
		"ToKebab":  ToKebabCase,
		"ToSnake":  ToSnakeCase,
		"contains": Contains,
	}).Parse(string(content))

	if err != nil {
		return nil, fmt.Errorf("error parsing template %s: %v", name, err)
	}

	return tmpl, nil
}

// ToLower is a wrapper for strings.ToLower
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper is a wrapper for strings.ToUpper
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Contains is a wrapper for strings.Contains
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

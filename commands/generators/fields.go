package generators

import (
	"strings"
	"unicode"

	pluralize "github.com/gertd/go-pluralize"
)

// Field represents a struct field in the entity
type Field struct {
	Name       string
	Type       string
	JSONName   string
	IsRequired bool
}

// ParseFields parses a slice of field strings into Field structs
func ParseFields(fieldStrs []string) []Field {
	fields := make([]Field, 0, len(fieldStrs))

	for _, fieldStr := range fieldStrs {
		parts := strings.Split(fieldStr, ":")

		name := parts[0]
		fieldType := "string"
		isRequired := false

		if len(parts) > 1 {
			fieldType = parts[1]
		}

		if strings.HasSuffix(fieldType, "!") {
			isRequired = true
			fieldType = strings.TrimSuffix(fieldType, "!")
		}

		// Convert to camelCase for JSON field name
		jsonName := ToCamelCase(name)

		fields = append(fields, Field{
			Name:       name,
			JSONName:   jsonName,
			Type:       fieldType,
			IsRequired: isRequired,
		})
	}

	return fields
}

// PluralizeEntityName converts a singular name to its plural form
func PluralizeEntityName(name string) string {
	// Use the go-pluralize package for accurate pluralization
	pluralizer := pluralize.NewClient()
	
	// If the word is already plural, return it as is
	if pluralizer.IsPlural(name) {
		return name
	}
	
	// Otherwise, return the plural form
	return pluralizer.Plural(name)
}

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	// Split by common separators
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})

	// Convert each word to title case
	var result strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			result.WriteString(strings.ToUpper(word[:1]) + strings.ToLower(word[1:]))
		}
	}

	return result.String()
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	// First convert to pascal case to handle the initial casing
	pascal := ToPascalCase(s)

	// Convert pascal case to kebab case
	var result strings.Builder
	for i, r := range pascal {
		if unicode.IsUpper(r) && i > 0 {
			result.WriteRune('-')
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	pascal := ToPascalCase(s)
	return strings.ToLower(pascal[:1]) + pascal[1:]
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// First convert to pascal case to handle the initial casing
	pascal := ToPascalCase(s)

	// Convert pascal case to snake case
	var result strings.Builder
	for i, r := range pascal {
		if unicode.IsUpper(r) && i > 0 {
			result.WriteRune('_')
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

// ToPlural converts a singular string to its plural form
func ToPlural(s string) string {
	return PluralizeEntityName(s)
}

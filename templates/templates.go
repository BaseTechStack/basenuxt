package templates

import (
	"embed"
	"io/fs"
)

//go:embed entity_templates/*
var EntityTemplates embed.FS

// GetEntityTemplates returns the embedded entity templates filesystem
func GetEntityTemplates() fs.FS {
	subFS, _ := fs.Sub(EntityTemplates, "entity_templates")
	return subFS
}

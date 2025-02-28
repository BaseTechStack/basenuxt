package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// Generator holds configuration for the entity generation process
type Generator struct {
	BaseDir    string
	EntityName string
	PluralName string
	Fields     []Field
}

// NewGenerator creates a new generator instance
func NewGenerator(baseDir, entityName string, fieldStrs []string) *Generator {
	fields := ParseFields(fieldStrs)
	pluralName := PluralizeEntityName(entityName)

	return &Generator{
		BaseDir:    baseDir,
		EntityName: entityName,
		PluralName: pluralName,
		Fields:     fields,
	}
}

// Generate creates all entity files from templates
func (g *Generator) Generate() error {
	// Create the entity directory under baseDir/structures
	entityDirName := ToKebabCase(g.PluralName)
	entityDir := filepath.Join(g.BaseDir, "structures", entityDirName)

	// Create subdirectories following the clients template structure
	componentsDir := filepath.Join(entityDir, "components")
	composablesDir := filepath.Join(entityDir, "composables")
	pagesDir := filepath.Join(entityDir, "pages")
	storesDir := filepath.Join(entityDir, "stores")
	servicesDir := filepath.Join(entityDir, "services")

	// Create the directories if they don't exist
	dirsToCreate := []string{
		entityDir,
		componentsDir,
		composablesDir,
		pagesDir,
		storesDir,
		servicesDir,
	}

	for _, dir := range dirsToCreate {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	fmt.Printf("Generating entity files for %s in %s\n", g.EntityName, entityDir)

	// Generate TypeScript interface
	if err := GenerateEntityType(g.BaseDir, storesDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating entity type: %v", err)
	}

	// Generate Entity Store
	if err := GenerateEntityStore(g.BaseDir, storesDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating entity store: %v", err)
	}

	// Generate Vue components
	if err := GenerateAddModal(g.BaseDir, componentsDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating AddModal: %v", err)
	}

	if err := GenerateEditModal(g.BaseDir, componentsDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating EditModal: %v", err)
	}

	if err := GenerateViewModal(g.BaseDir, componentsDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating ViewModal: %v", err)
	}

	if err := GenerateGrid(g.BaseDir, componentsDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Grid: %v", err)
	}

	if err := GenerateGridCard(g.BaseDir, componentsDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating GridCard: %v", err)
	}

	if err := GenerateTable(g.BaseDir, componentsDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Table: %v", err)
	}

	// Generate Composable
	if err := GenerateComposable(g.BaseDir, composablesDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Composable: %v", err)
	}

	// Generate Service
	if err := GenerateService(g.BaseDir, servicesDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Service: %v", err)
	}

	// Generate Page
	if err := GeneratePage(g.BaseDir, pagesDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Page: %v", err)
	}

	// Generate nuxt.config.ts
	if err := GenerateNuxtConfig(g.BaseDir, entityDir, g.EntityName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating nuxt.config.ts: %v", err)
	}

	// Update main nuxt.config.ts to include the new entity
	if err := UpdateMainNuxtConfig(g.BaseDir, entityDirName); err != nil {
		return fmt.Errorf("error updating main nuxt.config.ts: %v", err)
	}

	fmt.Printf("Entity generation completed successfully for %s\n", g.EntityName)
	return nil
}

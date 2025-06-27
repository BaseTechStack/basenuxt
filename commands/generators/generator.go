package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// Generator holds configuration for the entity generation process
type Generator struct {
	BaseDir    string
	StructName string
	PluralName string
	Fields     []Field
}

// NewGenerator creates a new generator instance
func NewGenerator(baseDir, entityName string, fieldStrs []string) *Generator {
	fields := ParseFields(fieldStrs)
	pluralName := PluralizeEntityName(entityName)

	return &Generator{
		BaseDir:    baseDir,
		StructName: entityName,
		PluralName: pluralName,
		Fields:     fields,
	}
}

// Generate creates all entity files from templates
func (g *Generator) Generate() error {
	// Force consistent casing for entity names
	fmt.Println("Struct name: ", g.StructName)
	fmt.Println("Plural name: ", g.PluralName)
	g.StructName = ToPascalCase(g.StructName)
	g.PluralName = ToPascalCase(PluralizeEntityName(g.StructName))

	// Debug - print both entity name and plural name to verify proper capitalization
	fmt.Printf("Using entity name: %s\n", g.StructName)
	fmt.Printf("Using plural name: %s\n", g.PluralName)

	// Debug logging
	fmt.Printf("Struct name: %s, Plural name: %s\n", g.StructName, g.PluralName)

	// Convert plural name to kebab-case ensuring word boundaries are preserved with hyphens
	// This ensures "ProductCategories" becomes "product-categories"
	entityDirName := ToKebabCase(g.PluralName)

	fmt.Printf("Creating directory with kebab-case name: %s\n", entityDirName)

	// Create structures directory first
	structuresDir := filepath.Join(g.BaseDir, "structures")
	if err := os.MkdirAll(structuresDir, 0755); err != nil {
		return fmt.Errorf("error creating structures directory: %v", err)
	}

	// Place entity directory inside structures
	entityDir := filepath.Join(structuresDir, entityDirName)

	fmt.Printf("Generating entity in directory: %s\n", filepath.Join("structures", entityDirName))

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

	fmt.Printf("Generating entity files for %s in %s\n", g.StructName, entityDir)

	// Generate TypeScript interface
	if err := GenerateEntityType(g.BaseDir, storesDir, g.StructName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating entity type: %v", err)
	}

	// Generate Entity Store
	if err := GenerateEntityStore(g.BaseDir, storesDir, g.StructName, g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating entity store: %v", err)
	}

	// Generate Vue components
	if err := GenerateAddModal(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Add%sModal: %v", g.StructName, err)
	}

	if err := GenerateEditModal(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Edit%sModal: %v", g.StructName, err)
	}

	if err := GenerateViewModal(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating View%sModal: %v", g.StructName, err)
	}

	if err := GenerateDeleteModal(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Delete%sModal: %v", g.StructName, err)
	}

	if err := GenerateGrid(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating %sGrid: %v", g.StructName, err)
	}

	if err := GenerateGridCard(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating %sGridCard: %v", g.StructName, err)
	}

	if err := GenerateTable(g.BaseDir, componentsDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating %sTable: %v", g.StructName, err)
	}

	// Generate Composable
	if err := GenerateComposable(g.BaseDir, composablesDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Composable: %v", err)
	}

	// Generate Service
	if err := GenerateService(g.BaseDir, servicesDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Service: %v", err)
	}

	// Generate Page
	if err := GeneratePage(g.BaseDir, pagesDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Page: %v", err)
	}

	// Generate Single Page
	if err := GenerateSinglePage(g.BaseDir, pagesDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating Single Page: %v", err)
	}

	// Generate nuxt.config.ts
	if err := GenerateNuxtConfig(g.BaseDir, entityDir, ToPascalCase(g.StructName), g.PluralName, g.Fields); err != nil {
		return fmt.Errorf("error generating nuxt.config.ts: %v", err)
	}

	// Update main nuxt.config.ts to include the new entity
	if err := UpdateMainNuxtConfig(g.BaseDir, entityDirName); err != nil {
		return fmt.Errorf("error updating main nuxt.config.ts: %v", err)
	}

	// Update the sidebar to include a link to the new entity
	if err := UpdateSidebar(g.BaseDir, g.StructName, g.PluralName); err != nil {
		// Just log the error but don't fail the generation process
		fmt.Printf("Warning: could not update sidebar: %v\n", err)
	}

	fmt.Printf("Entity generation completed successfully for %s\n", g.StructName)
	return nil
}

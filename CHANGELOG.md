# Changelog

All notable changes to the BaseNuxt CLI tool will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.5.0] - 2025-06-27

### Changed
- Improved directory structure generation to place entities under a 'structures' directory
- Fixed single page generator to correctly create pages in [id]/index.vue format
- Updated import paths in templates to reference files from the structures directory
- Improved template rendering with proper Vue template syntax escaping
- Enhanced button styling in single entity page templates

### Fixed
- Fixed Vue template syntax issues with double curly braces in Go templates
- Fixed JSONName field access in templates to properly access entity properties
- Fixed Edit component references in page templates
- Improved compatibility with Nuxt UI component styling

## [v0.4.2] - 2025-06-27

### Fixed
- Fixed BaseService interface missing `fetchById` method causing TypeScript errors
- Fixed illogical conditional in ViewModal template (`props.entity.createdAt && !props.entity.createdAt`)
- Fixed function signature mismatch in handleUpdate (EditModal vs page component expectations)
- Fixed import path casing issues in all component templates (now uses proper camelCase)
- Simplified type casting in store.create calls with proper Omit types
- Fixed service method calls to properly destructure response objects (`result.item`)

### Improved
- Simplified ViewModal template logic by removing unnecessary conditional complexity
- Enhanced template consistency across all generated components
- Improved TypeScript compatibility with cleaner type definitions
- Generated code now follows core @admin/app/ architecture patterns

### Templates Updated
- `view_modal.vue.tmpl` - Simplified fetchData logic
- `page_id.vue.tmpl` - Fixed handleUpdate signature and service response handling  
- `page_index.vue.tmpl` - Added proper type casting for store operations
- `edit_modal.vue.tmpl`, `add_modal.vue.tmpl`, `table.vue.tmpl`, `grid.vue.tmpl` - Fixed import casing
- `entity_service.ts.tmpl` - Enhanced service response handling

### Added
- Added new version release process

## [v0.4.1] - 2025-06-27

### Fixed
- Fixed file naming inconsistencies across generated code
- Added proper camelCase support for model files and imports
- Fixed entity naming to preserve PascalCase for multi-word entities
- Improved consistency between directory names and API endpoints

## [v0.4.0] - 2025-06-27

### Changed
- Standardized API endpoint format to consistently use kebab-case plural naming
- Fixed inconsistent templates that mixed toLower and toKebab for API endpoints
- Ensured consistent code generation for models, services, and stores
- Simplified API endpoint generation for product structure
- Improved entity code pattern following Nuxt/Vue/TypeScript conventions

### Fixed
- Fixed API endpoint format inconsistencies in templates
- Fixed incorrect template function reference

## [v0.3.1] - 2025-06-10

### Changed
- Renamed command from `basenuxt` to `bux` (BaseUX) with backward compatibility
- Added deprecation notice when using the old `basenuxt` command name
- Simplified command structure: `bux start` or `bux s` now exclusively runs development server
- Added new `bux run <command>` to directly pass any command to `bun run`
- Improved command documentation in README

## [v0.3.0] - 2025-05-26

### Added
- Added complete entity lifecycle management with sidebar integration
- Added automatic sidebar navigation updates for new entities
- Added functions to remove entities from sidebar when destroyed
- Improved store pattern with Pinia integration

### Changed
- Changed store implementation to use composition with Pinia's defineStore
- Updated entity-specific stores to follow consistent pattern
- Modified components to use entity-specific naming conventions to avoid auto-import conflicts
- Updated service implementations from singleton classes to composables

### Fixed
- Fixed prop naming consistency between different entity components
- Fixed pluralization issues in templates by using PluralName consistently
- Fixed import paths in components to use correct store names
- Fixed ViewModal template to address template parsing errors
- Fixed pagination implementation to explicitly pass page parameters
- Fixed DeleteModal component to use v-model:open for Nuxt UI 3 compatibility
- Fixed Pinia store initialization to prevent "getActivePinia() was called but there was no active Pinia" errors

## [v0.2.0] - 2025-04-22

### Added
- Added comprehensive APPNAME replacement throughout entire project
- Added pre-configured GitHub Actions workflow template in 'github/workflows' (without leading dot)
- Added detailed section in README about activating GitHub Actions
- Added support for API URL configuration in runtime config with project name

### Changed
- Improved project template with more informative README
- Enhanced file replacement utility to handle all text file formats
- Renamed .github directory to github to prevent accidental CI/CD activation

### Fixed
- Fixed file filtering to properly process github directory files
- Fixed case sensitivity issues when deploying to Linux environments
- Fixed import paths in store templates to use consistent snake_case naming
- Updated process.client references to import.meta.client for better TypeScript compatibility

## [v0.1.8] - 2025-04-22

### Added
- Added toSnake template function to all templates for consistent file naming
- Added support for API URL configuration in runtime config

### Changed
- Updated import paths in templates to use snake_case for better Linux compatibility
- Changed entity template file naming to use consistent snake_case (e.g., loan_entry_service.ts)
- Updated component imports to match the new file naming pattern

### Fixed
- Fixed case sensitivity issues when deploying to Linux environments
- Fixed import paths in store templates to correctly reference snake_case file names
- Updated process.client references to import.meta.client for better TypeScript compatibility
- Fixed APPNAME replacement to properly process .github directory files

## [v0.1.7] - 2025-03-05

### Added
- Added factory methods in entity.ts.tmpl for mapping JSON data to entity models
- Implemented fromJson and fromJsonList methods to handle both snake_case and camelCase properties
- Added toJson method to convert model instances back to JSON format

### Changed
- Updated entity_service.ts.tmpl to use the new factory methods instead of standalone utility functions
- Modified entities_store.ts.tmpl to use model conversion for API requests
- Updated view_modal.vue.tmpl to automatically convert raw data to model instances

### Fixed
- Fixed slot naming in view_modal.vue.tmpl to use #body instead of #content for USlideover component
- Fixed inconsistent pagination options in page_index.vue.tmpl
- Fixed event emission in add_modal.vue.tmpl to use entity-specific event names

## [v0.1.6] - 2025-03-03

### Added
- Added DeleteModal generator to create entity-specific delete modal components
- Updated templates to use entity-specific naming conventions for all components
- Improved UI for modal components with consistent styling

### Changed
- Modified delete_modal.vue.tmpl to use a red-themed UI for better user experience
- Updated view_modal.vue.tmpl to include title and description properties

### Fixed
- Fixed ViewModal generator to correctly pass StructName to the template

## [v0.1.5] - 2025-03-03

### Fixed
- Fixed ViewModal generator to correctly use StructName field instead of EntityName
- Fixed Grid and GridCard generators to use StructName field instead of EntityName
- Updated all generator functions to consistently use StructName throughout the codebase

## [v0.1.4] - 2025-03-03

### Fixed
- Fixed EditModal generator to correctly use StructName field instead of EntityName

## [v0.1.3] - 2025-03-03

### Fixed
- Fixed version checking to use the correct repository URLs
- Updated command references from 'base' to 'basenuxt' in upgrade messages
- Fixed EditModal generator to correctly use StructName field instead of EntityName

## [v0.1.2] - 2025-03-03

### Added
- Automatic dependency installation when creating new projects
- Support for Bun, Yarn, and npm package managers (in order of preference)

### Fixed
- Fixed generator function names to follow consistent naming conventions
- Corrected output paths for generated files
- Updated version number in code to match release version

## [v0.1.1] - 2025-03-03

### Fixed
- Fixed generator function names to follow consistent naming conventions
- Corrected output paths for generated files

## [v0.1.0] - 2025-03-03

### Added
- Initial release of BaseNuxt CLI
- Consistent naming convention across all files and references
- Improved installation scripts for macOS, Linux, and Windows
- Local installation script for development and testing
- Release script that builds binaries for multiple platforms

### Changed
- Renamed all references from "Base" to "BaseNuxt" for consistency
- Updated installation directory from `.base` to `.basenuxt`
- Updated binary name from `base` to `basenuxt`

### Fixed
- Fixed installation script to use correct naming conventions
- Fixed release script to handle existing tags
- Added fallback to version 0.1.0 if no version is found

## [v0.0.1] - 2025-03-01

### Added
- Complete entity generation system with:
  - Grid and Table components for data display
  - Add/Edit/View modals with form validation
  - Entity type and store generation
  - API composables for data fetching
  - Service layer that implements BaseService interface
  - Nuxt configuration generation and integration
- CLI commands for scaffolding new projects and entities
- Support for various field types (string, number, boolean, etc.)
- Automatic extension of main nuxt.config.ts when generating entities

### Changed
- Updated to use Vue 3 and Nuxt 3
- Updated to use TypeScript for better type safety
- Updated to use Pinia for state management
- Updated to use Vite for faster development
- Updated to use Tailwind CSS for styling

### Fixed
- Fixed issue with entity generation
 
[v0.1.7]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.7
[v0.1.6]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.6
[v0.1.5]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.5
[v0.1.4]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.4
[v0.1.3]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.3
[v0.1.2]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.2
[v0.1.1]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.1
[v0.1.0]: https://github.com/BaseTechStack/bux/releases/tag/v0.1.0
 
# Changelog

All notable changes to the BaseNuxt CLI tool will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
 
[v0.1.7]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.7
[v0.1.6]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.6
[v0.1.5]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.5
[v0.1.4]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.4
[v0.1.3]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.3
[v0.1.2]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.2
[v0.1.1]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.1
[v0.1.0]: https://github.com/BaseTechStack/basenuxt/releases/tag/v0.1.0
 
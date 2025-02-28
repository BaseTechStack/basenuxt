# Changelog

All notable changes to the BaseNuxt CLI tool will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.0.0] - 2025-03-01

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
 
[v1.0.0]: https://github.com/BaseTechStack/basenuxt/releases/tag/v1.0.0

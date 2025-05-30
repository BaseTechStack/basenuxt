# BaseUX (bux) - Command Line Tool for Nuxt Application Scaffolding

BaseUX (bux) is a powerful command-line tool designed to streamline development with Nuxt.js applications.
It offers scaffolding, module generation, and utilities to accelerate Vue/Nuxt development.

> **Note:** The command name has changed from `basenuxt` to `bux`. The old name is still supported but deprecated.

## Installation

### macOS and Linux

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/bux/main/install.sh | bash
```

If you need to install in a protected directory (like `/usr/local/bin`), use:

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/bux/main/install.sh | sudo bash
```

### Windows

#### Option 1: Using PowerShell (Recommended)

1. Open PowerShell as Administrator
2. Run:
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/BaseTechStack/bux/main/install.ps1'))
```

#### Option 2: Using Git Bash

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/bux/main/install.sh | bash
```

## Commands

### `bux new <project-name>`

Create a new Nuxt.js project with the BaseNuxt framework.

```bash
bux new myapp
```

### `bux generate` or `bux g`

Generate a new entity module with fields.

```bash
bux g <entity-name> [field:type ...] [options]
```

### `bux start` or `bux s`

Start the development server using Bun.

Examples:
```bash
# Start the development server
bux start

# Using the shorthand alias
bux s
```

### `bux run <command>`

Run any Bun script command.

Examples:
```bash
# Build the application
bux run build

# Generate static site
bux run generate

# Run linter
bux run lint

# Run linter with auto-fix
bux run lint:fix

# Run preview server
bux run preview
```

### `bux update`

Update framework core components:

```bash
bux update
```

### `bux upgrade`

Upgrade the BaseNuxt CLI tool:

```bash
bux upgrade
```

### `bux version`

Display version information:

```bash
bux version
```

## Entity Generation

BaseNuxt includes a sophisticated system for generating entity modules with full CRUD capabilities:

1. Each entity gets its own module with:
   - Vue components (Grid, Table, Add/Edit/View modals)
   - TypeScript stores
   - API composables
   - Service layer for interacting with backend APIs
   - Nuxt module configuration

2. The system supports various field types:
   - string: Text fields
   - number: Numeric fields  
   - boolean: True/false fields
   - date: Date fields
   - enum: Predefined options

3. Generated UI features:
   - Responsive Grid and Table views
   - Search functionality
   - Pagination
   - Create/Edit/View modals
   - Form validation

## Framework Features

BaseNuxt provides a robust foundation for Nuxt applications:

- **Vue 3 & Composition API**: Modern Vue features
- **Nuxt 3**: Nuxt 4 ready.
- **Tailwind CSS**: CSS framework
- **TypeScript**: Type-safe development
- **NuxtUI Components**: Integrated UI component library
- **Authentication**: Ready-to-use auth system
- **State Management**: Pinia store integration
- **API Integration**: Composables for API interactions
- **Module System**: Extendable architecture via Nuxt Layers.

## Field Types

BaseNuxt supports various field types for model generation:

Basic Types:
- `string`: String field
- `int`: Integer field
- `bool`: Boolean field
- `float`: Float field
- `text`: Text field (for longer strings)

  

## Contributing

Contributions to BaseNuxt are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

BaseNuxt is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

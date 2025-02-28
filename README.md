# BaseNuxt - Command Line Tool for Nuxt Application Scaffolding

BaseNuxt is a powerful command-line tool designed to streamline development with Nuxt.js applications.
It offers scaffolding, module generation, and utilities to accelerate Vue/Nuxt development.

## Installation

### macOS and Linux

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/basenuxt/main/install.sh | bash
```

If you need to install in a protected directory (like `/usr/local/bin`), use:

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/basenuxt/main/install.sh | sudo bash
```

### Windows

#### Option 1: Using PowerShell (Recommended)

1. Open PowerShell as Administrator
2. Run:
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/BaseTechStack/basenuxt/main/install.ps1'))
```

#### Option 2: Using Git Bash

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/basenuxt/main/install.sh | bash
```

## Commands

### `basenuxt new <project-name>`

Create a new Nuxt.js project with the BaseNuxt framework.

```bash
basenuxt new myapp
```

### `basenuxt generate` or `basenuxt g`

Generate a new entity module with fields.

```bash
basenuxt g <entity-name> [field:type ...] [options]
```

### `basenuxt start` or `basenuxt s`

Start the BaseNuxt application server.

Options:
- `--hot-reload`, `-r`: Enable hot reloading using air
- `--docs`, `-d`: Generate Swagger documentation

Examples:
```bash
# Start the server normally
basenuxt start

# Start with hot reloading
basenuxt start -r

# Start with Swagger documentation
basenuxt start -d

# Start with both hot reloading and Swagger docs
basenuxt start -r -d
```

### `basenuxt update`

Update framework core components:

```bash
basenuxt update
```

### `basenuxt upgrade`

Upgrade the BaseNuxt CLI tool:

```bash
basenuxt upgrade
```

### `basenuxt version`

Display version information:

```bash
basenuxt version
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

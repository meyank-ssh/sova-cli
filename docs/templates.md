# Project Templates

This document describes the available project templates and their structure in Sova CLI.

## API Template

The API template creates a Go web service with a clean architecture structure.

### Directory Structure
```
📦 project/
├── cmd/           # Application entry point
├── internal/      # Private application code
│   ├── handlers/  # HTTP handlers
│   ├── middleware/# Middleware components
│   ├── models/    # Data models
│   ├── server/    # Server implementation
│   └── service/   # Service layer
├── pkg/          # Public libraries
├── api/          # API definitions
├── routes/       # Route definitions
├── docs/         # Documentation
└── scripts/      # Build scripts
```

### Features
- Clean architecture structure
- HTTP server using Gin framework
- Environment configuration with .env
- Docker support with docker-compose
- Optional integrations:
  - PostgreSQL database
  - Redis cache
  - RabbitMQ message queue
  - Zap logging middleware

### Docker Services
When enabled, the following services are available:
- PostgreSQL (port: 5432)
- Redis (port: 6379)
- RabbitMQ (ports: 5672, 15672)

### Configuration
- Environment variables in `.env`
- Docker volumes for data persistence
- Customizable service configurations

## CLI Template

The CLI template creates a command-line application using Cobra.

### Directory Structure
```
📦 project/
├── cmd/
│   ├── root/             # Root command
│   └── version/          # Version command
├── internal/
│   ├── commands/         # Command implementations
│   ├── config/          # Configuration
│   └── utils/           # Utility functions
├── pkg/                 # Public packages
├── docs/               # Documentation
├── scripts/            # Build and deployment scripts
└── tests/              # Integration tests
```

### Features
- Cobra-based CLI structure
- Command management
- Configuration handling
- Utility functions for CLI operations

## Common Features

Both templates include:
- Go modules support
- `.gitignore` with appropriate exclusions
- Documentation structure
- Test setup
- Build scripts

## Recent Updates

1. Fixed Import Paths
   - Moved routes to `internal/routes`
   - Updated import paths in templates
   - Fixed module name references

2. Docker Compose
   - Removed obsolete version attribute
   - Added volume configurations
   - Improved service definitions

3. Project Structure
   - Reorganized internal packages
   - Added consistent directory structure
   - Improved template organization

4. Git Configuration
   - Added comprehensive `.gitignore` templates
   - Separate configurations for API and CLI projects
   - Docker-specific ignores for API projects

## Usage

Create a new API project:
```bash
sova-cli create api my-project
```

Create a new CLI project:
```bash
sova-cli create cli my-project
```

## Configuration Options

### API Projects
- `UsePostgres`: Enable PostgreSQL support
- `UseRedis`: Enable Redis support
- `UseRabbitMQ`: Enable RabbitMQ support
- `UseZap`: Enable Zap logging middleware

### CLI Projects
- Basic CLI structure with extensible commands
- Configuration management with Viper

## Creating Custom Templates

1. Create a template directory:
   ```bash
   mkdir -p ~/.sova/templates/my-template
   ```

2. Add template files:
   ```bash
   my-template/
   ├── template.yaml   # Template configuration
   ├── files/         # Template files
   └── hooks/         # Custom scripts
   ```

3. Template Configuration (template.yaml):
   ```yaml
   name: my-template
   description: My custom template
   version: 1.0.0
   files:
     - source: files/main.go
       target: cmd/main.go
     - source: files/config.go
       target: internal/config/config.go
   ```

4. Use your template:
   ```bash
   sova init my-project --template my-template
   ```

## Template Variables

Available variables in templates:

- `{{.ProjectName}}` - Project name
- `{{.Description}}` - Project description
- `{{.Author}}` - Author name
- `{{.Year}}` - Current year
- `{{.GoVersion}}` - Go version
- `{{.License}}` - License type

## Examples

1. **Custom main.go**:
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Welcome to {{.ProjectName}}!")
   }
   ```

2. **Custom README.md**:
   ```markdown
   # {{.ProjectName}}

   {{.Description}}

   ## Author
   {{.Author}}

   ## License
   {{.License}} © {{.Year}}
   ``` 
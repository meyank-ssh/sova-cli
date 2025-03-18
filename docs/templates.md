# Templates Guide

Sova CLI comes with two built-in templates to help you kickstart your projects.

## Available Templates

### 1. API Template
Complete structure for API applications.

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

### 2. CLI Template
Structure for command-line applications.

```
📦 project/
├── cmd/
│   ├── root/     # Root command
│   └── commands/ # Subcommands
├── internal/
│   ├── commands/ # Command implementations
│   ├── config/   # Configuration
│   └── utils/    # Utility functions
└── docs/         # Documentation
```

## Features

### API Template Features
- Complete API project structure
- Built-in middleware (logging, CORS, etc.)
- Service layer with PostgreSQL, Redis, and RabbitMQ support
- Environment configuration
- Docker support
- API documentation structure

### CLI Template Features
- Cobra-based CLI structure
- Command organization
- Configuration management
- Utility functions
- Documentation structure

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
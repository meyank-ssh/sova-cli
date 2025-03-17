# Templates Guide

Sova CLI comes with several built-in templates to help you kickstart your projects.

## Available Templates

### 1. Default Template
Basic Go project structure with essential directories.

```
📦 project/
├── cmd/           # Command-line interfaces
├── internal/      # Private application code
├── pkg/          # Public libraries
├── api/          # API definitions
├── docs/         # Documentation
└── scripts/      # Build scripts
```

### 2. Web Template
Complete structure for web applications.

```
📦 project/
├── cmd/
│   └── server/   # Server entry point
├── internal/
│   ├── handlers/ # HTTP handlers
│   ├── middleware/
│   ├── models/   # Data models
│   └── db/       # Database layer
├── web/
│   ├── templates/
│   └── static/
└── docs/
```

### 3. CLI Template
Structure for command-line applications.

```
📦 project/
├── cmd/
│   ├── root/     # Root command
│   └── commands/ # Subcommands
├── internal/
│   └── config/   # Configuration
└── docs/
```

### 4. Library Template
Structure for Go libraries.

```
📦 project/
├── pkg/          # Public API
├── internal/     # Private code
├── examples/     # Usage examples
└── docs/         # Documentation
```

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
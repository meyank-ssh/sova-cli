# 🚀 Sova CLI

A powerful and modern CLI tool for scaffolding projects with best practices and optimal structure. Sova CLI helps you jumpstart your development by generating production-ready project templates.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)

## ✨ Features

- 🏗️ Multiple project templates (Go Web, CLI, Library)
- 📁 Standardized project structure
- 🔧 Customizable templates
- 🚦 Built-in testing setup
- 📚 Automatic documentation generation
- 🛠️ Development tools integration

## 🔧 Installation

### Using Go Install

```bash
go install github.com/meyank/sova-cli@latest
```

### From Source

```bash
git clone https://github.com/meyank/sova-cli.git
cd sova-cli
go build
```

## 🚀 Quick Start

1. Create a new project:
   ```bash
   sova init my-awesome-project
   ```

2. Choose a template:
   ```bash
   sova init my-web-app --template go-web
   ```

3. Generate components:
   ```bash
   sova generate controller User
   ```

## 📖 Available Commands

### Project Initialization
```bash
# Basic project
sova init project-name

# Web project
sova init project-name --template go-web

# CLI project
sova init project-name --template cli

# Library project
sova init project-name --template library

# Force overwrite existing directory
sova init project-name --force
```

### Component Generation
```bash
# Generate a new controller
sova generate controller UserController

# Generate a model
sova generate model User

# Generate an API endpoint
sova generate api UserAPI
```

### Other Commands
```bash
# Show version
sova version

# Show verbose version info
sova version --verbose

# Show help
sova help
```

## 📁 Project Templates

### Default Template
```
📦 project/
├── cmd/           # Command-line interfaces
├── internal/      # Private application code
├── pkg/          # Public libraries
├── api/          # API definitions
├── docs/         # Documentation
├── scripts/      # Build and maintenance scripts
└── test/         # Additional test files
```

### Web Template
```
📦 project/
├── cmd/          # Entry points
├── internal/     # Private application code
│   ├── handlers/ # HTTP handlers
│   ├── models/   # Data models
│   └── db/       # Database interactions
├── pkg/          # Public libraries
├── web/         # Web-specific code
│   ├── templates/# HTML templates
│   ├── static/   # Static assets
│   └── routes/   # Route definitions
└── docs/         # Documentation
```

### CLI Template
```
📦 project/
├── cmd/          # CLI commands
│   ├── root/     # Root command
│   └── commands/ # Subcommands
├── internal/     # Private application code
├── pkg/          # Public libraries
└── docs/         # Documentation
```

## 🛠️ Development

### Prerequisites
- Go 1.21 or higher
- Git

### Building from Source
```bash
# Clone the repository
git clone https://github.com/meyank/sova-cli.git

# Change to project directory
cd sova-cli

# Install dependencies
go mod download

# Build the project
go build

# Run tests
go test ./...
```

### Adding Custom Templates

1. Create a new template in `templates/` directory
2. Register the template in `internal/project/structure.go`
3. Add template-specific logic in `internal/templates/`

## 📚 Documentation

- [API Documentation](docs/API.md)
- [Template Guide](docs/TEMPLATES.md)
- [Contributing Guide](CONTRIBUTING.md)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- The Go community for inspiration and support

## 📞 Support

- Create an issue for bug reports
- Start a discussion for feature requests
- Check our [FAQ](docs/FAQ.md) for common questions 
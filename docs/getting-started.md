# Getting Started with Sova CLI

Welcome to Sova CLI! This guide will help you get up and running quickly.

## Installation

### Prerequisites
- Go 1.21 or higher
- Git (for development)

### Install Methods

1. **Quick Install (Linux/macOS)**:
   ```bash
   curl -fsSL https://raw.githubusercontent.com/meyanksingh/go-sova/master/scripts/install.sh | bash
   ```

2. **Go Install**:
   ```bash
   go install github.com/meyanksingh/go-sova@latest
   ```

3. **Manual Installation**:
   - Download from [Releases](https://github.com/meyanksingh/go-sova/releases)
   - Extract and add to your PATH

## First Steps

1. **Create Your First Project**:
   ```bash
   sova init my-project
   cd my-project
   ```

2. **Project Structure**:
   ```
   my-project/
   ├── cmd/
   ├── internal/
   ├── pkg/
   ├── docs/
   └── README.md
   ```

3. **Generate Components**:
   ```bash
   sova generate controller UserController
   sova generate model User
   ```

## Project Templates

1. **Web Application**:
   ```bash
   sova init my-web --template web
   ```

2. **CLI Application**:
   ```bash
   sova init my-cli --template cli
   ```

3. **Library**:
   ```bash
   sova init my-lib --template library
   ```

## Next Steps

- Check out the [Templates Guide](templates.md)
- Learn about [Configuration](configuration.md)
- Read our [Contributing Guide](../CONTRIBUTING.md) 
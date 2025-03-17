# Sova CLI

Sova CLI is a powerful tool for initializing and generating project boilerplate code. It helps you quickly set up new projects with predefined templates and structures.

## Installation

```bash
go install github.com/meyank/sova-cli@latest
```

## Usage

### Initialize a new project

```bash
sova init [project-name]
```

### Generate project components

```bash
sova generate [component]
```

### Check version

```bash
sova version
```

### Get help

```bash
sova help
```

Or simply:

```bash
sova
```

## Project Structure

```
📦 sova-cli/                   # Root of your CLI project
├── 📂 cmd/                    # CLI commands
│   ├── init.go                # `sova init` command
│   ├── generate.go            # `sova generate` command
│   ├── version.go             # `sova version` command
│   ├── root.go                # Root command (entry point for Cobra)
│
├── 📂 internal/               # Business logic (not exposed externally)
│   ├── project/               # Project initialization logic
│   │   ├── create.go
│   │   ├── structure.go
│   │   ├── template_loader.go
│   ├── templates/             # Manages boilerplate template loading
│   │   ├── loader.go
│   │   ├── files.go
│   ├── utils/                 # Utility functions
│   │   ├── file_utils.go
│   │   ├── input_reader.go
│   │   ├── logger.go
│
├── 📂 templates/              # Predefined boilerplate templates
│   ├── go-main.tpl            # Go main.go template
│   ├── go-mod.tpl             # Go module template
│   ├── readme.tpl             # README template
│
├── 📂 tests/                  # Unit & integration tests
│
├── main.go                    # CLI entry point
├── go.mod                     # Go module file
├── README.md                  # Project documentation
├── LICENSE                    # License file
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
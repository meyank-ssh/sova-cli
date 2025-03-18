# Configuration Guide

Sova CLI can be configured using configuration files and command-line flags.

## Global Configuration

Create `.sova.yaml` in your home directory:

```yaml
# Default settings for new projects
defaults:
  template: web
  license: MIT
  goVersion: "1.21"
  author: "Meyank Singh"

# Template settings
templates:
  directory: ~/.sova/templates
  default: web

# Project settings
project:
  structure:
    enableTests: true
    enableDocs: true
    enableScripts: true

# Tool settings
tools:
  enableFormatting: true
  enableLinting: true
  enableTesting: true
```

## Project Configuration

Create `.sova.yaml` in your project directory:

```yaml
# Project information
name: my-awesome-project
version: 1.0.0
description: A fantastic Go project
author: Your Name
license: MIT

# Build settings
build:
  main: ./cmd/main.go
  output: ./bin/app
  ldflags: -s -w

# Dependencies
dependencies:
  - github.com/spf13/cobra
  - github.com/spf13/viper

# Development tools
tools:
  formatter: gofmt
  linter: golangci-lint
  testRunner: go test
```

## Environment Variables

Sova CLI respects the following environment variables:

```bash
# Configuration
SOVA_CONFIG=/path/to/config.yaml
SOVA_TEMPLATE_DIR=~/.sova/templates

# Project defaults
SOVA_DEFAULT_TEMPLATE=web
SOVA_DEFAULT_LICENSE=MIT
SOVA_DEFAULT_AUTHOR="Meyank Singh"

# Development
SOVA_DEBUG=true
SOVA_VERBOSE=true
```

## Command Line Flags

Global flags available for all commands:

```bash
# General
--config string     Config file path
--verbose          Enable verbose output
--debug           Enable debug mode

# Project initialization
--template string  Template to use
--force           Force overwrite existing files
--no-git          Don't initialize git repository

# Component generation
--output string    Output directory
--dry-run         Show what would be done
```

## Template Configuration

Template-specific configuration in `template.yaml`:

```yaml
name: custom-template
version: 1.0.0
description: Custom project template

# Files to include
files:
  - source: main.go.tmpl
    target: cmd/main.go
  - source: config.go.tmpl
    target: internal/config/config.go

# Directories to create
directories:
  - cmd
  - internal
  - pkg
  - docs

# Dependencies to add
dependencies:
  - name: github.com/spf13/cobra
    version: v1.7.0
  - name: github.com/spf13/viper
    version: v1.16.0

# Hooks
hooks:
  pre-generate:
    - command: go mod init
  post-generate:
    - command: go mod tidy
``` 
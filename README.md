# 🚀 Sova CLI

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)](https://golang.org/dl/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-sova/sova-cli)](https://goreportcard.com/report/github.com/go-sova/sova-cli)

A modern CLI tool for scaffolding Go projects with best practices. Generate production-ready project templates in seconds.

## 🚀 Quick Install

**Linux/macOS**:
```bash
curl -fsSL https://raw.githubusercontent.com/meyanksingh/go-sova/master/scripts/install.sh | bash
```

**Windows** (PowerShell Admin):
```powershell
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/meyanksingh/go-sova/master/scripts/install.sh" -OutFile "install.sh"; bash install.sh
```

**Using Go**:
```bash
go install github.com/go-sova/sova-cli@latest
```

**Manual Installation**:
Download the latest release from [GitHub Releases](https://github.com/go-sova/sova-cli/releases/latest)

## 💡 Usage

Create a new project:
```bash
# Basic project
sova init my-project

# Web project
sova init my-web --template web

# CLI project
sova init my-cli --template cli
```

## 📦 Features

- Multiple project templates (Web, CLI, Library)
- Standardized project structure
- Customizable templates
- Interactive prompts
- Cross-platform support

## 📚 Documentation

- [Getting Started](https://github.com/go-sova/sova-cli/wiki/getting-started)
- [Templates](https://github.com/go-sova/sova-cli/wiki/templates)
- [Configuration](https://github.com/go-sova/sova-cli/wiki/configuration)
- [Contributing](CONTRIBUTING.md)

## 🤝 Contributing

We love your input! Check out our [Contributing Guide](CONTRIBUTING.md) for ways to get started. Every contribution counts:

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing`)
5. Open a Pull Request

## 📝 License

Copyright © 2024 [Sova CLI Contributors](https://github.com/go-sova/sova-cli/graphs/contributors)

This project is [MIT](LICENSE) licensed. By contributing, you agree that your contributions will be licensed under its MIT License. 
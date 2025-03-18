# {{.ProjectName}}

{{.ProjectDescription}}

## Installation

```bash
go get -u github.com/yourusername/{{.ProjectName}}
```

## Usage

```bash
{{.ProjectName}} [command]
```

### Available Commands:

* command1: Description of command1
* command2: Description of command2
* help: Help about any command

### Flags:

* --config string: config file (default is $HOME/.{{.ProjectName}}.yaml)
* -h, --help: help for {{.ProjectName}}
* -t, --toggle: Help message for toggle

Use "{{.ProjectName}} [command] --help" for more information about a command.

## Development

1. Clone the repository
2. Install dependencies with `go mod tidy`
3. Run with `go run main.go`

## Building

Build a binary with:

```bash
go build -o {{.ProjectName}}
```

## License

This project is licensed under the {{.License}} License - see the LICENSE file for details. 
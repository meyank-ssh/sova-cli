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
   go install github.com/go-sova/sova-cli@latest
   ```

3. **Manual Installation**:
   - Download from [Releases](https://github.com/go-sova/sova-cli/releases)
   - Extract and add to your PATH

## Quick Start

### Creating an API Project

1. Create a new API project:
```bash
sova-cli create api my-api
```

2. Choose your integrations when prompted:
- PostgreSQL database
- Redis cache
- RabbitMQ message queue
- Zap logging

3. Navigate to your project:
```bash
cd my-api
```

4. Start the services:
```bash
docker compose up -d
```

5. Run your application:
```bash
go run cmd/main.go
```

Your API will be available at `http://localhost:8080`

### Creating a CLI Project

1. Create a new CLI project:
```bash
sova-cli create cli my-cli
```

2. Navigate to your project:
```bash
cd my-cli
```

3. Build your CLI:
```bash
go build -o my-cli cmd/main.go
```

4. Run your CLI:
```bash
./my-cli --help
```

## Project Structure

### API Project Structure
```
my-api/
├── cmd/main.go              # Entry point
├── internal/               # Internal packages
│   ├── config/            # Configuration
│   ├── handlers/          # HTTP handlers
│   ├── middleware/        # HTTP middleware
│   ├── models/           # Data models
│   ├── routes/           # Route definitions
│   ├── server/           # Server setup
│   └── service/          # Business logic
├── docker-compose.yml     # Docker services
└── .env                   # Environment variables
```

### CLI Project Structure
```
my-cli/
├── cmd/
│   ├── root/            # Root command
│   └── version/         # Version command
└── internal/
    ├── commands/        # Command implementations
    ├── config/         # Configuration
    └── utils/          # Utilities
```

## Configuration

### API Project

1. Environment Variables (`.env`):
```env
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/my-api
REDIS_URL=localhost:6379
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

2. Docker Services (`docker-compose.yml`):
- PostgreSQL (port: 5432)
- Redis (port: 6379)
- RabbitMQ (ports: 5672, 15672)

### CLI Project

Configuration is managed through:
- Command-line flags
- Configuration files
- Environment variables

## Development

### API Development

1. Start dependencies:
```bash
docker compose up -d
```

2. Run with hot reload (using air):
```bash
air
```

3. Access endpoints:
- Health check: `GET http://localhost:8080/api/health`
- Ping: `GET http://localhost:8080/api/ping`

### CLI Development

1. Add new commands:
```bash
sova-cli add command my-command
```

2. Build and test:
```bash
go build -o my-cli cmd/main.go
./my-cli my-command
```

## Testing

Run tests:
```bash
go test ./...
```

## Common Issues

1. Docker services not starting:
   - Check if Docker is running
   - Verify ports are not in use
   - Check docker-compose.yml configuration

2. Import path issues:
   - Verify module name in go.mod
   - Check import paths in source files
   - Run `go mod tidy`

3. Build errors:
   - Run `go mod tidy`
   - Check Go version compatibility
   - Verify all dependencies are installed

## Next Steps

1. Customize your project:
   - Add new routes (API)
   - Create new commands (CLI)
   - Implement business logic

2. Deploy your application:
   - Build for production
   - Set up CI/CD
   - Configure production environment

3. Explore advanced features:
   - Custom middleware
   - Additional integrations
   - Extended configuration
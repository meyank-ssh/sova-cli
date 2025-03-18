# Server Configuration
PORT=8080

{{if .UsePostgres}}# Database Configuration
DATABASE_URL=postgres://postgres:postgres@localhost:5432/{{.ProjectName}}?sslmode=disable
{{end}}

{{if .UseRedis}}# Redis Configuration
REDIS_URL=localhost:6379
{{end}}

{{if .UseRabbitMQ}}# RabbitMQ Configuration
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
{{end}} 
version: '3'
services:
  {{if .UsePostgres}}postgres:
    image: postgres:latest
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB={{.ProjectName}}
  {{end}}

  {{if .UseRedis}}redis:
    image: redis:latest
    ports:
      - "6379:6379"
  {{end}}

  {{if .UseRabbitMQ}}rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"
  {{end}} 
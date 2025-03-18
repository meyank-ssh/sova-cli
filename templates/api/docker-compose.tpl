services:
  {{if .UsePostgres}}postgres:
    image: postgres:latest
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB={{.ProjectName}}
    volumes:
      - postgres_data:/var/lib/postgresql/data{{end}}

  {{if .UseRedis}}redis:
    image: redis:latest
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data{{end}}

  {{if .UseRabbitMQ}}rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      - RABBITMQ_DEFAULT_USER=guest
      - RABBITMQ_DEFAULT_PASS=guest
    volumes:
      - rabbitmq_data:/var/lib/rabbitmq{{end}}

volumes:
  {{if .UsePostgres}}postgres_data:{{end}}
  {{if .UseRedis}}redis_data:{{end}}
  {{if .UseRabbitMQ}}rabbitmq_data:{{end}} 
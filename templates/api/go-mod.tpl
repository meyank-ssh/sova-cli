module {{.ModuleName}}

go {{.GoVersion}}

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
	{{if .UseZap}}go.uber.org/zap v1.27.0{{end}}
	{{if .UsePostgres}}github.com/lib/pq v1.10.9{{end}}
	{{if .UseRedis}}github.com/redis/go-redis/v9 v9.5.1{{end}}
	{{if .UseRabbitMQ}}github.com/rabbitmq/amqp091-go v1.9.0{{end}}
) 
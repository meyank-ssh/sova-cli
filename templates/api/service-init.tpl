package service

func InitServices() error {
	{{if .UsePostgres}}// Initialize PostgreSQL
	if err := InitPostgres(); err != nil {
		return err
	}
	{{end}}

	{{if .UseRedis}}// Initialize Redis
	if err := InitRedis(); err != nil {
		return err
	}
	{{end}}

	{{if .UseRabbitMQ}}// Initialize RabbitMQ
	if err := InitRabbitMQ(); err != nil {
		return err
	}
	{{end}}

	return nil
}

func CloseServices() {
	{{if .UsePostgres}}ClosePostgres(){{end}}
	{{if .UseRedis}}CloseRedis(){{end}}
	{{if .UseRabbitMQ}}CloseRabbitMQ(){{end}}
} 
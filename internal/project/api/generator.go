package api

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-sova/sova-cli/pkg/questions"
)

type APIProjectGenerator struct {
	ProjectName string
	ProjectDir  string
	Answers     *questions.ProjectAnswers
}

func NewAPIProjectGenerator(projectName, projectDir string, answers *questions.ProjectAnswers) *APIProjectGenerator {
	return &APIProjectGenerator{
		ProjectName: projectName,
		ProjectDir:  projectDir,
		Answers:     answers,
	}
}

func (g *APIProjectGenerator) Generate() (map[string]string, []string, error) {
	dirs := []string{
		"internal/handlers",
		"internal/middleware",
		"internal/models",
		"internal/server",
		"internal/service",
		"routes",
	}

	files := make(map[string]string)

	files["routes/routes.go"] = g.generateRoutesFile()

	if g.Answers.UsePostgres {
		files["internal/service/postgres.go"] = g.generatePostgresFile()
	}

	if g.Answers.UseRedis {
		files["internal/service/redis.go"] = g.generateRedisFile()
	}

	if g.Answers.UseRabbitMQ {
		files["internal/service/rabbitmq.go"] = g.generateRabbitmqFile()
	}

	files["internal/service/init.go"] = g.generateServiceInitFile()

	files["internal/handlers/ping.go"] = g.generatePingHandlerFile()

	if g.Answers.UseZap {
		files["internal/middleware/logging.go"] = g.generateLoggingMiddlewareFile()
	}

	files["internal/server/server.go"] = g.generateServerFile()

	files["cmd/main.go"] = g.generateMainFile()

	files[".env"] = g.generateEnvFile()

	files["docker-compose.yml"] = g.generateDockerComposeFile()

	files["go.mod"] = g.generateGoModFile()

	return files, dirs, nil
}

func (g *APIProjectGenerator) generateRoutesFile() string {
	imports := fmt.Sprintf(`"github.com/gin-gonic/gin"
	"%s/internal/handlers"`, g.ProjectName)

	if g.Answers.UseZap {
		imports += fmt.Sprintf(`
	"%s/internal/middleware"`, g.ProjectName)
	}

	middlewareSetup := ""
	if g.Answers.UseZap {
		middlewareSetup = "// Add logging middleware\n\trouter.Use(middleware.LoggingMiddleware())\n"
	}

	return fmt.Sprintf(`package routes

import (
	%s
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
	%s
	// API routes
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
	}
}
`, imports, middlewareSetup)
}

func (g *APIProjectGenerator) generatePostgresFile() string {
	return `package service

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() error {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	
	if err := DB.Ping(); err != nil {
		return err
	}
	
	return nil
}

func ClosePostgres() {
	if DB != nil {
		DB.Close()
	}
}`
}

func (g *APIProjectGenerator) generateRedisFile() string {
	return `package service

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	
	return nil
}

func CloseRedis() {
	if RedisClient != nil {
		RedisClient.Close()
	}
}`
}

func (g *APIProjectGenerator) generateRabbitmqFile() string {
	return `package service

import (
	"os"
	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQ *amqp.Connection

func InitRabbitMQ() error {
	var err error
	RabbitMQ, err = amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		return err
	}
	
	return nil
}

func CloseRabbitMQ() {
	if RabbitMQ != nil {
		RabbitMQ.Close()
	}
}`
}

func (g *APIProjectGenerator) generateServiceInitFile() string {
	var inits []string
	var closers []string

	if g.Answers.UsePostgres {
		inits = append(inits, `
	// Initialize PostgreSQL
	if err := InitPostgres(); err != nil {
		return err
	}`)
		closers = append(closers, "\tClosePostgres()")
	}

	if g.Answers.UseRedis {
		inits = append(inits, `
	// Initialize Redis
	if err := InitRedis(); err != nil {
		return err
	}`)
		closers = append(closers, "\tCloseRedis()")
	}

	if g.Answers.UseRabbitMQ {
		inits = append(inits, `
	// Initialize RabbitMQ
	if err := InitRabbitMQ(); err != nil {
		return err
	}`)
		closers = append(closers, "\tCloseRabbitMQ()")
	}

	initBody := strings.Join(inits, "\n")
	if initBody != "" {
		initBody += "\n"
	}
	initBody += "\treturn nil"

	closerBody := ""
	if len(closers) > 0 {
		closerBody = "\n\t" + strings.Join(closers, "\n\t")
	}

	return fmt.Sprintf(`package service

func InitServices() error {%s
}

func CloseServices() {%s
}`, initBody, closerBody)
}

func (g *APIProjectGenerator) generatePingHandlerFile() string {
	return `package handlers

import (
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
`
}

func (g *APIProjectGenerator) generateLoggingMiddlewareFile() string {
	return `package middleware

import (
	"time"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("request completed",
			zap.String("path", path),
			zap.String("method", method),
			zap.Int("status", status),
			zap.Duration("latency", latency),
		)
	}
}
`
}

func (g *APIProjectGenerator) generateServerFile() string {
	return fmt.Sprintf(`package server

import (
	"fmt"
	"os"
	"%s/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	// Setup routes
	routes.SetupRoutes(s.router)

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	return s.router.Run(fmt.Sprintf(":%s", port))
}
`, g.ProjectName)
}

func (g *APIProjectGenerator) generateMainFile() string {
	return fmt.Sprintf(`package main

import (
	"log"
	"%s/internal/server"
	"%s/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize all services
	if err := service.InitServices(); err != nil {
		log.Fatalf("Failed to initialize services: %%v", err)
	}
	defer service.CloseServices()

	// Create and start server
	srv := server.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %%v", err)
	}
}
`, g.ProjectName, g.ProjectName)
}

func (g *APIProjectGenerator) generateEnvFile() string {
	var envVars []string
	envVars = append(envVars, "# Server Configuration\nPORT=8080")

	if g.Answers.UsePostgres {
		envVars = append(envVars, `# Database Configuration
DATABASE_URL=postgres://postgres:postgres@localhost:5432/`+g.ProjectName+`?sslmode=disable`)
	}

	if g.Answers.UseRedis {
		envVars = append(envVars, `# Redis Configuration
REDIS_URL=localhost:6379`)
	}

	if g.Answers.UseRabbitMQ {
		envVars = append(envVars, `# RabbitMQ Configuration
RABBITMQ_URL=amqp://guest:guest@localhost:5672/`)
	}

	return strings.Join(envVars, "\n\n")
}

func (g *APIProjectGenerator) generateDockerComposeFile() string {
	var services []string

	if g.Answers.UsePostgres {
		services = append(services, `
  postgres:
    image: postgres:latest
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=`+g.ProjectName)
	}

	if g.Answers.UseRedis {
		services = append(services, `
  redis:
    image: redis:latest
    ports:
      - "6379:6379"`)
	}

	if g.Answers.UseRabbitMQ {
		services = append(services, `
  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"`)
	}

	serviceStr := strings.Join(services, "\n")
	return fmt.Sprintf(`version: '3'
services:%s`, serviceStr)
}

func (g *APIProjectGenerator) generateGoModFile() string {
	var deps []string
	deps = append(deps, "github.com/gin-gonic/gin v1.9.1", "github.com/joho/godotenv v1.5.1")

	if g.Answers.UseZap {
		deps = append(deps, "go.uber.org/zap v1.27.0")
	}
	if g.Answers.UsePostgres {
		deps = append(deps, "github.com/lib/pq v1.10.9")
	}
	if g.Answers.UseRedis {
		deps = append(deps, "github.com/redis/go-redis/v9 v9.5.1")
	}
	if g.Answers.UseRabbitMQ {
		deps = append(deps, "github.com/rabbitmq/amqp091-go v1.9.0")
	}

	formattedDeps := strings.Join(deps, "\n\t")

	return fmt.Sprintf(`module %s

go 1.21

require (
	%s
)
`, g.ProjectName, formattedDeps)
}

func (g *APIProjectGenerator) WriteFiles(files map[string]string) error {
	for filename, content := range files {
		filePath := filepath.Join(g.ProjectDir, filename)

		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filename, err)
		}

		fmt.Printf("Created file: %s\n", filePath)
	}

	return nil
}

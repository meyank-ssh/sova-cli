package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	initTemplate string
	initForce    bool
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Long: `Initialize a new project with the specified name.
This will create a new directory with the project name and
set up the basic structure and files needed for your project.

Example:
  sova init my-awesome-project
  sova init my-awesome-project --template go-web
  sova init my-awesome-project --force`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := ""
		if len(args) > 0 {
			projectName = args[0]
		}

		// Ask project questions
		answers, err := questions.AskProjectQuestions(projectName)
		if err != nil {
			PrintError("Failed to get project details: %v", err)
			os.Exit(1)
		}

		projectName = answers.ProjectName
		cwd, err := os.Getwd()
		if err != nil {
			PrintError("Failed to get current directory: %v", err)
			os.Exit(1)
		}

		projectDir := filepath.Join(cwd, projectName)
		fmt.Printf("Project directory will be: %s\n", projectDir)

		if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
			if !initForce {
				PrintError("Directory %s already exists. Use --force to overwrite.", projectName)
				os.Exit(1)
			}
			PrintWarning("Overwriting existing directory: %s", projectName)
			if err := os.RemoveAll(projectDir); err != nil {
				PrintError("Failed to remove existing directory: %v", err)
				os.Exit(1)
			}
		}

		PrintInfo("Initializing new project: %s", projectName)

		if err := os.MkdirAll(projectDir, 0755); err != nil {
			PrintError("Failed to create project directory: %v", err)
			os.Exit(1)
		}

		// Base directories for all project types
		dirs := []string{
			"cmd",
			"internal",
			"test",
			"routes",
		}

		// Add API specific directories and files if needed
		apiFiles := make(map[string]string)
		if answers.ProjectType == "API" {
			dirs = append(dirs,
				"internal/handlers",
				"internal/middleware",
				"internal/models",
				"internal/server",
				"internal/service",
			)

			// Add routes setup in root level routes directory
			apiFiles["routes/routes.go"] = `package routes

import (
	"` + projectName + `/internal/handlers"
	` + func() string {
				if answers.UseZap {
					return `"` + projectName + `/internal/middleware"`
				}
				return ""
			}() + `
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
	` + func() string {
				if answers.UseZap {
					return "// Add logging middleware\n\trouter.Use(middleware.LoggingMiddleware())\n"
				}
				return ""
			}() + `
	// API routes
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
	}
}
`

			// Add service files for connections
			if answers.UsePostgres {
				apiFiles["internal/service/postgres.go"] = `package service

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

			if answers.UseRedis {
				apiFiles["internal/service/redis.go"] = `package service

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

			if answers.UseRabbitMQ {
				apiFiles["internal/service/rabbitmq.go"] = `package service

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

			// Add service initializer
			apiFiles["internal/service/init.go"] = `package service

func InitServices() error {` + func() string {
				var inits []string
				if answers.UsePostgres {
					inits = append(inits, `
	// Initialize PostgreSQL
	if err := InitPostgres(); err != nil {
		return err
	}`)
				}
				if answers.UseRedis {
					inits = append(inits, `
	// Initialize Redis
	if err := InitRedis(); err != nil {
		return err
	}`)
				}
				if answers.UseRabbitMQ {
					inits = append(inits, `
	// Initialize RabbitMQ
	if err := InitRabbitMQ(); err != nil {
		return err
	}`)
				}
				return strings.Join(inits, "\n") + `
	
	return nil`
			}() + `
}

func CloseServices() {` + func() string {
				var closers []string
				if answers.UsePostgres {
					closers = append(closers, "\tClosePostgres()")
				}
				if answers.UseRedis {
					closers = append(closers, "\tCloseRedis()")
				}
				if answers.UseRabbitMQ {
					closers = append(closers, "\tCloseRabbitMQ()")
				}
				return "\n\t" + strings.Join(closers, "\n\t")
			}() + `
}`

			// Add handlers
			apiFiles["internal/handlers/ping.go"] = `package handlers

import (
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
`

			if answers.UseZap {
				apiFiles["internal/middleware/logging.go"] = `package middleware

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

			// Update server setup with root level routes import
			apiFiles["internal/server/server.go"] = `package server

import (
	"fmt"
	"os"
	"` + projectName + `/routes"
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
`

			// Update main.go with cleaner initialization
			apiFiles["cmd/main.go"] = `package main

import (
	"log"
	"` + projectName + `/internal/server"
	"` + projectName + `/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize all services
	if err := service.InitServices(); err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer service.CloseServices()

	// Create and start server
	srv := server.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
`

			// Add .env file with conditional variables
			apiFiles[".env"] = `# Server Configuration
PORT=8080

` + func() string {
				var envVars []string
				if answers.UsePostgres {
					envVars = append(envVars, `# Database Configuration
DATABASE_URL=postgres://postgres:postgres@localhost:5432/`+projectName+`?sslmode=disable`)
				}
				if answers.UseRedis {
					envVars = append(envVars, `# Redis Configuration
REDIS_URL=localhost:6379`)
				}
				if answers.UseRabbitMQ {
					envVars = append(envVars, `# RabbitMQ Configuration
RABBITMQ_URL=amqp://guest:guest@localhost:5672/`)
				}
				return strings.Join(envVars, "\n\n")
			}()

			// Add docker-compose.yml with conditional services
			apiFiles["docker-compose.yml"] = `version: '3'
services:` + func() string {
				var services []string
				if answers.UsePostgres {
					services = append(services, `
  postgres:
    image: postgres:latest
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=`+projectName)
				}
				if answers.UseRedis {
					services = append(services, `
  redis:
    image: redis:latest
    ports:
      - "6379:6379"`)
				}
				if answers.UseRabbitMQ {
					services = append(services, `
  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"`)
				}
				return strings.Join(services, "\n")
			}()

			// Update go.mod for API with conditional dependencies
			baseAPIFiles := map[string]string{
				"go.mod": fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
	`+func() string {
					var deps []string
					if answers.UseZap {
						deps = append(deps, "go.uber.org/zap v1.27.0")
					}
					if answers.UsePostgres {
						deps = append(deps, "github.com/lib/pq v1.10.9")
					}
					if answers.UseRedis {
						deps = append(deps, "github.com/redis/go-redis/v9 v9.5.1")
					}
					if answers.UseRabbitMQ {
						deps = append(deps, "github.com/rabbitmq/amqp091-go v1.9.0")
					}
					return strings.Join(deps, "\n\t")
				}()+`
)
`, projectName),
			}

			// Merge base API files with conditional files
			for k, v := range baseAPIFiles {
				apiFiles[k] = v
			}
		}

		// Create directories
		for _, dir := range dirs {
			dirPath := filepath.Join(projectDir, dir)
			err := os.MkdirAll(dirPath, 0755)
			if err != nil {
				PrintError("Failed to create directory %s: %v", dir, err)
				os.Exit(1)
			}
			fmt.Printf("Created directory: %s\n", dirPath)
		}

		// Create all files
		for filename, content := range apiFiles {
			filePath := filepath.Join(projectDir, filename)
			// Ensure the directory exists
			dir := filepath.Dir(filePath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				PrintError("Failed to create directory %s: %v", dir, err)
				os.Exit(1)
			}
			// Write the file
			if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
				PrintError("Failed to create file %s: %v", filename, err)
				os.Exit(1)
			}
			fmt.Printf("Created file: %s\n", filePath)
		}

		PrintSuccess("Project %s initialized successfully!", projectName)
		if answers.ProjectType == "API" {
			PrintInfo("\nTo start your API server:")
			PrintInfo("cd %s", projectName)
			PrintInfo("go mod tidy")
			PrintInfo("go run cmd/main.go")
			PrintInfo("\nThen visit http://localhost:8080/api/ping")

			if answers.UsePostgres {
				PrintInfo("\nMake sure PostgreSQL is running:")
				PrintInfo("docker-compose up postgres -d")
			}
			if answers.UseRedis {
				PrintInfo("\nMake sure Redis is running:")
				PrintInfo("docker-compose up redis -d")
			}
			if answers.UseRabbitMQ {
				PrintInfo("\nMake sure RabbitMQ is running:")
				PrintInfo("docker-compose up rabbitmq -d")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&initTemplate, "template", "t", "default", "Template to use for project initialization")
	initCmd.Flags().BoolVarP(&initForce, "force", "f", false, "Force initialization even if directory exists")

	viper.BindPFlag("init.template", initCmd.Flags().Lookup("template"))
	viper.BindPFlag("init.force", initCmd.Flags().Lookup("force"))
}

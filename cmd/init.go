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
			"pkg",
			"docs",
			"scripts",
			"test",
		}

		// Add API specific directories and files if needed
		apiFiles := make(map[string]string)
		if answers.ProjectType == "API" {
			dirs = append(dirs, "api")

			// Add feature-specific directories and files
			if answers.UseZap {
				dirs = append(dirs, "internal/logger")
				apiFiles["internal/logger/logger.go"] = `package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

// Initialize sets up the logger
func Initialize() error {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		return err
	}
	return nil
}

// Get returns the global logger instance
func Get() *zap.Logger {
	return log
}
`
			}

			if answers.UseRabbitMQ {
				dirs = append(dirs, "internal/broker")
				apiFiles["internal/broker/rabbitmq.go"] = `package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ holds the connection and channel
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ connection
func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

// Close closes the connection and channel
func (r *RabbitMQ) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}
`
			}

			if answers.UseRedis {
				dirs = append(dirs, "internal/cache")
				apiFiles["internal/cache/redis.go"] = `package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Redis holds the Redis client
type Redis struct {
	client *redis.Client
}

// NewRedis creates a new Redis connection
func NewRedis(addr string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &Redis{
		client: client,
	}
}

// Get retrieves a value from Redis
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Set stores a value in Redis
func (r *Redis) Set(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}
`
			}

			if answers.UsePostgres {
				dirs = append(dirs, "internal/database")
				apiFiles["internal/database/postgres.go"] = `package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Postgres holds the database connection
type Postgres struct {
	db *sql.DB
}

// NewPostgres creates a new Postgres connection
func NewPostgres(connStr string) (*Postgres, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

// Close closes the database connection
func (p *Postgres) Close() error {
	return p.db.Close()
}
`
			}

			// Add handlers and middleware directories
			dirs = append(dirs, "internal/handlers", "internal/middleware")
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

		// Base files for all projects
		files := map[string]string{
			"go.mod": fmt.Sprintf(`module %s

go 1.21

require (
`, projectName),
			"README.md": fmt.Sprintf(`# %s

This project was generated using Sova CLI.

## Project Type
%s

## Features
`, projectName, answers.ProjectType),
		}

		// Add main.go based on project type
		if answers.ProjectType == "CLI" {
			files["main.go"] = `package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "` + projectName + `",
		Short: "A brief description of your CLI application",
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`
			files["go.mod"] += `	github.com/spf13/cobra v1.9.1
)
`
			files["README.md"] += `- Command Line Interface
- Built with Cobra
`
		} else {
			// API project
			mainImports := []string{
				`	"log"`,
				`	"net/http"`,
				`	"github.com/gin-gonic/gin"`,
			}

			mainInit := []string{}
			if answers.UseZap {
				mainImports = append(mainImports, `	"`+projectName+`/internal/logger"`)
				mainInit = append(mainInit, `	if err := logger.Initialize(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}`)
			}

			if answers.UseRabbitMQ {
				mainImports = append(mainImports, `	"`+projectName+`/internal/broker"`)
				mainInit = append(mainInit, `	rmq, err := broker.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rmq.Close()`)
			}

			if answers.UseRedis {
				mainImports = append(mainImports, `	"`+projectName+`/internal/cache"`)
				mainInit = append(mainInit, `	redis := cache.NewRedis("localhost:6379")`)
			}

			if answers.UsePostgres {
				mainImports = append(mainImports, `	"`+projectName+`/internal/database"`)
				mainInit = append(mainInit, `	db, err := database.NewPostgres("postgres://postgres:postgres@localhost:5432/`+projectName+`?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()`)
			}

			files["main.go"] = `package main

import (
` + strings.Join(mainImports, "\n") + `
)

func main() {
` + strings.Join(mainInit, "\n") + `

	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
`
			// Add dependencies based on selected features
			deps := []string{
				"github.com/gin-gonic/gin v1.9.1",
			}

			if answers.UseZap {
				deps = append(deps, "go.uber.org/zap v1.27.0")
				files["README.md"] += "- Zap logging\n"
			}

			if answers.UseRabbitMQ {
				deps = append(deps, "github.com/rabbitmq/amqp091-go v1.9.0")
				files["README.md"] += "- RabbitMQ integration\n"
			}

			if answers.UseRedis {
				deps = append(deps, "github.com/redis/go-redis/v9 v9.5.1")
				files["README.md"] += "- Redis caching\n"
			}

			if answers.UsePostgres {
				deps = append(deps, "github.com/lib/pq v1.10.9")
				files["README.md"] += "- PostgreSQL database\n"
			}

			files["go.mod"] += "	" + strings.Join(deps, "\n\t") + "\n)"
		}

		// Merge API-specific files with base files
		for filename, content := range apiFiles {
			files[filename] = content
		}

		// Create all files
		for filename, content := range files {
			filePath := filepath.Join(projectDir, filename)
			err := ioutil.WriteFile(filePath, []byte(content), 0644)
			if err != nil {
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
			PrintInfo("go run main.go")
			PrintInfo("\nThen visit http://localhost:8080/ping")

			if answers.UsePostgres {
				PrintInfo("\nMake sure PostgreSQL is running and create a database named '%s'", projectName)
			}
			if answers.UseRedis {
				PrintInfo("\nMake sure Redis server is running on localhost:6379")
			}
			if answers.UseRabbitMQ {
				PrintInfo("\nMake sure RabbitMQ server is running on localhost:5672")
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

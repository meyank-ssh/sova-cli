package api

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-sova/sova-cli/internal/templates"
	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/go-sova/sova-cli/pkg/utils"
)

type APIProjectGenerator struct {
	ProjectName    string
	ProjectDir     string
	Answers        *questions.ProjectAnswers
	templateLoader *templates.TemplateLoader
	fileGenerator  *templates.FileGenerator
	logger         *utils.Logger
}

func NewAPIProjectGenerator(projectName, projectDir string, answers *questions.ProjectAnswers) *APIProjectGenerator {
	// Try multiple locations for the templates
	templateDirs := []string{
		"templates",                                  // Current directory
		filepath.Join("..", "templates"),             // One level up
		filepath.Join("..", "..", "templates"),       // Two levels up
		filepath.Join("..", "..", "..", "templates"), // Three levels up
	}

	// If running as an executable, also try relative to the executable
	execPath, err := os.Executable()
	if err == nil {
		execDir := filepath.Dir(execPath)
		templateDirs = append(templateDirs,
			filepath.Join(execDir, "templates"),
			filepath.Join(execDir, "..", "templates"),
			filepath.Join(execDir, "..", "..", "templates"),
			filepath.Join(execDir, "..", "..", "..", "templates"),
		)
	}

	// Try to find an existing template directory
	var templateDir string
	for _, dir := range templateDirs {
		if _, err := os.Stat(dir); err == nil {
			// Verify that it contains the expected subdirectories
			if _, err := os.Stat(filepath.Join(dir, "web")); err == nil {
				templateDir = dir
				break
			}
		}
	}

	// If no template directory found, use the default one
	if templateDir == "" {
		templateDir = "templates"
	}

	loader := templates.NewTemplateLoader(templateDir)
	return &APIProjectGenerator{
		ProjectName:    projectName,
		ProjectDir:     projectDir,
		Answers:        answers,
		templateLoader: loader,
		fileGenerator:  templates.NewFileGenerator(loader),
		logger:         utils.NewLoggerWithPrefix(utils.Info, "APIProjectGenerator"),
	}
}

func (g *APIProjectGenerator) SetLogger(logger *utils.Logger) {
	g.logger = logger
	g.templateLoader.SetLogger(logger)
	g.fileGenerator.SetLogger(logger)
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

	// Map of file paths to template names
	fileTemplates := map[string]string{
		"routes/routes.go":          "api/routes.tpl",
		"internal/handlers/ping.go": "api/handlers.tpl",
		"internal/server/server.go": "api/server.tpl",
		"cmd/main.go":               "api/main.tpl",
		".env":                      "api/env.tpl",
		"docker-compose.yml":        "api/docker-compose.tpl",
		"go.mod":                    "api/go-mod.tpl",
	}

	// Add conditional files based on answers
	if g.Answers.UsePostgres {
		fileTemplates["internal/service/postgres.go"] = "api/postgres.tpl"
	}

	if g.Answers.UseRedis {
		fileTemplates["internal/service/redis.go"] = "api/redis.tpl"
	}

	if g.Answers.UseRabbitMQ {
		fileTemplates["internal/service/rabbitmq.go"] = "api/rabbitmq.tpl"
	}

	if g.Answers.UseZap {
		fileTemplates["internal/middleware/logging.go"] = "api/logging.tpl"
	}

	// Add service initialization file
	fileTemplates["internal/service/init.go"] = "api/service-init.tpl"

	return fileTemplates, dirs, nil
}

func (g *APIProjectGenerator) WriteFiles(files map[string]string) error {
	for filePath, templateName := range files {
		fullPath := filepath.Join(g.ProjectDir, filePath)

		// Prepare template data
		data := map[string]interface{}{
			"ProjectName":        g.ProjectName,
			"ProjectDescription": "A Go API project created with Sova CLI",
			"ModuleName":         g.ProjectName,
			"GoVersion":          "1.21",
			"Author":             "Meyank Singh",
			"License":            "MIT",
			"Year":               fmt.Sprintf("%d", time.Now().Year()),
			"UseZap":             g.Answers.UseZap,
			"UsePostgres":        g.Answers.UsePostgres,
			"UseRedis":           g.Answers.UseRedis,
			"UseRabbitMQ":        g.Answers.UseRabbitMQ,
		}

		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}

		if err := g.fileGenerator.GenerateFile(templateName, fullPath, data); err != nil {
			return fmt.Errorf("failed to generate file %s from template %s: %v", filePath, templateName, err)
		}

		fmt.Printf("Created file: %s\n", fullPath)
	}

	return nil
}

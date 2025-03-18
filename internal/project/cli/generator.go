package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-sova/sova-cli/internal/templates"
	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/go-sova/sova-cli/pkg/utils"
)

type CLIProjectGenerator struct {
	ProjectName    string
	ProjectDir     string
	Answers        *questions.ProjectAnswers
	templateLoader *templates.TemplateLoader
	fileGenerator  *templates.FileGenerator
	logger         *utils.Logger
}

func NewCLIProjectGenerator(projectName, projectDir string, answers *questions.ProjectAnswers) *CLIProjectGenerator {
	execPath, err := os.Executable()
	templateDir := "templates"
	if err == nil {
		execDir := filepath.Dir(execPath)
		templateDir = filepath.Join(execDir, "templates")
	}

	loader := templates.NewTemplateLoader(templateDir)
	return &CLIProjectGenerator{
		ProjectName:    projectName,
		ProjectDir:     projectDir,
		Answers:        answers,
		templateLoader: loader,
		fileGenerator:  templates.NewFileGenerator(loader),
		logger:         utils.NewLoggerWithPrefix(utils.Info, "CLIProjectGenerator"),
	}
}

func (g *CLIProjectGenerator) SetLogger(logger *utils.Logger) {
	g.logger = logger
	g.templateLoader.SetLogger(logger)
	g.fileGenerator.SetLogger(logger)
}

func (g *CLIProjectGenerator) Generate() (map[string]string, []string, error) {
	dirs := []string{
		"cmd",
		"internal",
		"test",
	}

	// Map of file paths to template names
	fileTemplates := map[string]string{
		"cmd/root.go":        "cli/root.tpl",
		"cmd/command1.go":    "cli/command.tpl",
		"cmd/command2.go":    "cli/command.tpl",
		"internal/utils.go":  "cli/utils.tpl",
		"internal/config.go": "cli/config.tpl",
		"main.go":            "cli/main.tpl",
		"README.md":          "cli/readme.tpl",
		"go.mod":             "cli/go-mod.tpl",
	}

	// Generate files using templates
	files := make(map[string]string)
	for filePath, templateName := range fileTemplates {
		files[filePath] = templateName
	}

	return files, dirs, nil
}

func (g *CLIProjectGenerator) WriteFiles(files map[string]string) error {
	for filePath, templateName := range files {
		fullPath := filepath.Join(g.ProjectDir, filePath)

		// Prepare template data
		data := map[string]interface{}{
			"ProjectName":        g.ProjectName,
			"ProjectDescription": "A CLI application built with Go and Cobra",
			"ModuleName":         g.ProjectName,
			"GoVersion":          "1.21",
			"Author":             "Meyank Singh",
			"License":            "MIT",
			"Year":               fmt.Sprintf("%d", time.Now().Year()),
		}

		// Special handling for command files
		if filepath.Base(filePath) == "command1.go" {
			data["CommandName"] = "command1"
		} else if filepath.Base(filePath) == "command2.go" {
			data["CommandName"] = "command2"
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

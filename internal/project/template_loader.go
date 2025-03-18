package project

import (
	"fmt"
	"path/filepath"

	"github.com/go-sova/sova-cli/internal/templates"
	"github.com/go-sova/sova-cli/pkg/utils"
)

type TemplateManager struct {
	templateLoader *templates.TemplateLoader
	logger         *utils.Logger
}

func NewTemplateManager(templateDir string) *TemplateManager {
	return &TemplateManager{
		templateLoader: templates.NewTemplateLoader(templateDir),
		logger:         utils.NewLoggerWithPrefix(utils.Info, "TemplateManager"),
	}
}

func (m *TemplateManager) SetLogger(logger *utils.Logger) {
	m.logger = logger
	m.templateLoader.SetLogger(logger)
}

func (m *TemplateManager) ListTemplates() ([]string, error) {
	m.logger.Debug("Listing templates")

	categories, err := m.templateLoader.ListTemplateCategories()
	if err != nil {
		return nil, err
	}

	builtInTemplates := []string{"default", "go-web", "cli", "library"}

	allTemplates := append(builtInTemplates, categories...)

	return allTemplates, nil
}

func (m *TemplateManager) GetTemplateDescription(templateName string) (string, error) {
	m.logger.Debug("Getting description for template: %s", templateName)

	switch templateName {
	case "default":
		return "A basic Go project with a minimal structure", nil
	case "go-web":
		return "A Go web application with a complete structure for web development", nil
	case "cli":
		return "A command-line interface application with Cobra", nil
	case "library":
		return "A Go library with examples and documentation", nil
	}

	categoryDir := filepath.Join(m.templateLoader.GetTemplateDir(), templateName)
	if utils.DirExists(categoryDir) {
		descFile := filepath.Join(categoryDir, "description.txt")
		if utils.FileExists(descFile) {
			content, err := utils.ReadFile(descFile)
			if err == nil {
				return string(content), nil
			}
		}

		return fmt.Sprintf("Custom template: %s", templateName), nil
	}

	return "", fmt.Errorf("unknown template: %s", templateName)
}

func (m *TemplateManager) ValidateTemplate(templateName string) error {
	m.logger.Debug("Validating template: %s", templateName)

	switch templateName {
	case "default", "go-web", "cli", "library":
		return nil
	}

	categoryDir := filepath.Join(m.templateLoader.GetTemplateDir(), templateName)
	if utils.DirExists(categoryDir) {
		return nil
	}

	return fmt.Errorf("unknown template: %s", templateName)
}

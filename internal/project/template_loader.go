package project

import (
	"fmt"

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
	return []string{"api", "cli"}, nil
}

func (m *TemplateManager) GetTemplateDescription(templateName string) (string, error) {
	m.logger.Debug("Getting description for template: %s", templateName)

	switch templateName {
	case "api":
		return "A Go API project with a complete structure for API development", nil
	case "cli":
		return "A command-line interface application with Cobra", nil
	}

	return "", fmt.Errorf("unknown template: %s", templateName)
}

func (m *TemplateManager) ValidateTemplate(templateName string) error {
	m.logger.Debug("Validating template: %s", templateName)

	switch templateName {
	case "api", "cli":
		return nil
	}

	return fmt.Errorf("unknown template: %s", templateName)
}

package project

import (
	"fmt"

	"github.com/go-sova/sova-cli/pkg/utils"
	"github.com/go-sova/sova-cli/templates"
)

type TemplateManager struct {
	logger         *utils.Logger
	templateLoader *templates.TemplateLoader
}

func NewTemplateManager() *TemplateManager {
	loader := templates.NewTemplateLoader()
	return &TemplateManager{
		logger:         utils.NewLoggerWithPrefix(utils.Info, "TemplateManager"),
		templateLoader: loader,
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

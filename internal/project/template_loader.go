package project

import (
	"fmt"
	"path/filepath"

	"github.com/meyanksingh/go-sova/internal/templates"
	"github.com/meyanksingh/go-sova/internal/utils"
)

// TemplateManager manages project templates
type TemplateManager struct {
	templateLoader *templates.TemplateLoader
	logger         *utils.Logger
}

// NewTemplateManager creates a new template manager
func NewTemplateManager(templateDir string) *TemplateManager {
	return &TemplateManager{
		templateLoader: templates.NewTemplateLoader(templateDir),
		logger:         utils.NewLoggerWithPrefix(utils.Info, "TemplateManager"),
	}
}

// SetLogger sets the logger for the template manager
func (m *TemplateManager) SetLogger(logger *utils.Logger) {
	m.logger = logger
	m.templateLoader.SetLogger(logger)
}

// ListTemplates lists all available templates
func (m *TemplateManager) ListTemplates() ([]string, error) {
	m.logger.Debug("Listing templates")

	// First, check for template categories
	categories, err := m.templateLoader.ListTemplateCategories()
	if err != nil {
		return nil, err
	}

	// Add built-in templates
	builtInTemplates := []string{"default", "go-web", "cli", "library"}

	// Combine built-in templates with categories
	allTemplates := append(builtInTemplates, categories...)

	return allTemplates, nil
}

// GetTemplateDescription returns the description of a template
func (m *TemplateManager) GetTemplateDescription(templateName string) (string, error) {
	m.logger.Debug("Getting description for template: %s", templateName)

	// Check built-in templates first
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

	// Check if it's a custom template category
	categoryDir := filepath.Join(m.templateLoader.GetTemplateDir(), templateName)
	if utils.DirExists(categoryDir) {
		// Try to read description from a description.txt file
		descFile := filepath.Join(categoryDir, "description.txt")
		if utils.FileExists(descFile) {
			content, err := utils.ReadFile(descFile)
			if err == nil {
				return string(content), nil
			}
		}

		// Return a generic description
		return fmt.Sprintf("Custom template: %s", templateName), nil
	}

	return "", fmt.Errorf("unknown template: %s", templateName)
}

// ValidateTemplate validates that a template exists
func (m *TemplateManager) ValidateTemplate(templateName string) error {
	m.logger.Debug("Validating template: %s", templateName)

	// Check built-in templates first
	switch templateName {
	case "default", "go-web", "cli", "library":
		return nil
	}

	// Check if it's a custom template category
	categoryDir := filepath.Join(m.templateLoader.GetTemplateDir(), templateName)
	if utils.DirExists(categoryDir) {
		return nil
	}

	return fmt.Errorf("unknown template: %s", templateName)
}

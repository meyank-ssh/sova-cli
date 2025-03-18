package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-sova/sova-cli/pkg/utils"
)

// TemplateLoader loads templates from the filesystem
type TemplateLoader struct {
	templateDir string
	logger      *utils.Logger
}

// NewTemplateLoader creates a new template loader
func NewTemplateLoader(templateDir string) *TemplateLoader {
	return &TemplateLoader{
		templateDir: templateDir,
		logger:      utils.NewLoggerWithPrefix(utils.Info, "TemplateLoader"),
	}
}

// SetLogger sets the logger for the template loader
func (l *TemplateLoader) SetLogger(logger *utils.Logger) {
	l.logger = logger
}

// GetTemplateDir returns the template directory
func (l *TemplateLoader) GetTemplateDir() string {
	return l.templateDir
}

// LoadTemplate loads a template from the filesystem
func (l *TemplateLoader) LoadTemplate(name string) (*template.Template, error) {
	templatePath := filepath.Join(l.templateDir, name)
	l.logger.Debug("Loading template: %s", templatePath)
	
	if !utils.FileExists(templatePath) {
		return nil, fmt.Errorf("template not found: %s", name)
	}
	
	return template.ParseFiles(templatePath)
}

// LoadTemplateWithFuncs loads a template from the filesystem with custom functions
func (l *TemplateLoader) LoadTemplateWithFuncs(name string, funcs template.FuncMap) (*template.Template, error) {
	templatePath := filepath.Join(l.templateDir, name)
	l.logger.Debug("Loading template with funcs: %s", templatePath)
	
	if !utils.FileExists(templatePath) {
		return nil, fmt.Errorf("template not found: %s", name)
	}
	
	return template.New(filepath.Base(name)).Funcs(funcs).ParseFiles(templatePath)
}

// ListTemplates lists all templates in the template directory
func (l *TemplateLoader) ListTemplates() ([]string, error) {
	var templates []string
	
	l.logger.Debug("Listing templates in: %s", l.templateDir)
	
	if !utils.DirExists(l.templateDir) {
		return nil, fmt.Errorf("template directory not found: %s", l.templateDir)
	}
	
	err := filepath.Walk(l.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && filepath.Ext(path) == ".tpl" {
			relPath, err := filepath.Rel(l.templateDir, path)
			if err != nil {
				return err
			}
			templates = append(templates, relPath)
		}
		
		return nil
	})
	
	return templates, err
}

// ListTemplateCategories lists all template categories (subdirectories)
func (l *TemplateLoader) ListTemplateCategories() ([]string, error) {
	var categories []string
	
	l.logger.Debug("Listing template categories in: %s", l.templateDir)
	
	if !utils.DirExists(l.templateDir) {
		return nil, fmt.Errorf("template directory not found: %s", l.templateDir)
	}
	
	entries, err := os.ReadDir(l.templateDir)
	if err != nil {
		return nil, err
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			categories = append(categories, entry.Name())
		}
	}
	
	return categories, nil
}

// GetTemplatesInCategory lists all templates in a specific category
func (l *TemplateLoader) GetTemplatesInCategory(category string) ([]string, error) {
	var templates []string
	
	categoryDir := filepath.Join(l.templateDir, category)
	l.logger.Debug("Listing templates in category: %s", categoryDir)
	
	if !utils.DirExists(categoryDir) {
		return nil, fmt.Errorf("template category not found: %s", category)
	}
	
	err := filepath.Walk(categoryDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && filepath.Ext(path) == ".tpl" {
			relPath, err := filepath.Rel(categoryDir, path)
			if err != nil {
				return err
			}
			templates = append(templates, relPath)
		}
		
		return nil
	})
	
	return templates, err
} 
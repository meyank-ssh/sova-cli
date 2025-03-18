package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-sova/sova-cli/pkg/utils"
)

type TemplateLoader struct {
	templateDir string
	logger      *utils.Logger
}

func NewTemplateLoader(templateDir string) *TemplateLoader {
	return &TemplateLoader{
		templateDir: templateDir,
		logger:      utils.NewLoggerWithPrefix(utils.Info, "TemplateLoader"),
	}
}

func (l *TemplateLoader) SetLogger(logger *utils.Logger) {
	l.logger = logger
}

func (l *TemplateLoader) GetTemplateDir() string {
	return l.templateDir
}

// LoadTemplate loads a template by name, searching first in category subdirectories
func (l *TemplateLoader) LoadTemplate(name string) (*template.Template, error) {
	// Try to find the template directly (backward compatibility)
	directPath := filepath.Join(l.templateDir, name)
	l.logger.Debug("Trying to load template directly: %s", directPath)
	
	if utils.FileExists(directPath) {
		l.logger.Debug("Found template at path: %s", directPath)
		return template.ParseFiles(directPath)
	}
	
	// If not found directly, try to find in subdirectories
	categories, err := l.ListTemplateCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to list template categories: %w", err)
	}
	
	for _, category := range categories {
		categoryPath := filepath.Join(l.templateDir, category, name)
		l.logger.Debug("Trying to load template from category %s: %s", category, categoryPath)
		
		if utils.FileExists(categoryPath) {
			l.logger.Debug("Found template in category %s: %s", category, categoryPath)
			return template.ParseFiles(categoryPath)
		}
	}
	
	return nil, fmt.Errorf("template not found: %s", name)
}

// LoadTemplateFromCategory loads a template from a specific category
func (l *TemplateLoader) LoadTemplateFromCategory(category, name string) (*template.Template, error) {
	templatePath := filepath.Join(l.templateDir, category, name)
	l.logger.Debug("Loading template from category %s: %s", category, templatePath)
	
	if !utils.FileExists(templatePath) {
		return nil, fmt.Errorf("template not found in category %s: %s", category, name)
	}
	
	return template.ParseFiles(templatePath)
}

func (l *TemplateLoader) LoadTemplateWithFuncs(name string, funcs template.FuncMap) (*template.Template, error) {
	// Try to find the template directly (backward compatibility)
	directPath := filepath.Join(l.templateDir, name)
	l.logger.Debug("Trying to load template with funcs directly: %s", directPath)
	
	if utils.FileExists(directPath) {
		l.logger.Debug("Found template at path: %s", directPath)
		return template.New(filepath.Base(name)).Funcs(funcs).ParseFiles(directPath)
	}
	
	// If not found directly, try to find in subdirectories
	categories, err := l.ListTemplateCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to list template categories: %w", err)
	}
	
	for _, category := range categories {
		categoryPath := filepath.Join(l.templateDir, category, name)
		l.logger.Debug("Trying to load template with funcs from category %s: %s", category, categoryPath)
		
		if utils.FileExists(categoryPath) {
			l.logger.Debug("Found template in category %s: %s", category, categoryPath)
			return template.New(filepath.Base(name)).Funcs(funcs).ParseFiles(categoryPath)
		}
	}
	
	return nil, fmt.Errorf("template not found: %s", name)
}

// LoadTemplateWithFuncsFromCategory loads a template with custom functions from a specific category
func (l *TemplateLoader) LoadTemplateWithFuncsFromCategory(category, name string, funcs template.FuncMap) (*template.Template, error) {
	templatePath := filepath.Join(l.templateDir, category, name)
	l.logger.Debug("Loading template with funcs from category %s: %s", category, templatePath)
	
	if !utils.FileExists(templatePath) {
		return nil, fmt.Errorf("template not found in category %s: %s", category, name)
	}
	
	return template.New(filepath.Base(name)).Funcs(funcs).ParseFiles(templatePath)
}

func (l *TemplateLoader) ListTemplates() ([]string, error) {
	var templates []string
	
	l.logger.Debug("Listing all templates in: %s", l.templateDir)
	
	if !utils.DirExists(l.templateDir) {
		return nil, fmt.Errorf("template directory not found: %s", l.templateDir)
	}
	
	// First, add templates in the root directory (backward compatibility)
	entries, err := os.ReadDir(l.templateDir)
	if err != nil {
		return nil, err
	}
	
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".tpl" {
			templates = append(templates, entry.Name())
		}
	}
	
	// Then add templates in category subdirectories
	categories, err := l.ListTemplateCategories()
	if err != nil {
		return nil, err
	}
	
	for _, category := range categories {
		categoryTemplates, err := l.GetTemplatesInCategory(category)
		if err != nil {
			l.logger.Warning("Failed to list templates in category %s: %v", category, err)
			continue
		}
		
		for _, tmpl := range categoryTemplates {
			// Add the category as a prefix to avoid name conflicts
			templates = append(templates, fmt.Sprintf("%s/%s", category, tmpl))
		}
	}
	
	return templates, nil
}

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

func (l *TemplateLoader) GetTemplatesInCategory(category string) ([]string, error) {
	var templates []string
	
	categoryDir := filepath.Join(l.templateDir, category)
	l.logger.Debug("Listing templates in category: %s", categoryDir)
	
	if !utils.DirExists(categoryDir) {
		return nil, fmt.Errorf("template category not found: %s", category)
	}
	
	entries, err := os.ReadDir(categoryDir)
	if err != nil {
		return nil, err
	}
	
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".tpl" {
			templates = append(templates, entry.Name())
		}
	}
	
	return templates, nil
} 
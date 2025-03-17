package templates

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/meyanksingh/go-sova/internal/utils"
)

// FileGenerator generates files from templates
type FileGenerator struct {
	loader *TemplateLoader
	logger *utils.Logger
}

// NewFileGenerator creates a new file generator
func NewFileGenerator(loader *TemplateLoader) *FileGenerator {
	return &FileGenerator{
		loader: loader,
		logger: utils.NewLoggerWithPrefix(utils.Info, "FileGenerator"),
	}
}

// SetLogger sets the logger for the file generator
func (g *FileGenerator) SetLogger(logger *utils.Logger) {
	g.logger = logger
}

// GenerateFile generates a file from a template
func (g *FileGenerator) GenerateFile(templateName, outputPath string, data interface{}) error {
	g.logger.Debug("Generating file from template: %s -> %s", templateName, outputPath)
	
	// Check if output file already exists
	if utils.FileExists(outputPath) {
		g.logger.Warning("Output file already exists: %s", outputPath)
	}
	
	// Load template
	tmpl, err := g.loader.LoadTemplate(templateName)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}
	
	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	
	// Create the directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := utils.CreateDirIfNotExists(dir); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	// Write file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	g.logger.Info("Generated file: %s", outputPath)
	return nil
}

// GenerateFileWithFuncs generates a file from a template with custom functions
func (g *FileGenerator) GenerateFileWithFuncs(templateName, outputPath string, data interface{}, funcs template.FuncMap) error {
	g.logger.Debug("Generating file from template with funcs: %s -> %s", templateName, outputPath)
	
	// Check if output file already exists
	if utils.FileExists(outputPath) {
		g.logger.Warning("Output file already exists: %s", outputPath)
	}
	
	// Load template
	tmpl, err := g.loader.LoadTemplateWithFuncs(templateName, funcs)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}
	
	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	
	// Create the directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := utils.CreateDirIfNotExists(dir); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	// Write file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	g.logger.Info("Generated file: %s", outputPath)
	return nil
}

// GenerateMultipleFiles generates multiple files from templates
func (g *FileGenerator) GenerateMultipleFiles(templates map[string]string, outputDir string, data interface{}) error {
	g.logger.Debug("Generating multiple files in: %s", outputDir)
	
	for templateName, outputFile := range templates {
		outputPath := filepath.Join(outputDir, outputFile)
		if err := g.GenerateFile(templateName, outputPath, data); err != nil {
			return fmt.Errorf("failed to generate file %s: %w", outputFile, err)
		}
	}
	
	return nil
}

// GenerateMultipleFilesWithFuncs generates multiple files from templates with custom functions
func (g *FileGenerator) GenerateMultipleFilesWithFuncs(templates map[string]string, outputDir string, data interface{}, funcs template.FuncMap) error {
	g.logger.Debug("Generating multiple files with funcs in: %s", outputDir)
	
	for templateName, outputFile := range templates {
		outputPath := filepath.Join(outputDir, outputFile)
		if err := g.GenerateFileWithFuncs(templateName, outputPath, data, funcs); err != nil {
			return fmt.Errorf("failed to generate file %s: %w", outputFile, err)
		}
	}
	
	return nil
} 
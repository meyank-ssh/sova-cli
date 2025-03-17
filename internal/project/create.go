package project

import (
	"fmt"

	"github.com/meyanksingh/go-sova/internal/templates"
	"github.com/meyanksingh/go-sova/internal/utils"
)

// ProjectCreator creates new projects
type ProjectCreator struct {
	logger         *utils.Logger
	templateLoader *templates.TemplateLoader
	fileGenerator  *templates.FileGenerator
}

// NewProjectCreator creates a new project creator
func NewProjectCreator(templateDir string) *ProjectCreator {
	loader := templates.NewTemplateLoader(templateDir)
	return &ProjectCreator{
		logger:         utils.NewLoggerWithPrefix(utils.Info, "ProjectCreator"),
		templateLoader: loader,
		fileGenerator:  templates.NewFileGenerator(loader),
	}
}

// SetLogger sets the logger for the project creator
func (c *ProjectCreator) SetLogger(logger *utils.Logger) {
	c.logger = logger
	c.templateLoader.SetLogger(logger)
	c.fileGenerator.SetLogger(logger)
}

// ProjectData contains data for project templates
type ProjectData struct {
	ProjectName        string
	ProjectDescription string
	ModuleName         string
	GoVersion          string
	Author             string
	License            string
	Year               string
}

// CreateProject creates a new project
func (c *ProjectCreator) CreateProject(projectName, projectDir, templateName string, force bool) error {
	c.logger.Info("Creating project: %s in directory: %s", projectName, projectDir)
	c.logger.Info("Using template: %s", templateName)

	// Check if the directory already exists
	if utils.DirExists(projectDir) {
		if !force {
			return fmt.Errorf("directory already exists: %s", projectDir)
		}
		c.logger.Warning("Overwriting existing directory: %s", projectDir)
	}

	// Get project structure
	structure, err := GetProjectStructure(templateName, projectName)
	if err != nil {
		return err
	}

	// Get project data
	projectData, err := c.getProjectData(projectName, structure.Description)
	if err != nil {
		return err
	}

	// Create project directories
	dirs, files := structure.GetAbsolutePaths(projectDir)
	for _, dir := range dirs {
		c.logger.Debug("Creating directory: %s", dir)
		if err := utils.CreateDirIfNotExists(dir); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// Generate project files
	for filePath, templateName := range files {
		c.logger.Debug("Generating file: %s from template: %s", filePath, templateName)
		if err := c.fileGenerator.GenerateFile(templateName, filePath, projectData); err != nil {
			return fmt.Errorf("failed to generate file: %w", err)
		}
	}

	c.logger.Info("Project created successfully!")
	return nil
}

// getProjectData gets data for project templates
func (c *ProjectCreator) getProjectData(projectName, description string) (*ProjectData, error) {
	// Get module name
	moduleName, err := utils.ReadInputWithDefault("Module name", "github.com/example/"+projectName)
	if err != nil {
		return nil, err
	}

	// Get project description
	projectDescription, err := utils.ReadInputWithDefault("Project description", description)
	if err != nil {
		return nil, err
	}

	// Get Go version
	goVersion, err := utils.ReadInputWithDefault("Go version", "1.21")
	if err != nil {
		return nil, err
	}

	// Get author
	author, err := utils.ReadInputWithDefault("Author", "")
	if err != nil {
		return nil, err
	}

	// Get license
	license, err := utils.ReadInputWithOptions("License", []string{
		"MIT",
		"Apache-2.0",
		"GPL-3.0",
		"BSD-3-Clause",
		"None",
	}, "MIT")
	if err != nil {
		return nil, err
	}

	// Get current year
	year := utils.GetCurrentYear()

	return &ProjectData{
		ProjectName:        projectName,
		ProjectDescription: projectDescription,
		ModuleName:         moduleName,
		GoVersion:          goVersion,
		Author:             author,
		License:            license,
		Year:               year,
	}, nil
}

// ListAvailableTemplates lists all available templates
func (c *ProjectCreator) ListAvailableTemplates() ([]string, error) {
	return []string{"default", "go-web", "cli", "library"}, nil
}

// GetTemplateDescription returns the description of a template
func (c *ProjectCreator) GetTemplateDescription(templateName string) (string, error) {
	switch templateName {
	case "default":
		return "A basic Go project with a minimal structure", nil
	case "go-web":
		return "A Go web application with a complete structure for web development", nil
	case "cli":
		return "A command-line interface application with Cobra", nil
	case "library":
		return "A Go library with examples and documentation", nil
	default:
		return "", fmt.Errorf("unknown template: %s", templateName)
	}
}

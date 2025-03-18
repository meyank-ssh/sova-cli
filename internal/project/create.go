package project

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-sova/sova-cli/internal/templates"
	"github.com/go-sova/sova-cli/pkg/utils"
)

type ProjectCreator struct {
	logger         *utils.Logger
	templateLoader *templates.TemplateLoader
	fileGenerator  *templates.FileGenerator
}

func NewProjectCreator(templateDir string) *ProjectCreator {
	loader := templates.NewTemplateLoader(templateDir)
	return &ProjectCreator{
		logger:         utils.NewLoggerWithPrefix(utils.Info, "ProjectCreator"),
		templateLoader: loader,
		fileGenerator:  templates.NewFileGenerator(loader),
	}
}

func (c *ProjectCreator) SetLogger(logger *utils.Logger) {
	c.logger = logger
	c.templateLoader.SetLogger(logger)
	c.fileGenerator.SetLogger(logger)
}

type ProjectData struct {
	ProjectName        string
	ProjectDescription string
	ModuleName         string
	GoVersion          string
	Author             string
	License            string
	Year               string
}

func (c *ProjectCreator) CreateProject(projectName, projectDir, templateName string, force bool) error {
	c.logger.Info("Creating project: %s in directory: %s", projectName, projectDir)
	c.logger.Info("Using template: %s", templateName)

	if utils.DirExists(projectDir) {
		if !force {
			return fmt.Errorf("directory already exists: %s", projectDir)
		}
		c.logger.Warning("Overwriting existing directory: %s", projectDir)
	}

	structure, err := GetProjectStructure(templateName, projectName)
	if err != nil {
		return err
	}

	projectData, err := c.getProjectData(projectName, structure.Description)
	if err != nil {
		return err
	}

	dirs, files := structure.GetAbsolutePaths(projectDir)
	for _, dir := range dirs {
		c.logger.Debug("Creating directory: %s", dir)
		if err := utils.CreateDirIfNotExists(dir); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// Loop through files and handle template category subdirectories
	for filePath, templateName := range files {
		// If the template doesn't have a path separator, check the appropriate category directory
		if filepath.Base(templateName) == templateName {
			// For templates like "go-mod.tpl", look in the category directory first
			categoryTemplate := filepath.Join(templateName, templateName)
			c.logger.Debug("Looking for template in category directory: %s", categoryTemplate)

			// Try to find this template in the appropriate category subdirectory
			if templateName == "default" {
				categoryTemplate = filepath.Join("default", templateName)
			} else if templateName == "cli" {
				categoryTemplate = filepath.Join("cli", templateName)
			} else if templateName == "api" {
				categoryTemplate = filepath.Join("api", templateName)

			// Check if category template exists
			templatePath := filepath.Join(c.templateLoader.GetTemplateDir(), categoryTemplate)
			if utils.FileExists(templatePath) {
				templateName = categoryTemplate
			}
		}

		c.logger.Debug("Generating file: %s from template: %s", filePath, templateName)
		if err := c.fileGenerator.GenerateFile(templateName, filePath, projectData); err != nil {
			return fmt.Errorf("failed to generate file: %w", err)
		}
	}

	c.logger.Info("Project created successfully!")
	return nil
}

func (c *ProjectCreator) getProjectData(projectName, projectDescription string) (*ProjectData, error) {
	// TODO: Get project data from user or default values
	return &ProjectData{
		ProjectName:        projectName,
		ProjectDescription: projectDescription,
		ModuleName:         projectName,
		GoVersion:          "1.21",
		Author:             "Meyank Singh",
		License:            "MIT",
		Year:               fmt.Sprintf("%d", time.Now().Year()),
	}, nil
}

func (c *ProjectCreator) ListAvailableTemplates() ([]string, error) {
	return []string{"default", "go-api", "cli"}, nil
}

func (c *ProjectCreator) GetTemplateDescription(templateName string) (string, error) {
	switch templateName {
	case "default":
		return "A basic Go project with a minimal structure", nil
	case "go-api":
		return "A Go web application with a complete structure for web development", nil
	case "cli":
		return "A command-line interface application with Cobra", nil
	default:
		return "", fmt.Errorf("unknown template: %s", templateName)
	}
}

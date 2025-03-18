package project

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/go-sova/sova-cli/pkg/utils"
	"github.com/go-sova/sova-cli/templates"
)

type ProjectCreator struct {
	logger         *utils.Logger
	templateLoader *templates.TemplateLoader
	fileGenerator  *templates.FileGenerator
}

func NewProjectCreator() *ProjectCreator {
	loader := templates.NewTemplateLoader()
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
		c.logger.Debug("Generating file: %s from template: %s", filePath, templateName)
		if err := c.fileGenerator.GenerateFile(templateName, filePath, projectData); err != nil {
			return fmt.Errorf("failed to generate file: %w", err)
		}
	}

	c.logger.Info("Project created successfully!")
	return nil
}

func (c *ProjectCreator) getProjectData(projectName, projectDescription string) (*ProjectData, error) {
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

func CreateProject(projectName, projectDir string, answers *questions.ProjectAnswers) error {
	structure, err := GetProjectStructure(answers.ProjectType, projectName)
	if err != nil {
		return fmt.Errorf("failed to get project structure: %v", err)
	}

	dirs, files := structure.GetAbsolutePaths(projectDir)

	for _, dir := range dirs {
		if err := utils.CreateDirIfNotExists(dir); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	for path, template := range files {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}

		if err := utils.CopyFile(template, path); err != nil {
			return fmt.Errorf("failed to copy file %s to %s: %v", template, path, err)
		}
	}

	return nil
}

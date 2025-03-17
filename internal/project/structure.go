package project

import (
	"fmt"
	"path/filepath"
)

// ProjectStructure defines the structure of a project
type ProjectStructure struct {
	Name        string
	Description string
	Directories []string
	Files       map[string]string // Map of file path to template name
}

// DefaultProjectStructure returns the default project structure
func DefaultProjectStructure(projectName string) *ProjectStructure {
	return &ProjectStructure{
		Name:        projectName,
		Description: "A project created with Sova CLI",
		Directories: []string{
			"cmd",
			"internal",
			"pkg",
			"docs",
			"scripts",
			"tests",
		},
		Files: map[string]string{
			"main.go":   "go-main.tpl",
			"go.mod":    "go-mod.tpl",
			"README.md": "readme.tpl",
		},
	}
}

// GoWebProjectStructure returns a Go web project structure
func GoWebProjectStructure(projectName string) *ProjectStructure {
	structure := DefaultProjectStructure(projectName)
	structure.Description = "A Go web project created with Sova CLI"

	// Add additional directories
	structure.Directories = append(structure.Directories,
		"api",
		"internal/handlers",
		"internal/middleware",
		"internal/models",
		"internal/database",
		"internal/config",
		"web/templates",
		"web/static/css",
		"web/static/js",
		"web/static/img",
	)

	// Add additional files
	additionalFiles := map[string]string{
		"cmd/server/main.go":                "go-web-main.tpl",
		"internal/config/config.go":         "go-web-config.tpl",
		"internal/handlers/handlers.go":     "go-web-handlers.tpl",
		"internal/middleware/middleware.go": "go-web-middleware.tpl",
		"internal/models/models.go":         "go-web-models.tpl",
		"api/api.go":                        "go-web-api.tpl",
		"web/templates/index.html":          "go-web-index.tpl",
		"web/static/css/style.css":          "go-web-style.tpl",
		"web/static/js/app.js":              "go-web-app-js.tpl",
		"Dockerfile":                        "go-web-dockerfile.tpl",
		"docker-compose.yml":                "go-web-docker-compose.tpl",
		".gitignore":                        "gitignore.tpl",
	}

	for path, template := range additionalFiles {
		structure.Files[path] = template
	}

	return structure
}

// CLIProjectStructure returns a CLI project structure
func CLIProjectStructure(projectName string) *ProjectStructure {
	structure := DefaultProjectStructure(projectName)
	structure.Description = "A CLI project created with Sova CLI"

	// Add additional directories
	structure.Directories = append(structure.Directories,
		"cmd/root",
		"internal/commands",
		"internal/config",
	)

	// Add additional files
	additionalFiles := map[string]string{
		"cmd/root/root.go":          "cli-root.tpl",
		"cmd/version/version.go":    "cli-version.tpl",
		"internal/commands/cmd.go":  "cli-commands.tpl",
		"internal/config/config.go": "cli-config.tpl",
		".gitignore":                "gitignore.tpl",
	}

	for path, template := range additionalFiles {
		structure.Files[path] = template
	}

	return structure
}

// LibraryProjectStructure returns a library project structure
func LibraryProjectStructure(projectName string) *ProjectStructure {
	structure := DefaultProjectStructure(projectName)
	structure.Description = "A Go library created with Sova CLI"

	// Add additional directories
	structure.Directories = append(structure.Directories,
		"examples",
	)

	// Add additional files
	additionalFiles := map[string]string{
		"pkg/library.go":      "lib-main.tpl",
		"examples/example.go": "lib-example.tpl",
		".gitignore":          "gitignore.tpl",
	}

	for path, template := range additionalFiles {
		structure.Files[path] = template
	}

	return structure
}

// GetProjectStructure returns a project structure based on the template name
func GetProjectStructure(templateName, projectName string) (*ProjectStructure, error) {
	switch templateName {
	case "default":
		return DefaultProjectStructure(projectName), nil
	case "go-web":
		return GoWebProjectStructure(projectName), nil
	case "cli":
		return CLIProjectStructure(projectName), nil
	case "library":
		return LibraryProjectStructure(projectName), nil
	default:
		return nil, fmt.Errorf("unknown template: %s", templateName)
	}
}

// GetAbsolutePaths returns the absolute paths for all directories and files
func (s *ProjectStructure) GetAbsolutePaths(baseDir string) ([]string, map[string]string) {
	dirs := make([]string, len(s.Directories))
	files := make(map[string]string)

	// Get absolute paths for directories
	for i, dir := range s.Directories {
		dirs[i] = filepath.Join(baseDir, dir)
	}

	// Get absolute paths for files
	for path, template := range s.Files {
		files[filepath.Join(baseDir, path)] = template
	}

	return dirs, files
}

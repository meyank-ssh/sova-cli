package project

import (
	"fmt"
	"path/filepath"
)

type ProjectStructure struct {
	Name        string
	Description string
	Directories []string
	Files       map[string]string
}

func APIProjectStructure(projectName string) *ProjectStructure {
	structure := &ProjectStructure{
		Name:        projectName,
		Description: "A Go API project created with Sova CLI",
		Directories: []string{
			"cmd",
			"internal",
			"pkg",
			"api",
			"docs",
			"scripts",
			"test",
			"internal/handlers",
			"internal/middleware",
			"internal/models",
			"internal/server",
			"internal/service",
			"routes",
		},
		Files: map[string]string{
			"cmd/main.go":                       "api/main.tpl",
			"internal/config/config.go":         "api/config.tpl",
			"internal/handlers/handlers.go":     "api/handlers.tpl",
			"internal/middleware/middleware.go": "api/middleware.tpl",
			"internal/models/models.go":         "api/models.tpl",
			"internal/server/server.go":         "api/server.tpl",
			"routes/routes.go":                  "api/routes.tpl",
			"internal/service/service.go":       "api/service-init.tpl",
			"internal/service/postgres.go":      "api/postgres.tpl",
			"internal/service/redis.go":         "api/redis.tpl",
			"internal/service/rabbitmq.go":      "api/rabbitmq.tpl",
			"internal/middleware/logging.go":    "api/logging.tpl",
			".env":                              "api/env.tpl",
			"docker-compose.yml":                "api/docker-compose.tpl",
			"go.mod":                            "api/go-mod.tpl",
			".gitignore":                        "api/gitignore.tpl",
		},
	}

	return structure
}

func CLIProjectStructure(projectName string) *ProjectStructure {
	structure := &ProjectStructure{
		Name:        projectName,
		Description: "A CLI project created with Sova CLI",
		Directories: []string{
			"cmd",
			"internal",
			"pkg",
			"docs",
			"scripts",
			"test",
			"cmd/root",
			"internal/commands",
			"internal/config",
		},
		Files: map[string]string{
			"cmd/root/root.go":          "cli/root.tpl",
			"cmd/version/version.go":    "cli/version.tpl",
			"internal/commands/cmd.go":  "cli/commands.tpl",
			"internal/config/config.go": "cli/config.tpl",
			"internal/utils/utils.go":   "cli/utils.tpl",
			".gitignore":                "cli/gitignore.tpl",
		},
	}

	return structure
}

func GetProjectStructure(templateName, projectName string) (*ProjectStructure, error) {
	switch templateName {
	case "api":
		return APIProjectStructure(projectName), nil
	case "cli":
		return CLIProjectStructure(projectName), nil
	default:
		return nil, fmt.Errorf("unknown template: %s", templateName)
	}
}

func (s *ProjectStructure) GetAbsolutePaths(baseDir string) ([]string, map[string]string) {
	dirs := make([]string, len(s.Directories))
	files := make(map[string]string)

	for i, dir := range s.Directories {
		dirs[i] = filepath.Join(baseDir, dir)
	}

	for path, template := range s.Files {
		files[filepath.Join(baseDir, path)] = template
	}

	return dirs, files
}

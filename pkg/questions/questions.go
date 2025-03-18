package questions

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type ProjectAnswers struct {
	ProjectName string
	ProjectType string
	UseZap      bool
	UsePostgres bool
	UseRedis    bool
	UseRabbitMQ bool
}

func AskProjectName() (string, error) {
	var name string
	prompt := &survey.Input{
		Message: "What is your project name?",
		Help:    "The name of your new project",
	}

	err := survey.AskOne(prompt, &name)
	if err != nil {
		return "", fmt.Errorf("failed to get project name: %v", err)
	}

	if name == "" {
		return "", fmt.Errorf("project name cannot be empty")
	}

	return name, nil
}

func AskProjectType() (string, error) {
	var projectType string
	prompt := &survey.Select{
		Message: "What type of project are you building?",
		Options: []string{"api", "cli"},
		Default: "api",
	}

	err := survey.AskOne(prompt, &projectType)
	if err != nil {
		return "", fmt.Errorf("failed to get project type: %v", err)
	}

	return projectType, nil
}

func AskProjectQuestions(projectType string) (*ProjectAnswers, error) {
	answers := &ProjectAnswers{
		ProjectType: projectType,
	}

	switch projectType {
	case "api":
		prompt := &survey.Confirm{
			Message: "Would you like to use zap as a logger?",
			Default: true,
		}
		err := survey.AskOne(prompt, &answers.UseZap)
		if err != nil {
			return nil, err
		}

		prompt = &survey.Confirm{
			Message: "Would you like to use PostgreSQL?",
			Default: true,
		}
		err = survey.AskOne(prompt, &answers.UsePostgres)
		if err != nil {
			return nil, err
		}

		prompt = &survey.Confirm{
			Message: "Would you like to use Redis?",
			Default: false,
		}
		err = survey.AskOne(prompt, &answers.UseRedis)
		if err != nil {
			return nil, err
		}

		prompt = &survey.Confirm{
			Message: "Would you like to use RabbitMQ?",
			Default: false,
		}
		err = survey.AskOne(prompt, &answers.UseRabbitMQ)
		if err != nil {
			return nil, err
		}
	case "cli":
		fmt.Printf("Debug: Processing CLI project type\n")

		prompt := &survey.Confirm{
			Message: "Would you like to use zap as a logger?",
			Default: false,
		}
		err := survey.AskOne(prompt, &answers.UseZap)
		if err != nil {
			return nil, err
		}

		answers.UsePostgres = false
		answers.UseRedis = false
		answers.UseRabbitMQ = false
	default:
		fmt.Printf("Debug: Unsupported project type: '%s'\n", projectType)
		return nil, fmt.Errorf("unsupported project type: %s", projectType)
	}

	return answers, nil
}

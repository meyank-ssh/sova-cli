package questions

import (
	"github.com/AlecAivazis/survey/v2"
)

type ProjectAnswers struct {
	ProjectName string
	ProjectType string
	UseZap      bool
	UseRabbitMQ bool
	UseRedis    bool
	UsePostgres bool
}

func AskProjectQuestions(defaultProjectName string) (*ProjectAnswers, error) {
	answers := &ProjectAnswers{}

	// Project name question
	if defaultProjectName == "" {
		namePrompt := &survey.Input{
			Message: "What is your project name?",
			Help:    "The name of your new project",
		}
		if err := survey.AskOne(namePrompt, &answers.ProjectName); err != nil {
			return nil, err
		}
	} else {
		answers.ProjectName = defaultProjectName
	}

	// Project type question
	projectTypePrompt := &survey.Select{
		Message: "What type of project are you building?",
		Options: []string{"CLI", "API"},
		Default: "CLI",
	}
	if err := survey.AskOne(projectTypePrompt, &answers.ProjectType); err != nil {
		return nil, err
	}

	// If API project is selected, ask additional questions
	if answers.ProjectType == "API" {
		// Zap logging question
		zapPrompt := &survey.Confirm{
			Message: "Do you need Zap for logging?",
			Default: true,
		}
		if err := survey.AskOne(zapPrompt, &answers.UseZap); err != nil {
			return nil, err
		}

		// RabbitMQ question
		rabbitPrompt := &survey.Confirm{
			Message: "Do you require RabbitMQ message broker?",
			Default: false,
		}
		if err := survey.AskOne(rabbitPrompt, &answers.UseRabbitMQ); err != nil {
			return nil, err
		}

		// Redis question
		redisPrompt := &survey.Confirm{
			Message: "Do you want to use Redis?",
			Default: false,
		}
		if err := survey.AskOne(redisPrompt, &answers.UseRedis); err != nil {
			return nil, err
		}

		// Postgres question
		postgresPrompt := &survey.Confirm{
			Message: "Do you want to use PostgreSQL?",
			Default: false,
		}
		if err := survey.AskOne(postgresPrompt, &answers.UsePostgres); err != nil {
			return nil, err
		}
	}

	return answers, nil
}

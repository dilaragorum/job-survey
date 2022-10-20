package questionary

import "github.com/AlecAivazis/survey/v2"

type BeginAnswer struct {
	Job string
}

var BeginningQuestion = []*survey.Question{
	{
		Name: "Job",
		Prompt: &survey.Select{
			Message: "Which job do you want to have?",
			Options: JobOptions,
		},
		Validate: survey.Required,
	},
}

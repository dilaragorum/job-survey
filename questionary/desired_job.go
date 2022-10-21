package questionary

import "github.com/AlecAivazis/survey/v2"

type DesiredJobAnswer struct {
	Job string
}

var DesiredJobQuestion = []*survey.Question{
	{
		Name: "Job",
		Prompt: &survey.Select{
			Message: "Which job do you want to have?",
			Options: JobOptions,
		},
		Validate: survey.Required,
	},
}

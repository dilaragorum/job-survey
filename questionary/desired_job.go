package questionary

import "github.com/AlecAivazis/survey/v2"

var DesiredJobQuestion = &survey.Select{
	Message: "Which job do you want to have?",
	Options: JobOptions,
}

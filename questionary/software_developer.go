package questionary

import "github.com/AlecAivazis/survey/v2"

type SDAnswer struct {
	SelfLearning  bool
	Collaboration bool
}

var SoftwareDeveloperQuestions = []*survey.Question{
	{
		Name: "SelfLearning",
		Prompt: &survey.Confirm{
			Message: "Do you want to be life long learner?",
		},
		Validate: survey.Required,
	},
	{
		Name: "Collaboration",
		Prompt: &survey.Confirm{
			Message: "Do you like to collaborate with other people and are you open to gain different point of view from them?",
		},
		Validate: survey.Required,
	},
}

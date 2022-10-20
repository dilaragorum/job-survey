package questionary

import "github.com/AlecAivazis/survey/v2"

type BotAnswer struct {
	LocalFlora      bool
	LabExercise     bool
	NewPlantSpecies bool
	Passion         bool
}

var BotanistQuestions = []*survey.Question{
	{
		Name: "LocalFlora",
		Prompt: &survey.Confirm{
			Message: "Do you like to interest studying the local flora around your city?",
		},
		Validate: survey.Required,
	},
	{
		Name: "LabExercise",
		Prompt: &survey.Confirm{
			Message: "Do you like to participate lab exercises that used plants as experimental systems in process of science?",
		},
		Validate: survey.Required,
	},
	{
		Name: "NewPlantSpecies",
		Prompt: &survey.Confirm{
			Message: "Do you excited when new plant species are identified and understood especially in such biologically rich areas as tropical rain forests?",
		},
		Validate: survey.Required,
	},
	{
		Name: "Passion",
		Prompt: &survey.Confirm{
			Message: "Many plant species are especially sensitive to certain pollutants. Do you have passion to study the effects of different types of pollution on plants?",
		},
		Validate: survey.Required,
	},
}

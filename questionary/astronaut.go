package questionary

import (
	"github.com/AlecAivazis/survey/v2"
	"strconv"
)

const (
	BiologicalScienceDegree = "Biological Science"
	ComputerScienceDegree   = "Computer Science"
	EngineeringDegree       = "Engineering"
	MathematicsDegree       = "Mathematics"
	SocialScienceDegree     = "Social Science"
	OtherDegrees            = "Other"
)

var Departments = []string{
	BiologicalScienceDegree,
	ComputerScienceDegree,
	EngineeringDegree,
	MathematicsDegree,
	SocialScienceDegree,
	OtherDegrees,
}

const (
	IncurableEyeSightProblem = "Yes, I have incurable eyesight problem"
	CurableEyeSightProblem   = "Yes, I have eyesight problem but it is curable"
	NoEyeSightProblem        = "No, I have no eyesight problem"
)

var EyeSightOptions = []string{
	IncurableEyeSightProblem,
	CurableEyeSightProblem,
	NoEyeSightProblem,
}

type AstAnswer struct {
	Tall           int
	EyeSight       Select
	BachelorDegree Select
	MasterDegree   Select
}

// TODO:
const (
	MinRestrictionHeight = 149
	MaxRestrictionHeight = 193
)

var AstronautQuestions = []*survey.Question{
	{
		Name:      "Tall",
		Prompt:    &survey.Input{Message: "How tall are you? (cm)"},
		Validate:  survey.Required,
		Transform: transformInt,
	},
	{
		Name: "Eyesight",
		Prompt: &survey.Select{
			Message: "Do you have eyesight problem?",
			Options: EyeSightOptions,
		},
		Validate: survey.Required,
	},
	{
		Name: "BachelorDegree",
		Prompt: &survey.Select{
			Message: "What is your bachelor degree?",
			Options: Departments,
		},
		Validate: survey.Required,
	},
	{
		Name: "MasterDegree",
		Prompt: &survey.Select{
			Message: "What is your master degree/s?",
			Options: Departments,
		},
		Validate: survey.Required,
	},
}

var transformInt = func(ans interface{}) (newAns interface{}) {
	strValue, ok := ans.(string)
	if !ok {
		return 0
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0
	}

	return intValue
}

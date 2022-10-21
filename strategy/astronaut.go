package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	AstronautInappropriateMsgText = "Sorry, There are really strict rules to be selected. Better to prefer another job."
	AstronautAppropriateMsgText   = "You have met the pre-conditions. If you selected, you need to work hard :)"
)

type Astronaut struct {
	validDegree map[string]bool
}

func NewAstronaut() *Astronaut {
	return &Astronaut{validDegree: map[string]bool{
		questionary.BiologicalScienceDegree: true,
		questionary.ComputerScienceDegree:   true,
		questionary.EngineeringDegree:       true,
		questionary.MathematicsDegree:       true,
	}}
}

func (ast *Astronaut) Check(answersBytes []byte) (CheckResult, error) {
	var answers questionary.AstronautAnswer
	err := json.Unmarshal(answersBytes, &answers)
	if err != nil {
		return CheckResult{}, err
	}

	if ok := ast.validDegree[answers.BachelorDegree.Value.(string)]; !ok {
		return CheckResult{Text: AstronautInappropriateMsgText, Status: false}, nil
	}

	if ok := ast.validDegree[answers.MasterDegree.Value.(string)]; !ok {
		return CheckResult{Text: AstronautInappropriateMsgText, Status: false}, nil
	}

	eyeSightVal := answers.EyeSight.Value

	if (eyeSightVal == questionary.AstronautNoEyeSightProblem || eyeSightVal == questionary.AstronautCurableEyeSightProblem) &&
		(answers.Tall > questionary.AstronautMinRestrictionHeight && answers.Tall < questionary.AstronautMaxRestrictionHeight) {
		return CheckResult{Text: AstronautAppropriateMsgText, Status: true}, nil
	}

	return CheckResult{Text: AstronautInappropriateMsgText, Status: false}, nil
}

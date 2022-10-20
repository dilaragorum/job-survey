package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	InappropriateMsgTextAst = "Sorry, There are really strict rules to be selected. Better to prefer another job."
	AppropriateMsgTextAst   = "You have met the pre-conditions. If you selected, you need to work hard :)"
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
	var answers questionary.AstAnswer
	err := json.Unmarshal(answersBytes, &answers)
	if err != nil {
		return CheckResult{}, err
	}

	if ok := ast.validDegree[answers.BachelorDegree.Value.(string)]; !ok {
		return CheckResult{Text: InappropriateMsgTextAst, Status: false}, nil
	}

	if ok := ast.validDegree[answers.MasterDegree.Value.(string)]; !ok {
		return CheckResult{Text: InappropriateMsgTextAst, Status: false}, nil
	}

	eyeSightVal := answers.EyeSight.Value

	if (eyeSightVal == questionary.NoEyeSightProblem || eyeSightVal == questionary.CurableEyeSightProblem) &&
		(answers.Tall > questionary.MinRestrictionHeight && answers.Tall < questionary.MaxRestrictionHeight) {
		return CheckResult{Text: AppropriateMsgTextAst, Status: true}, nil
	}

	return CheckResult{Text: InappropriateMsgTextAst, Status: false}, nil
}

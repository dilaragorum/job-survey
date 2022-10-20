package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	InappropriateMsgTextAst = "Sorry, There are really strict rules to be selected. Better to prefer another job."
	AppropriateMsgTextAst   = "You meet the pre-conditions. If you selected, you need to work hard :)"
)

type Astronaut struct {
	validDegree map[string]bool
}

func NewAstronaut() *Astronaut {
	return &Astronaut{validDegree: map[string]bool{
		questionary.BiologicalScience: true,
		questionary.ComputerScience:   true,
		questionary.Engineering:       true,
		questionary.Mathematics:       true,
	}}
}

func (ast *Astronaut) Check(answersBytes []byte) (questionary.MessageForClient, error) {
	var answers questionary.AstAnswer
	err := json.Unmarshal(answersBytes, &answers)
	if err != nil {
		return questionary.MessageForClient{}, err
	}

	if ok := ast.validDegree[answers.BachelorDegree.Value.(string)]; !ok {
		return questionary.MessageForClient{Text: InappropriateMsgTextAst, Status: false}, nil
	}

	if ok := ast.validDegree[answers.MasterDegree.Value.(string)]; !ok {
		return questionary.MessageForClient{Text: InappropriateMsgTextAst, Status: false}, nil
	}

	eyeSightVal := answers.EyeSight.Value

	if (eyeSightVal == questionary.NoEyeSightProblem || eyeSightVal == questionary.CurableEyeSightProblem) &&
		(answers.Tall > 149 && answers.Tall < 193) {
		return questionary.MessageForClient{Text: AppropriateMsgTextAst, Status: true}, nil
	}

	return questionary.MessageForClient{Text: InappropriateMsgTextAst, Status: false}, nil
}

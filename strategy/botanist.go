package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	BotanistInappropriateMessage = "Becoming botanist really needs strong interest and passion. You need to whip these." +
		"Don't lose your hope."
	BotanistAppropriateMessageForBot = "You are ideal candidate, however you need to work hard to be Botanist."
)

type Botanist struct{}

func (bot *Botanist) Check(answersBytes []byte) (CheckResult, error) {
	var answers questionary.BotanistAnswer
	err := json.Unmarshal(answersBytes, &answers)
	if err != nil {
		return CheckResult{}, err
	}

	if answers.LocalFlora && answers.NewPlantSpecies && answers.LabExercise && answers.Passion {
		return CheckResult{Text: BotanistAppropriateMessageForBot, Status: true}, nil
	}

	return CheckResult{Text: BotanistInappropriateMessage, Status: false}, nil
}

package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	InappropriateMessageForBot = "Becoming a botanist really needs strong interest and passion. You need to whip these." +
		"Don't lose your hope."
	AppropriateMessageForBot = "You are an ideal candidate. However, you need to work hard to be Botanist."
)

type Botanist struct{}

func (bot *Botanist) Check(answersBytes []byte) (questionary.MessageForClient, error) {
	var answers questionary.BotAnswer
	err := json.Unmarshal(answersBytes, &answers)
	if err != nil {
		return questionary.MessageForClient{}, err
	}

	if answers.LocalFlora && answers.NewPlantSpecies && answers.LabExercise && answers.Passion {
		return questionary.MessageForClient{AppropriateMessageForBot, true}, nil
	}

	return questionary.MessageForClient{Text: InappropriateMessageForBot, Status: false}, nil
}

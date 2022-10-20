package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	InappropriateMessageForSoftDev = "Sorry, You need to change your mindset. Don't lose your hope."
	AppropriateMessageForSoftDev   = "You are ideal candidate, however you need to work hard to be Software Developer"
)

type SoftwareDeveloper struct{}

func (sci *SoftwareDeveloper) Check(answersBytes []byte) (questionary.MessageForClient, error) {
	var answers questionary.SDAnswer
	if err := json.Unmarshal(answersBytes, &answers); err != nil {
		return questionary.MessageForClient{}, err
	}

	if answers.SelfLearning && answers.Collaboration {
		return questionary.MessageForClient{Text: AppropriateMessageForSoftDev, Status: true}, nil
	}

	return questionary.MessageForClient{Text: InappropriateMessageForSoftDev, Status: false}, nil
}

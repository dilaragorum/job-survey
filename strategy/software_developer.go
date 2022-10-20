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

func (sci *SoftwareDeveloper) Check(answersBytes []byte) (CheckResult, error) {
	var answers questionary.SDAnswer
	if err := json.Unmarshal(answersBytes, &answers); err != nil {
		return CheckResult{}, err
	}

	if answers.SelfLearning && answers.Collaboration {
		return CheckResult{Text: AppropriateMessageForSoftDev, Status: true}, nil
	}

	return CheckResult{Text: InappropriateMessageForSoftDev, Status: false}, nil
}

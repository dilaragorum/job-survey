package strategy

import (
	"encoding/json"
	"github.com/AlecAivazis/survey/v2"
	"github.com/dilaragorum/job-survey/questionary"
)

const (
	SoftwareDeveloperInappropriateMessage = "You need to change your mindset. Don't lose your hope. Try and try again."
	SoftwareDeveloperAppropriateMessage   = "You are ideal candidate, however you need to work hard to be Software Developer"
)

type SoftwareDeveloper struct{}

func (sci *SoftwareDeveloper) Check(answersBytes []byte) (CheckResult, error) {
	var answers questionary.SDAnswer
	if err := json.Unmarshal(answersBytes, &answers); err != nil {
		return CheckResult{}, err
	}

	if answers.SelfLearning && answers.Collaboration {
		return CheckResult{Text: SoftwareDeveloperAppropriateMessage, Status: true}, nil
	}

	return CheckResult{Text: SoftwareDeveloperInappropriateMessage, Status: false}, nil
}

func (sci *SoftwareDeveloper) GetQuestions() []*survey.Question {
	return questionary.SoftwareDeveloperQuestions
}

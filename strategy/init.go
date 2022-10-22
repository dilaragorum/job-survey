package strategy

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/dilaragorum/job-survey/questionary"
)

var Job = make(StrategyMap)

func init() {
	JobRegister()
}

func JobRegister() {
	Job.Register(questionary.JobBotanist, &Botanist{})
	Job.Register(questionary.JobAstronaut, &Astronaut{})
	Job.Register(questionary.JobSoftwareDeveloper, &SoftwareDeveloper{})
}

type CheckResult struct {
	Text   string
	Status bool
}

type JobChecker interface {
	Check(answersBytes []byte) (CheckResult, error)
	GetQuestions() []*survey.Question
}

type StrategyMap map[string]JobChecker

func (m StrategyMap) Register(key string, value JobChecker) {
	Job[key] = value
}

func (m StrategyMap) GetQuestions(desiredJob string) []*survey.Question {
	return Job[desiredJob].GetQuestions()
}

func (m StrategyMap) Check(desiredJob string, answersBytes []byte) (CheckResult, error) {
	return Job[desiredJob].Check(answersBytes)
}

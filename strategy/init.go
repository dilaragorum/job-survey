package strategy

import "github.com/dilaragorum/job-survey/questionary"

var Job = make(StrategyMap)

func init() {
	JobRegister()
}

func JobRegister() {
	Job.Register(questionary.JobBotanist, &Botanist{})
	Job.Register(questionary.JobAstronaut, NewAstronaut())
	Job.Register(questionary.JobSoftwareDeveloper, &SoftwareDeveloper{})
}

type JobChecker interface {
	Check(answersBytes []byte) (questionary.MessageForClient, error)
}

type StrategyMap map[string]JobChecker

func (m StrategyMap) Register(key string, value JobChecker) {
	Job[key] = value
}

func (m StrategyMap) Check(desiredJob string, answersBytes []byte) (questionary.MessageForClient, error) {
	return Job[desiredJob].Check(answersBytes)
}

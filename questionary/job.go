package questionary

import "github.com/AlecAivazis/survey/v2"

const (
	JobAstronaut         = "Astronaut"
	JobSoftwareDeveloper = "Software Developer"
	JobBotanist          = "Botanist"
	JobAskQuestion       = "Ask"
)

var JobOptions = []string{
	JobAstronaut,
	JobSoftwareDeveloper,
	JobBotanist,
}

var JobQuestionsMap = map[string][]*survey.Question{
	JobAskQuestion:       DesiredJobQuestion,
	JobAstronaut:         AstronautQuestions,
	JobBotanist:          BotanistQuestions,
	JobSoftwareDeveloper: SoftwareDeveloperQuestions,
}

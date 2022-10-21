package cmd

import (
	"encoding/json"
	"github.com/AlecAivazis/survey/v2"
	"github.com/dilaragorum/job-survey/questionary"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"

	"github.com/dilaragorum/job-survey/strategy"
)

var rootCmd = &cobra.Command{
	Use:     "job-survey",
	Short:   "This allows you to decide whether you are suitable for selected job",
	Example: "job-survey",
	RunE: func(cmd *cobra.Command, args []string) error {
		var beginningAnswers questionary.BeginAnswer
		err := survey.Ask(questionary.JobQuestionsMap[questionary.JobAskQuestion], &beginningAnswers)
		if err != nil {
			return err
		}

		answersBytes, err := askQuestions(beginningAnswers.Job)
		if err != nil {
			return err
		}

		message, err := strategy.Job.Check(beginningAnswers.Job, answersBytes)
		if err != nil {
			return err
		}

		surveyResult(message)

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command. Reason: %v", err)
	}
}

func surveyResult(result strategy.CheckResult) {
	if result.Status {
		color.Green("%s", result.Text)
		return
	}

	color.Red("%s", result.Text)
}

func askQuestions(desiredJob string) ([]byte, error) {
	qs := questionary.JobQuestionsMap[desiredJob]

	answers := make(map[string]interface{})
	if err := survey.Ask(qs, &answers); err != nil {
		return nil, err
	}

	answersBytes, _ := json.Marshal(answers)
	return answersBytes, nil
}

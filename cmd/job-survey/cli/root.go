package cli

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
	Short:   "This allows you to see whether you are suitable for selected job",
	Aliases: []string{"js"},
	Example: "job-survey",
	RunE: func(cmd *cobra.Command, args []string) error {
		job, err := startQuestion()
		if err != nil {
			return err
		}

		questions := strategy.Job.GetQuestions(job)

		answersBytes, err := askQuestions(questions)
		if err != nil {
			return err
		}

		message, err := strategy.Job.Check(job, answersBytes)
		if err != nil {
			return err
		}

		surveyResult(message)

		return nil
	},
}

func startQuestion() (string, error) {
	var job string
	err := survey.AskOne(questionary.DesiredJobQuestion, &job)
	return job, err
}

func askQuestions(questions []*survey.Question) ([]byte, error) {
	answers := make(map[string]interface{})
	if err := survey.Ask(questions, &answers); err != nil {
		return nil, err
	}

	// trick: in order to provide dynamism
	answersBytes, _ := json.Marshal(answers)
	return answersBytes, nil
}

func surveyResult(result strategy.CheckResult) {
	if result.Status {
		color.Green("%s", result.Text)
		return
	}

	color.Red("%s", result.Text)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command. Reason: %v", err)
	}
}

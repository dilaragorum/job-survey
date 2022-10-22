package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAstronaut_Check(t *testing.T) {
	t.Run("Appropriate_Answers", func(t *testing.T) {
		// Given
		astronaut := Astronaut{}
		answer := questionary.AstronautAnswer{
			Tall: 165,
			EyeSight: questionary.Select{
				Index: 2,
				Value: questionary.AstronautNoEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "You have met the pre-conditions. If you selected, you need to work hard :)"
		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.True(t, check.Status)
	})
	t.Run("When_Tall_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := Astronaut{}
		answer := questionary.AstronautAnswer{
			Tall: 120,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.AstronautNoEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})
	t.Run("When_Bachelor_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := Astronaut{}
		answer := questionary.AstronautAnswer{
			Tall: 120,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.AstronautCurableEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.SocialScienceDegree,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})
	t.Run("When_Master_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := Astronaut{}
		answer := questionary.AstronautAnswer{
			Tall: 120,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.AstronautCurableEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.OtherDegrees,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})
	t.Run("When_Have_Incurable_Eyesight_Problem", func(t *testing.T) {
		// Given
		astronaut := Astronaut{}
		answer := questionary.AstronautAnswer{
			Tall: 165,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.AstronautIncurableEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScienceDegree,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.OtherDegrees,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})
}

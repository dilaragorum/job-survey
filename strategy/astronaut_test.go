package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

// go build -o calculate    ./calculate

// TODO: goreleaser: https://www.youtube.com/watch?v=GG8m-aJuptk
//
//	https://github.com/Abdulsametileri/vX
//
// asciinema rec -- exit
func TestAstronaut_Check(t *testing.T) {
	t.Run("Appropriate_Answers", func(t *testing.T) {
		// Given
		astronaut := NewAstronaut()
		answer := questionary.AstAnswer{
			Tall: 165,
			EyeSight: questionary.Select{
				Index: 2,
				Value: questionary.NoEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "You meet the pre-conditions. If you selected, you need to work hard :)"
		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.True(t, check.Status)
	})
	t.Run("When_Tall_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := NewAstronaut()
		answer := questionary.AstAnswer{
			Tall: 120,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.NoEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
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
		astronaut := NewAstronaut()
		answer := questionary.AstAnswer{
			Tall: 120,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.CurableEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.SocialScience,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
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
		astronaut := NewAstronaut()
		answer := questionary.AstAnswer{
			Tall: 120,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.CurableEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.Other,
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
		astronaut := NewAstronaut()
		answer := questionary.AstAnswer{
			Tall: 165,
			EyeSight: questionary.Select{
				Index: 0,
				Value: questionary.IncurableEyeSightProblem,
			},
			BachelorDegree: questionary.Select{
				Index: 0,
				Value: questionary.BiologicalScience,
			},
			MasterDegree: questionary.Select{
				Index: 0,
				Value: questionary.Other,
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

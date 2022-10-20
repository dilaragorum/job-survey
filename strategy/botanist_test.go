package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBotanist_Check(t *testing.T) {
	t.Run("When_Answers_Are_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      true,
			LabExercise:     true,
			NewPlantSpecies: true,
			Passion:         true,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "You are ideal candidate, however you need to work hard to be Botanist."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.True(t, check.Status)
	})

	t.Run("When_LocalFlora_Answer_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      false,
			LabExercise:     true,
			NewPlantSpecies: true,
			Passion:         true,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_LabExercise_Answer_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      true,
			LabExercise:     false,
			NewPlantSpecies: true,
			Passion:         true,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_NewPlantSpecies_Answer_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      true,
			LabExercise:     true,
			NewPlantSpecies: false,
			Passion:         true,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_Passion_Answer_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      true,
			LabExercise:     true,
			NewPlantSpecies: true,
			Passion:         false,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_LocalFlora_And_Passion_Aswers_Are_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      false,
			LabExercise:     true,
			NewPlantSpecies: true,
			Passion:         false,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_LocalFlora_Passion_And_NewPlants_Answers_Are_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      false,
			LabExercise:     true,
			NewPlantSpecies: false,
			Passion:         false,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_All_Answers_Are_Not_Appropriate", func(t *testing.T) {
		// Given
		botanist := Botanist{}

		answer := questionary.BotAnswer{
			LocalFlora:      false,
			LabExercise:     false,
			NewPlantSpecies: false,
			Passion:         false,
		}

		answerByte, _ := json.Marshal(answer)

		// When
		check, _ := botanist.Check(answerByte)
		expectedText := "Becoming botanist really needs strong interest and passion. You need to whip these." +
			"Don't lose your hope."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})
}

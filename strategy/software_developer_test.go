package strategy

import (
	"encoding/json"
	"github.com/dilaragorum/job-survey/questionary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoftwareDeveloper_Check(t *testing.T) {
	t.Run("When_All_Answers_Are_Appropriate", func(t *testing.T) {
		// Given
		developer := SoftwareDeveloper{}

		answer := questionary.SDAnswer{
			SelfLearning:  true,
			Collaboration: true,
		}

		answerBytes, _ := json.Marshal(answer)

		// When
		check, _ := developer.Check(answerBytes)
		expectedText := "You are ideal candidate, however you need to work hard to be Software Developer"

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.True(t, check.Status)
	})

	t.Run("When_Self_Learning_Answer_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		developer := SoftwareDeveloper{}

		answer := questionary.SDAnswer{
			SelfLearning:  false,
			Collaboration: true,
		}

		answerBytes, _ := json.Marshal(answer)

		// When
		check, _ := developer.Check(answerBytes)
		expectedText := "You need to change your mindset. Don't lose your hope. Try and try again."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_Collaboration_Answer_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		developer := SoftwareDeveloper{}

		answer := questionary.SDAnswer{
			SelfLearning:  true,
			Collaboration: false,
		}

		answerBytes, _ := json.Marshal(answer)

		// When
		check, _ := developer.Check(answerBytes)
		expectedText := "You need to change your mindset. Don't lose your hope. Try and try again."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

	t.Run("When_All_Answers_Are_Not_Appropriate", func(t *testing.T) {
		// Given
		developer := SoftwareDeveloper{}

		answer := questionary.SDAnswer{
			SelfLearning:  false,
			Collaboration: false,
		}

		answerBytes, _ := json.Marshal(answer)

		// When
		check, _ := developer.Check(answerBytes)
		expectedText := "You need to change your mindset. Don't lose your hope. Try and try again."

		// Then
		assert.Equal(t, expectedText, check.Text)
		assert.False(t, check.Status)
	})

}

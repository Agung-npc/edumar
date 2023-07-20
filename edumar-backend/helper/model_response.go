package helper

import (
	. "github.com/dwirobbin/edumar-backend/model/domain"
	. "github.com/dwirobbin/edumar-backend/model/web"
)

func ToLoginResponse(user UserDomain, token string) LoginResponse {
	return LoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}
}

func ToRegisterResponse(user UserDomain, email string) RegisterResponse {
	return RegisterResponse{
		Username: user.Username,
		Email:    email,
	}
}

func ToCategoryResponses(categories []CategoryDomain) []CategoryResponse {
	var categoryResponses []CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, CategoryResponse{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return categoryResponses
}

func ToQuizResponses(questions []QuizDomain, incorrectAnswer IncorrectAnswerDomain, category CategoryDomain) []QuizResponse {
	var incorrectAnswerResponses []string
	incorrectAnswerResponses = append(
		incorrectAnswerResponses, incorrectAnswer.OptionOne, incorrectAnswer.OptionTwo,
	)

	var questionResponses []QuizResponse
	for _, question := range questions {
		questionResponses = append(questionResponses, QuizResponse{
			Id:               question.Id,
			Category:         category.Name,
			Question:         question.Question,
			CorrectAnswer:    question.CorrectAnswer,
			IncorrectAnswers: incorrectAnswerResponses,
		})
	}

	return questionResponses
}

func ToAnswersAttemptResponse(result ResultDomain) AnswerAttemptResponse {
	return AnswerAttemptResponse{
		Correct:  result.Correct,
		Wrong:    result.Wrong,
		Duration: result.Duration,
	}
}

func ToScoreBoardResponses(users []UserDomain, scoresBoard []ResultDomain) []ScoreBoardResponse {
	var scoreBoardResponses []ScoreBoardResponse

	for _, scoreBoard := range scoresBoard {
		var username string
		for _, user := range users {
			if user.Id == scoreBoard.UserId {
				username = user.Username
			}
		}

		scoreBoardResponses = append(scoreBoardResponses, ScoreBoardResponse{
			Username: username,
			Score:    scoreBoard.Correct * 10,
			Duration: scoreBoard.Duration,
		})
	}

	return scoreBoardResponses
}

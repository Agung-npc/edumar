package service

import (
	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/model/domain"
	. "github.com/dwirobbin/edumar-backend/model/web"
	. "github.com/dwirobbin/edumar-backend/pkg/repository"
)

type FeServiceImpl struct {
	authRepository AuthRepository
	feRepo         FeRepository
}

func NewFeService(authRepo AuthRepository, feRepo FeRepository) *FeServiceImpl {
	return &FeServiceImpl{
		authRepository: authRepo,
		feRepo:         feRepo,
	}
}

func (srvc *FeServiceImpl) GetCategories() ([]CategoryResponse, error) {
	categories, err := srvc.feRepo.FindCategories()
	PanicIfError(err)

	return ToCategoryResponses(categories), nil
}

func (srvc *FeServiceImpl) GetQuizByCategoryIdWithPagination(categoryId, page, limit uint) ([]QuizResponse, error) {
	questions, err := srvc.feRepo.FindQuizByCategoryIdWithPagination(categoryId, page, limit)
	PanicIfError(err)

	category, err := srvc.feRepo.FindCategoryById(categoryId)
	PanicIfError(err)

	var incorrectAnswer IncorrectAnswerDomain
	for _, question := range questions {
		incorrectAnswerDomain, err := srvc.feRepo.FindIncorrectAnswersByQuizId(question.Id)
		PanicIfError(err)

		incorrectAnswer = incorrectAnswerDomain
	}

	return ToQuizResponses(questions, incorrectAnswer, category), nil
}

func (srvc *FeServiceImpl) PostAnswerAttempt(userId uint, answerAttemptReq AnswerAttemptRequest) (AnswerAttemptResponse, error) {
	var answersAttempt []AnswerAttemptDomain

	for _, answer := range answerAttemptReq.Answers {
		answersAttempt = append(answersAttempt, AnswerAttemptDomain{
			UserId: userId, QuizId: answer.QuizId, Answer: answer.Answer,
		})
	}

	_, err := srvc.feRepo.SaveAnswerAttempt(userId, answersAttempt)
	PanicIfError(err)

	_, err = srvc.feRepo.SaveResult(
		answerAttemptReq.Duration, userId, answerAttemptReq.CategoryId,
	)
	PanicIfError(err)

	result, err := srvc.feRepo.FindResultByCategoryId(answerAttemptReq.CategoryId)
	PanicIfError(err)

	return ToAnswersAttemptResponse(result), nil
}

func (srvc *FeServiceImpl) GetScoresBoardByCategoryId(categoryId uint) ([]ScoreBoardResponse, error) {
	usersResp, err := srvc.authRepository.FindUsers()
	PanicIfError(err)

	var users []UserDomain
	for _, userResp := range usersResp {
		user, err := srvc.authRepository.FindUserById(userResp.Id)
		PanicIfError(err)

		users = append(users, user)
	}

	scoresBoard, err := srvc.feRepo.FindScoresBoardByCategoryId(categoryId)
	PanicIfError(err)

	return ToScoreBoardResponses(users, scoresBoard), nil
}

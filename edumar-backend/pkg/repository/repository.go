package repository

import (
	"database/sql"

	. "github.com/dwirobbin/edumar-backend/model/domain"
)

type AuthRepository interface {
	FindUser(email, password string) (UserDomain, error)
	FindUsers() ([]UserDomain, error)
	FindUserById(id uint) (UserDomain, error)
	Save(user UserDomain, email string) (UserDomain, error)
}

type FeRepository interface {
	FindCategories() ([]CategoryDomain, error)
	FindCategoryById(categoryId uint) (CategoryDomain, error)
	FindQuizByCategoryIdWithPagination(categoryId, page, limit uint) ([]QuizDomain, error)
	FindIncorrectAnswersByQuizId(quizId uint) (IncorrectAnswerDomain, error)
	SaveAnswerAttempt(userId uint, answersAttempt []AnswerAttemptDomain) (bool, error)
	SaveResult(duration string, userId, categoryId uint) (bool, error)
	FindResultByCategoryId(categoryId uint) (ResultDomain, error)
	FindScoresBoardByCategoryId(categoryId uint) ([]ResultDomain, error)
}

type Repository struct {
	AuthRepository
	FeRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
		FeRepository:   NewFeRepository(db),
	}
}

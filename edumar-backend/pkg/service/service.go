package service

import (
	. "github.com/dwirobbin/edumar-backend/model/web"
	. "github.com/dwirobbin/edumar-backend/pkg/repository"
)

type AuthService interface {
	GenerateToken(request LoginRequest) (LoginResponse, error)
	Register(request RegisterRequest, email string) (RegisterResponse, error)
}

type FeService interface {
	GetCategories() ([]CategoryResponse, error)
	GetQuizByCategoryIdWithPagination(categoryId, page, limit uint) ([]QuizResponse, error)
	PostAnswerAttempt(userId uint, req AnswerAttemptRequest) (AnswerAttemptResponse, error)
	GetScoresBoardByCategoryId(categoryId uint) ([]ScoreBoardResponse, error)
}

type Service struct {
	AuthService
	FeService
}

func NewService(repo *Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo.AuthRepository),
		FeService:   NewFeService(repo.AuthRepository, repo.FeRepository),
	}
}

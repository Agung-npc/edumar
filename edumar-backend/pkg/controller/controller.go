package controller

import (
	"net/http"

	. "github.com/dwirobbin/edumar-backend/pkg/service"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type FrontendController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
	GetQuizByCategoryIdWithPagination(w http.ResponseWriter, r *http.Request)
	SubmitAnswersAttempts(w http.ResponseWriter, r *http.Request)
	GetScoresBoardByCategoryId(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	AuthController
	FrontendController
}

func NewController(service *Service) *Controller {
	return &Controller{
		AuthController:     NewAuthController(service.AuthService),
		FrontendController: NewFrontendController(service.FeService),
	}
}

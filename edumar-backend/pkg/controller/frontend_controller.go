package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/middleware"
	. "github.com/dwirobbin/edumar-backend/model/web"
	. "github.com/dwirobbin/edumar-backend/pkg/service"
)

type FeControllerImpl struct {
	feService FeService
}

func NewFrontendController(feService FeService) *FeControllerImpl {
	return &FeControllerImpl{
		feService: feService,
	}
}

func (ctrl *FeControllerImpl) GetCategories(w http.ResponseWriter, r *http.Request) {
	categoryResponses, err := ctrl.feService.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	WriteToResponseBody(w, WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    categoryResponses,
	})
}

func (ctrl *FeControllerImpl) GetQuizByCategoryIdWithPagination(w http.ResponseWriter, r *http.Request) {
	categoryId, _ := strconv.Atoi(r.URL.Query().Get("category_id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	quizResponses, err := ctrl.feService.GetQuizByCategoryIdWithPagination(
		uint(categoryId), uint(page), uint(limit),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	WriteToResponseBody(w, WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    quizResponses,
	})
}

func (ctrl *FeControllerImpl) SubmitAnswersAttempts(w http.ResponseWriter, r *http.Request) {
	userId, _ := GetUserId()
	var answerAttemptReq AnswerAttemptRequest

	err := json.NewDecoder(r.Body).Decode(&answerAttemptReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = answerAttemptReq.ValidateAnswerAttempt()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	answerAttemptResponses, err := ctrl.feService.PostAnswerAttempt(userId, answerAttemptReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	WriteToResponseBody(w, WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    answerAttemptResponses,
	})
}

func (ctrl *FeControllerImpl) GetScoresBoardByCategoryId(w http.ResponseWriter, r *http.Request) {
	categoryId, _ := strconv.Atoi(r.URL.Query().Get("category_id"))

	scoreBoardResponse, err := ctrl.feService.GetScoresBoardByCategoryId(uint(categoryId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	WriteToResponseBody(w, WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    scoreBoardResponse,
	})
}

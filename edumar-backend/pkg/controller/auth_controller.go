package controller

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/model/web"
	. "github.com/dwirobbin/edumar-backend/pkg/service"
)

type AuthControllerImpl struct {
	authService AuthService
}

func NewAuthController(authService AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (ctrl *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = loginRequest.ValidateLogin()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	userLoginResponse, err := ctrl.authService.GenerateToken(loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   userLoginResponse.Token,
		Expires: time.Now().Add(time.Minute * 60),
		Path:    "/",
	})

	WriteToResponseBody(w, WebResponse{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    userLoginResponse,
	})
}

func (ctrl *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = registerRequest.ValidateRegister()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	registerResponse, err := ctrl.authService.Register(registerRequest, registerRequest.Email)
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
		Message: "SUCCESS",
		Data:    registerResponse,
	})
}

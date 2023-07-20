package router

import (
	"fmt"
	"net/http"
	"os"

	. "github.com/dwirobbin/edumar-backend/exception"
	. "github.com/dwirobbin/edumar-backend/middleware"
	. "github.com/dwirobbin/edumar-backend/pkg/controller"
)

func NewRouter(ctrl *Controller) *http.ServeMux {
	router := http.NewServeMux()

	fmt.Println("\t============================")
	fmt.Println("\t[INFO] Running in Port: " + os.Getenv("API_PORT"))
	fmt.Println("\t============================")

	router.Handle("/api/auth/login",
		POST(http.HandlerFunc(ctrl.Login)),
	)

	router.Handle("/api/auth/register",
		POST(http.HandlerFunc(ctrl.Register)),
	)

	router.Handle("/api/home/categories",
		GET(AuthMiddleWare(http.HandlerFunc(ctrl.GetCategories))),
	)

	router.Handle("/api/home/quizzes",
		GET(AuthMiddleWare(http.HandlerFunc(ctrl.GetQuizByCategoryIdWithPagination))),
	)

	router.Handle("/api/home/process-and-result",
		POST(AuthMiddleWare(http.HandlerFunc(ctrl.SubmitAnswersAttempts))),
	)

	router.Handle("/api/home/score-boards",
		GET(AuthMiddleWare(http.HandlerFunc(ctrl.GetScoresBoardByCategoryId))),
	)

	return router
}

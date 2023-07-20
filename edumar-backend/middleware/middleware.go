package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/dwirobbin/edumar-backend/model/web"
	. "github.com/dwirobbin/edumar-backend/pkg/service"
)

const (
	authorizationHeader = "Authorization"
	userIdKey           = "user_id"
)

var userIdVal uint

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    "Need Authorization header",
			})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    "Invalid Authorization header",
			})
			return
		}

		userId, err := ParseToken(r.Context(), headerParts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, userId)
		userIdVal = ctx.Value(userIdKey).(uint)

		next.ServeHTTP(w, r)
	})
}

func GetUserId() (uint, error) {
	if userIdVal == 0 {
		return 0, fmt.Errorf("user id is not set")
	}

	return userIdVal, nil
}

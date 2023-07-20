package service

import (
	"context"
	"fmt"
	"os"
	"time"

	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/model/domain"
	. "github.com/dwirobbin/edumar-backend/model/web"
	. "github.com/dwirobbin/edumar-backend/pkg/repository"
	. "github.com/dwirobbin/edumar-backend/security"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

type AuthServiceImpl struct {
	authRepository AuthRepository
}

func NewAuthService(authRepo AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepo,
	}
}

func (service *AuthServiceImpl) GenerateToken(request LoginRequest) (LoginResponse, error) {
	userDomain, err := service.authRepository.FindUser(
		request.Email, GeneratePasswordHash(request.Password),
	)
	PanicIfError(err)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userDomain.Id,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	PanicIfError(err)

	return ToLoginResponse(userDomain, tokenString), nil
}

func (service *AuthServiceImpl) Register(request RegisterRequest, email string) (RegisterResponse, error) {
	hashed := GeneratePasswordHash(request.Password)
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")

	userDomain := UserDomain{
		Username:  request.Username,
		Email:     "",
		Password:  hashed,
		CreatedAt: time.Now().In(timeLoc),
		UpdatedAt: time.Now().In(timeLoc),
	}

	authRepo, err := service.authRepository.Save(userDomain, email)
	PanicIfError(err)

	return ToRegisterResponse(authRepo, email), nil
}

func ParseToken(ctx context.Context, token string) (uint, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	newCtx := context.WithValue(ctx, "user_id", claims.UserId)
	_ = context.WithValue(newCtx, "props", claims)

	return claims.UserId, nil
}

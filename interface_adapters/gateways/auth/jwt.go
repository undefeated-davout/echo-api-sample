package auth

import (
	"context"
	"time"
	"undefeated-davout/echo-api-sample/entities"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entities.UserID) error
	Load(ctx context.Context, key string) (entities.UserID, error)
}

type JWTer struct {
	Store        Store
	Clocker      entities.Clocker
	JWTSecretKey string
}

type JWTCustomClaims struct {
	Name string `json:"user_name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func NewJWTer(s Store, c entities.Clocker, secretKey string) *JWTer {
	return &JWTer{Store: s, Clocker: c, JWTSecretKey: secretKey}
}

func (j *JWTer) setCustomClaims(name string, role string) *JWTCustomClaims {
	claims := &JWTCustomClaims{
		Name: name,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}
	return claims
}

func (j *JWTer) getSignedToken(claims *JWTCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.JWTSecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *JWTer) GenerateToken(ctx context.Context, user entities.User) (string, error) {
	claims := j.setCustomClaims(user.Name, user.Role)
	token, err := j.getSignedToken(claims)
	if err != nil {
		return "", err
	}

	jwtID := uuid.New().String()
	// Redisに保存
	if err := j.Store.Save(ctx, jwtID, user.ID); err != nil {
		return "", err
	}

	return token, nil
}

func (j *JWTer) GetUserName(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	return claims.Name
}

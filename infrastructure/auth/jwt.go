package auth

import (
	"errors"
	"example.com/gouserservice/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTHandler struct {
	secretKey string
}

func NewJWTHandler(secretKey string) JWTHandler {
	return JWTHandler{secretKey}
}

func (handler *JWTHandler) AccessToken(user *domain.User, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.UserID
	claims["role"] = role 
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(handler.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (handler *JWTHandler) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.secretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}
	userID := claims["user_id"].(string)
	role := claims["role"].(string)
	return handler.AccessToken(&domain.User{UserID: userID}, role)
}

func (handler *JWTHandler) GenerateEmailVerToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.UserID
	claims["role"] = "email_verification"
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(handler.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


func (handler *JWTHandler) ValidateToken(tokenString string) (*domain.User, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.secretKey), nil
	})
	if err != nil {
		return nil, "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, "", errors.New("invalid token")
	}
	user := &domain.User{
		UserID: claims["user_id"].(string),
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil, "", errors.New("role not found in token")
	}
	return user, role, nil
}

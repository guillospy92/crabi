package pkgjwt

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/guillospy92/crabi/internal/core/errors"
	"github.com/guillospy92/crabi/resources"
)

// JWTData struct content claims and data token
type JWTData[T any] struct {
	Payload T
	jwt.RegisteredClaims
}

// GenerateJWTToken generate token
func GenerateJWTToken[T any](payload T, expirationDate time.Time) (tokenString string, err error) {
	claims := JWTData[T]{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(resources.ConfigurationEnv().JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validate token
func ValidateToken[T any](tokenString string) (*T, error) {
	var claims JWTData[T]

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(*jwt.Token) (any, error) {
		return []byte(resources.ConfigurationEnv().JWTSecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			err = errorlogic.GenerateErrorSinceMessage(
				errorlogic.ErrorCode(errorlogic.TokenExpired),
				errorlogic.StatusCode(http.StatusUnauthorized),
				errorlogic.Message("token is expired"),
			)
		}

		return nil, err
	}

	return &claims.Payload, err
}

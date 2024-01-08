package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var APPLICATION_NAME = "Rest Api Fiber First"
var LOGIN_EXPIRATION_DURATION = time.Now().Add(1 * time.Minute)
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("AC9D72E4B61B7935B28D412F72193")

func GenerateToken(username string) string {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(LOGIN_EXPIRATION_DURATION),
		},
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return err.Error()
	}

	return signedToken
}

func VerifyToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}
	return true
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

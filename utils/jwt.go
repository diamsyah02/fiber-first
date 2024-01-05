package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func GenerateToken(username string) string {
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		jwt.MapClaims{},
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return err.Error()
	}

	return signedToken
}

func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return false
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}
	return true
}

package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(ex time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	if ex == 0 {
		claims["exp"] = time.Now().Add(time.Hour * ex).Unix()
	}
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

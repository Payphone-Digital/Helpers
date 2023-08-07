package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateStateOauthCookie(name string, w http.ResponseWriter) string {
	var expiration = time.Now().Add(2 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:     name,
		Value:    state,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return state
}

func CreateJWT(ex time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	if ex == 0 {
		claims["exp"] = time.Now().Add(time.Hour * ex).Unix()
	}
	tokenStr, err := token.SignedString(secret)
	return tokenStr, err
}

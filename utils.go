package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
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

func CreateJWT(secret, id, name, email, telp string, verify bool, ex time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["email"] = email
	claims["telp"] = email
	claims["verify"] = verify
	if ex > 0 {
		claims["exp"] = time.Now().Add(time.Hour * ex).Unix()
	}
	tokenStr, err := token.SignedString(secret)
	return tokenStr, err
}

func IntOnly(str string) (int, error) {
	// Membagi string menjadi angka-angka dan operator
	var nums []string
	var operators []string
	num := ""
	for _, char := range str {
		if char >= '0' && char <= '9' {
			num += string(char)
		} else if IsOperator(char) {
			if num != "" {
				nums = append(nums, num)
				num = ""
			}
			operators = append(operators, string(char))
		}
	}
	if num != "" {
		nums = append(nums, num)
	}

	// Menghitung hasil berdasarkan operator
	result, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(operators); i++ {
		num, err := strconv.Atoi(nums[i+1])
		if err != nil {
			return 0, err
		}
		operator := operators[i]
		switch operator {
		case "+":
			result += num
		case "-":
			result -= num
		case "*":
			result *= num
		case "/":
			result /= num
		}
	}

	// Menampilkan hasil
	return result, nil
}

func Uws(input string) string {
	return strings.ReplaceAll(input, "_", " ")
}

func IsOperator(char rune) bool {
	operators := "+-*/!@#$%^&()-=" // Anda dapat menambahkan operator lain yang mungkin digunakan
	return strings.ContainsRune(operators, char)
}

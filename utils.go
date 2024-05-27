package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
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

func CreateJWT(data map[string]interface{}, secret string, ex time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}
	if ex > 0 {
		claims["exp"] = time.Now().Add(time.Hour * ex).Unix()
	}
	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}

func RandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		randomBytes[i] = charset[randomBytes[i]%byte(len(charset))]
	}

	return string(randomBytes), nil
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

func IsOperator(char rune) bool {
	operators := "+-*/!@#$%^&()-=" // Anda dapat menambahkan operator lain yang mungkin digunakan
	return strings.ContainsRune(operators, char)
}

func Uws(input string) string {
	return strings.ReplaceAll(input, "_", " ")
}

func IsObject(jsonData string) bool {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return false
	}

	_, ok := data.(map[string]interface{})
	return ok
}

// Fungsi untuk mengevaluasi kondisi ==
func Equals(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a == b
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a == b
		}
	case string:
		if b, ok := b.(string); ok {
			return a == b
		}
	}
	return false
}

// Fungsi untuk mengevaluasi kondisi !=
func NotEquals(a, b interface{}) bool {
	return !Equals(a, b)
}

// Fungsi untuk mengevaluasi kondisi >
func GreaterThan(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a > b
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a > b
		}
	}
	return false
}

// Fungsi untuk mengevaluasi kondisi >=
func GreaterThanOrEquals(a, b interface{}) bool {
	return GreaterThan(a, b) || Equals(a, b)
}

// Fungsi untuk mengevaluasi kondisi <
func LessThan(a, b interface{}) bool {
	return !GreaterThan(a, b) && !Equals(a, b)
}

// Fungsi untuk mengevaluasi kondisi <=
func LessThanOrEquals(a, b interface{}) bool {
	return !GreaterThan(a, b)
}

// Fungsi untuk mengevaluasi penjumlahan
func Add(a, b interface{}) (interface{}, error) {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a + b, nil
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a + b, nil
		}
	}
	return nil, errors.New("unsupported types for addition")
}

// Fungsi untuk mengevaluasi pengurangan
func Subtract(a, b interface{}) (interface{}, error) {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a - b, nil
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a - b, nil
		}
	}
	return nil, errors.New("unsupported types for subtraction")
}

// Fungsi untuk mengevaluasi perkalian
func Multiply(a, b interface{}) (interface{}, error) {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a * b, nil
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a * b, nil
		}
	}
	return nil, errors.New("unsupported types for multiplication")
}

// Fungsi untuk mengevaluasi pembagian
func Divide(a, b interface{}) (interface{}, error) {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			if b == 0 {
				return nil, errors.New("division by zero")
			}
			return a / b, nil
		}
	case float64:
		if b, ok := b.(float64); ok {
			if b == 0.0 {
				return nil, errors.New("division by zero")
			}
			return a / b, nil
		}
	}
	return nil, errors.New("unsupported types for division")
}

// Fungsi untuk memeriksa apakah suatu nilai ada dalam slice
func contains(slice []string, key string) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

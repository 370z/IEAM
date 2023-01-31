package utils

import (
	"ieam-backend/models"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashing(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}
	return string(hashed)
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	if err != nil {
		return err
	}
	return nil
}

func Token(user models.Account) string {
	hmacSampleSecret := []byte("IEAM_IT")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"level":  user.Level,
		"exp":    time.Now().Add(24 * 60 * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(hmacSampleSecret)
	return tokenString
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
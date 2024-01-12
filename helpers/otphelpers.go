package helpers

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateOtp() int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(8999)
	return n + 1000
}
func SendOtp(otp, email string) {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")
	to := []string{email}
	message := "Subject: Otp verification\nyour verification otp is " + otp
	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, []byte(message))
	if err != nil {
		fmt.Println("failed to send otp")
		return
	}
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}

func CheckPassword(dbPassword,inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword),[]byte(inputPassword))
}

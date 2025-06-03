package auth

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"sync"
	"time"

	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/utils"
)

type secret struct {
	OTP            string
	ExpirationTime time.Time
}

var (
	otpSecrets       = make(map[string]secret)
	otpLock          sync.RWMutex
	gmailUsername, _ = utils.EnvVars["GMAIL_USERNAME"]
	gmailPassword, _ = utils.EnvVars["GMAIL_PASSWORD"]
)

func GenerateOTP(accID string) string {
	otpLock.Lock()
	defer otpLock.Unlock()

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	otpSecrets[accID] = secret{
		OTP:            otp,
		ExpirationTime: time.Now().Add(5 * time.Minute),
	}
	return otp
}

func ValidateOTP(accID, otp string) error {
	otpLock.Lock()
	defer otpLock.Unlock()

	secret, ok := otpSecrets[accID]
	println(otp)

	if !ok || secret.OTP != otp {
		return errs.ErrWrongOTP
	}

	delete(otpSecrets, accID)

	if time.Now().After(secret.ExpirationTime) {
		return errs.ErrExpiredOTP
	}

	return nil
}

func SendOTPEmail(recipient, otp string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Your OTP Code\n"
	mime := "MIME-Version: 1.0\nContent-Type: text/plain; charset=\"UTF-8\"\n\n"
	body := fmt.Sprintf("Your OTP code is: %s\n\nThis code will expire in 5 minutes.", otp)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", gmailUsername, gmailPassword, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		gmailUsername,
		[]string{recipient},
		message,
	)

	return err
}

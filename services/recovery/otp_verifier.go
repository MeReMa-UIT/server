package recovery

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
)

func sendOTPEmail(recipient, otp string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")

	// Email content
	subject := "Subject: Your OTP Code\n"
	mime := "MIME-Version: 1.0\nContent-Type: text/plain; charset=\"UTF-8\"\n\n"
	body := fmt.Sprintf("Your OTP code is: %s\n\nThis code will expire in 5 minutes.", otp)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{recipient},
		message,
	)

	return err
}

func SendRecoveryEmail(ctx context.Context, req models.AccountRecoverRequest) error {

	trueEmail, err := repo.GetEmailByCitizenID(ctx, req.CitizenID)
	if err != nil {
		return err
	}

	if trueEmail != req.Email {
		return models.ErrEmailDoesNotMatchCitizenID
	}

	otp := generateOTP(req.CitizenID)

	return sendOTPEmail(req.Email, otp)
}

func VerifyRecoveryOTP(ctx context.Context, req models.AccountRecoverConfirmRequest) error {
	return validateOTP(req.CitizenID, req.OTP)
}

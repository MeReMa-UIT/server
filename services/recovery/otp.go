package recovery

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/errors"
)

type secret struct {
	OTP            string
	ExpirationTime time.Time
}

// temporary storage for otp, will be replaced with a database later
var (
	otpSecrets = make(map[string]secret)
	otpLock    sync.RWMutex
)

func generateOTP(citizenID string) string {
	otpLock.Lock()
	defer otpLock.Unlock()

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	otpSecrets[citizenID] = secret{
		OTP:            otp,
		ExpirationTime: time.Now().Add(5 * time.Minute),
	}
	return otp
}

func validateOTP(req models.AccountRecoverConfirmRequest) error {
	otpLock.Lock()
	defer otpLock.Unlock()

	secret, ok := otpSecrets[req.CitizenID]

	if !ok || secret.OTP != req.OTP {
		return errors.ErrWrongOTP
	}

	delete(otpSecrets, req.CitizenID)

	if time.Now().After(secret.ExpirationTime) {
		return errors.ErrExpiredOTP
	}

	return nil
}

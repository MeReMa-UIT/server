package recovery

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	errs "github.com/merema-uit/server/models/errors"
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

func generateOTP(accID string) string {
	otpLock.Lock()
	defer otpLock.Unlock()

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	otpSecrets[accID] = secret{
		OTP:            otp,
		ExpirationTime: time.Now().Add(5 * time.Minute),
	}
	return otp
}

func validateOTP(accID, otp string) error {
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

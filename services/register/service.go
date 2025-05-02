package register

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	"golang.org/x/crypto/bcrypt"
)

// 1 acc -> n patients
// 1 acc -> 1 staff

func RegisterAccount(ctx context.Context, req models.AccountRegisterRequest) error {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(password_hash)
	err = repo.StoreAccountInfo(ctx, req)
	return err
}

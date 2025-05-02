package register

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAccount(ctx context.Context, req models.AccountRegisterRequest) error {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(password_hash)
	err = repo.StoreAccountInfo(ctx, req)
	return err
}

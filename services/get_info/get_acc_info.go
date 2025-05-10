package getinfo

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetAccountInfo(ctx context.Context, authHeader string) (models.AccountInfo, error) {
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
	if err != nil {
		return models.AccountInfo{}, err
	}

	accountInfo, err := repo.GetAccountInfo(ctx, claims.ID)
	if err != nil {
		return models.AccountInfo{}, err
	}

	return accountInfo, nil
}

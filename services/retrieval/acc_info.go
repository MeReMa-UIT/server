package retrieval

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
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

func GetAccountList(ctx context.Context, authHeader string) ([]models.AccountInfo, error) {
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
	if err != nil {
		return nil, err
	}

	if claims.Permission != permission.Admin.String() {
		return nil, errs.ErrPermissionDenied
	}

	return repo.GetAccountList(ctx)
}

package retrieval_services

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func GetDiagnosisList(ctx context.Context, authHeader string) ([]models.DiagnosisInfo, error) {
	_, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	return repo.GetDiagnosisList(ctx)
}

func GetDiagnosisInfo(ctx context.Context, authHeader, icdCode string) (models.DiagnosisInfo, error) {
	_, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.DiagnosisInfo{}, err
	}
	return repo.GetDiagnosisInfo(ctx, icdCode)
}

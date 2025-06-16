package retrieval_services

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func GetMedicationList(ctx context.Context, authHeader string) ([]models.MedicationInfo, error) {
	_, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	return repo.GetMedicationList(ctx)
}

func GetMedicationInfo(ctx context.Context, authHeader string, medicationID string) (models.MedicationInfo, error) {
	_, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.MedicationInfo{}, err
	}
	medicationInfo, err := repo.GetMedicationInfo(ctx, medicationID)
	if err != nil {
		return models.MedicationInfo{}, err
	}
	return medicationInfo, nil
}

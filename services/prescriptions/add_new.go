package prescription_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func AddNewPrescription(ctx context.Context, authHeader string, req models.NewPrescriptionRequest) (models.NewPrescriptionResponse, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return models.NewPrescriptionResponse{}, err
	}
	if claims.Permission != permission.Doctor.String() {
		return models.NewPrescriptionResponse{}, errs.ErrPermissionDenied
	}
	return repo.StorePrescription(ctx, req)
}

func AddPrescriptionDetails(ctx context.Context, authHeader, prescriptionID string, detail []models.PrescriptionDetailInfo) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	for _, detail := range detail {
		if detail.MorningDosage < 0 || detail.AfternoonDosage < 0 || detail.EveningDosage < 0 || detail.DurationDays <= 0 || (detail.MorningDosage == 0 && detail.AfternoonDosage == 0 && detail.EveningDosage == 0) {
			return errs.ErrInvalidDosage
		}
		if detail.TotalDosage != (detail.MorningDosage+detail.AfternoonDosage+detail.EveningDosage)*float32(detail.DurationDays) {
			return errs.ErrWrongDosageCalulation
		}
	}

	return repo.StorePrescriptionDetails(ctx, prescriptionID, detail)
}

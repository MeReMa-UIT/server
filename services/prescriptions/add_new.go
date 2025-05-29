package prescription_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func AddNewPrescription(ctx context.Context, authHeader string, req models.NewPrescriptionRequest) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}
	for _, detail := range req.Details {
		if detail.MorningDosage < 0 || detail.AfternoonDosage < 0 || detail.EveningDosage < 0 || detail.DurationDays <= 0 || (detail.MorningDosage == 0 && detail.AfternoonDosage == 0 && detail.EveningDosage == 0) {
			return errs.ErrInvalidDosage
		}
		if detail.TotalDosage != (detail.MorningDosage+detail.AfternoonDosage+detail.EveningDosage)*float32(detail.DurationDays) {
			return errs.ErrWrongDosageCalulation
		}
	}
	return repo.StorePrescription(ctx, req)
}

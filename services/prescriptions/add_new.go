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
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}
	var totalDosage []float32
	for _, detail := range req.Details {
		totalDosage = append(totalDosage, (detail.MorningDosage+detail.AfternoonDosage+detail.EveningDosage)*float32(detail.DurationDays))
	}
	return repo.StorePrescription(ctx, req, totalDosage)
}

package prescription_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func ConfirmReceivingPrescription(ctx context.Context, authHeader, prescriptionID string) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}
	return repo.ComfirmReceivingPrescription(ctx, prescriptionID)
}

func UpdatePrescription(ctx context.Context, authHeader, prescriptionID string, req models.PrescriptionUpdateRequest) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	return repo.UpdatePrescription(ctx, prescriptionID, req)
}

func UpdatePrescriptionDetail(ctx context.Context, authHeader, prescriptionID, medID string, detail models.PrescriptionDetailInfo) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	if detail.MorningDosage < 0 || detail.AfternoonDosage < 0 || detail.EveningDosage < 0 || detail.DurationDays <= 0 || (detail.MorningDosage == 0 && detail.AfternoonDosage == 0 && detail.EveningDosage == 0) {
		return errs.ErrInvalidDosage
	}
	if detail.TotalDosage != (detail.MorningDosage+detail.AfternoonDosage+detail.EveningDosage)*float32(detail.DurationDays) {
		return errs.ErrWrongDosageCalulation
	}

	return repo.UpdatePrescriptionDetail(ctx, prescriptionID, medID, detail)
}

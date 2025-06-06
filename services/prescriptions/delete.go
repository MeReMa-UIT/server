package prescription_services

import (
	"context"

	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func DeletePrescriptionDetail(ctx context.Context, authHeader, prescriptionID, detailID string) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	return repo.DeletePrescriptionDetail(ctx, prescriptionID, detailID)
}

package registration

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func RegisterPatient(ctx context.Context, req models.PatientRegistrationRequest, authHeader string) error {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if claims.Permission != permission.PatientRegistration.String() {
		return errs.ErrPermissionDenied
	}
	err = repo.StorePatientInfo(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

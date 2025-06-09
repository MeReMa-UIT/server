package registration_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func RegisterStaff(ctx context.Context, req models.StaffRegistrationRequest, authHeader string) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.StaffRegistration.String() {
		return errs.ErrPermissionDenied
	}
	err = repo.StoreStaffInfo(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

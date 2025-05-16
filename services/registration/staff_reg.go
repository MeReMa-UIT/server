package registration

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func RegisterStaff(ctx context.Context, req models.StaffRegistrationRequest, authHeader string) error {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
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

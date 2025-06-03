package info_update

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func UpdateStaffInfo(ctx context.Context, authHeader string, staffID string, updatedInfo models.StaffInfoUpdateRequest) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))

	if err != nil {
		return err
	}

	if claims.Permission != permission.Admin.String() {
		return errs.ErrPermissionDenied
	}
	return repo.UpdateStaffInfo(ctx, staffID, updatedInfo)
}

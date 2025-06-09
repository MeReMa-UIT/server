package info_update_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func UpdateStaffInfo(ctx context.Context, authHeader string, staffID string, updatedInfo models.StaffInfoUpdateRequest) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))

	if err != nil {
		return err
	}

	if claims.Permission != permission.Admin.String() {
		return errs.ErrPermissionDenied
	}
	return repo.UpdateStaffInfo(ctx, staffID, updatedInfo)
}

package retrieval

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetStaffList(ctx context.Context, authHeader string) ([]models.StaffInfo, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	if claims.Permission != permission.Admin.String() {
		return nil, errs.ErrPermissionDenied
	}
	return repo.GetStaffList(ctx)
}

func GetStaffInfo(ctx context.Context, authHeader string, staffID string) (models.StaffInfo, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return models.StaffInfo{}, err
	}

	var info models.StaffInfo
	switch claims.Permission {
	case permission.Admin.String():
		info, err = repo.GetStaffInfo(ctx, staffID, "0")
	case permission.Doctor.String(), permission.Receptionist.String():
		info, err = repo.GetStaffInfo(ctx, "0", claims.ID)
	default:
		return models.StaffInfo{}, errs.ErrPermissionDenied
	}

	if err != nil {
		return models.StaffInfo{}, err
	}

	return info, nil
}

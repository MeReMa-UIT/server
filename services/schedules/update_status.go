package schedule_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func UpdateScheduleStatus(ctx context.Context, authHeader string, req models.UpdateScheduleStatusRequest) error {
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if claims.Permission != permission.Receptionist.String() {
		return errs.ErrPermissionDenied
	}
	if req.NewStatus != models.ScheduleStatus.Completed && req.NewStatus != models.ScheduleStatus.Cancelled {
		return errs.ErrInvalidScheduleStatus
	}
	return repo.UpdateScheduleStatus(ctx, req)
}

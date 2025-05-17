package schedule_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func BookSchedule(ctx context.Context, authHeader string, req models.ScheduleBookingRequest) (models.ScheduleBookingResponse, error) {

	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)

	if err != nil {
		return models.ScheduleBookingResponse{}, err
	}

	if claims.Permission != permission.Patient.String() {
		return models.ScheduleBookingResponse{}, errs.ErrPermissionDenied
	}

	if req.Type != models.ScheduleType.Regular && req.Type != models.ScheduleType.Service {
		return models.ScheduleBookingResponse{}, errs.ErrInvalidExaminationType
	}

	queueNumber, err := repo.GetQueueNumber(ctx, req.ExaminationDate)
	if err != nil {
		return models.ScheduleBookingResponse{}, err
	}

	createdSchedule, err := repo.CreateSchedule(ctx, req, queueNumber, claims.ID)
	if err != nil {
		return models.ScheduleBookingResponse{}, err
	}

	return createdSchedule, nil
}

package comm_services

import (
	"context"
	"strconv"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func SendMessage(ctx context.Context, authHeader string, message models.SendingMessage) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Patient.String() {
		return errs.ErrPermissionDenied
	}
	accID, _ := strconv.Atoi(claims.ID)
	if message.ToAccID == accID {
		return errs.ErrInvalidRecipient
	}
	return repo.StoreMessage(ctx, claims.ID, message)
}

func LoadConversation(ctx context.Context, authHeader, contactID string) ([]models.Message, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Patient.String() {
		return nil, errs.ErrPermissionDenied
	}
	return repo.LoadConversation(ctx, claims.ID, contactID)
}

package record_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
)

func DeleteRecordAttachment(ctx context.Context, authHeader string, recordID string, req models.DeleteRecordAttachmentRequest) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}

	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	prefix := utils.GetAttachmentPrefix(req.AttachmentFileName)
	if prefix == "" {
		return errs.ErrInvalidAttachmentPrefix
	}

	err = repo.DeleteRecordAttachment(ctx, recordID, prefix, req)
	if err != nil {
		return err
	}

	return nil
}

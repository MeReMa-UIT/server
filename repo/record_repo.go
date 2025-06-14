package repo

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/utils"
)

func GetRecordIDListByAccID(ctx context.Context, accID string) ([]int, error) {
	const query = `
		SELECT record_id
		FROM records r JOIN patients p ON r.patient_id = p.patient_id
		WHERE acc_id = $1::BIGINT
	`

	var recordIDList []int
	rows, _ := dbpool.Query(ctx, query, accID)
	recordIDList, err := pgx.AppendRows(recordIDList, rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return recordIDList, nil
}

func StoreMedicalRecord(ctx context.Context, doctorID int, req *models.NewMedicalRecordRequest, additionalInfo models.ExtractedRecordInfo) (models.NewMedicalRecordResponse, error) {
	const query = `
		INSERT INTO records (doctor_id, patient_id, type_id, record_detail, primary_diagnosis, secondary_diagnosis)
		VALUES ($1, $2, $3, $4, NULLIF($5, ''), NULLIF($6, ''))
		RETURNING record_id
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}
	defer tx.Rollback(ctx)

	var createdRecordID int
	err = tx.QueryRow(ctx, query, doctorID, req.PatientID, req.TypeID, req.RecordDetail, additionalInfo.PrimaryDiagnosis, additionalInfo.SecondaryDiagnosis).Scan(&createdRecordID)
	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	return models.NewMedicalRecordResponse{RecordID: createdRecordID}, nil
}

func StoreMedicalRecordAttachments(ctx context.Context, recordID string, attachments []*multipart.FileHeader, attType string) error {
	const query = `
		INSERT INTO record_attachments (record_id, type, file_path)
		VALUES ($1, $2, $3)
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	storagePath := path.Join(os.ExpandEnv(utils.EnvVars["FILE_STORAGE_PATH"]), "records", recordID, attType)

	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return fmt.Errorf("Failed to create storage directory: %w", err)
	}

	for _, attachment := range attachments {

		attPath := path.Join(storagePath, attachment.Filename)
		err := utils.StoreFile(attachment, attPath)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, query, recordID, attType, attPath)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func GetMedicalRecordList(ctx context.Context, patientID int) ([]models.MedicalRecordBriefInfo, error) {
	const query = `
		SELECT record_id, patient_id, doctor_id, type_id, primary_diagnosis, secondary_diagnosis
		FROM records
		WHERE (patient_id = $1 OR $1::BIGINT = 0)
	`

	rows, _ := dbpool.Query(ctx, query, patientID)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MedicalRecordBriefInfo])

	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetMedicalRecordInfo(ctx context.Context, recordID string) (models.MedicalRecordInfo, error) {
	const query = `
		SELECT record_id, patient_id, doctor_id, type_id, record_detail, created_at, expired_at
		FROM records
		WHERE record_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, recordID)
	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.MedicalRecordInfo])

	if err != nil {
		return models.MedicalRecordInfo{}, err
	}
	return info, nil
}

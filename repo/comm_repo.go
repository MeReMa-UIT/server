package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
)

func GetContactListDoctor(ctx context.Context, accID string) ([]models.ContactInfo, error) {
	const query = `
		SELECT DISTINCT ON (p.acc_id) p.acc_id, p.full_name, $1::text as role
		FROM records r JOIN (SELECT * FROM patients ORDER BY date_of_birth) p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
		WHERE s.acc_id = $2
		ORDER BY p.acc_id
	`

	rows, _ := dbpool.Query(ctx, query, permission.Patient.String(), accID)
	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ContactInfo])
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func GetContactListPatient(ctx context.Context, accID string) ([]models.ContactInfo, error) {
	const query = `
		SELECT DISTINCT s.acc_id, s.full_name, $1::text as role
		FROM records r JOIN patients p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
		WHERE p.acc_id = $2
	`

	rows, _ := dbpool.Query(ctx, query, permission.Doctor.String(), accID)

	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ContactInfo])
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func CheckValidContact(ctx context.Context, accID1, accID2 string) error {
	const query = `
		SELECT EXISTS (
			SELECT 1
			FROM records r JOIN patients p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
			WHERE (p.acc_id = $1 AND s.acc_id = $2) OR (p.acc_id = $2 AND s.acc_id = $1)
		)
	`
	var exists bool
	err := dbpool.QueryRow(ctx, query, accID1, accID2).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errs.ErrInvalidRecipient
	}
	return nil
}

func StoreMessage(ctx context.Context, senderID string, message models.SendingMessage) error {
	const query = `
		INSERT INTO messages (from_acc_id, to_acc_id, content)
		VALUES ($1, $2, $3)
	`
	err := CheckValidContact(ctx, senderID, fmt.Sprint(message.ToAccID))
	if err != nil {
		return errs.ErrInvalidRecipient
	}

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, query, senderID, message.ToAccID, message.Content)
	if err != nil {
		return err
	}
	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func LoadConversation(ctx context.Context, accID, contactID string) ([]models.Message, error) {
	const query = `
		SELECT content, sent_at, from_acc_id  as sender_id 
		FROM messages
		WHERE (from_acc_id = $1 AND to_acc_id = $2) OR (from_acc_id = $2 AND to_acc_id = $1)
		ORDER BY sent_at
	`
	rows, _ := dbpool.Query(ctx, query, accID, contactID)
	messages, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Message])
	if err != nil {
		return nil, err
	}
	return messages, nil
}

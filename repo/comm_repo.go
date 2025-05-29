package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

func GetContactListDoctor(ctx context.Context, accID string) ([]models.ContactInfo, error) {
	const query = `
		SELECT DISTINCT p.acc_id, p.full_name
		FROM records r JOIN patients p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
		WHERE s.acc_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, accID)
	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ContactInfo])
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func GetContactListPatient(ctx context.Context, accID string) ([]models.ContactInfo, error) {
	const query = `
		SELECT DISTINCT s.acc_id, s.full_name
		FROM records r JOIN patients p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
		WHERE p.acc_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, accID)
	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ContactInfo])
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func StoreMessage(ctx context.Context, senderID string, message models.SendingMessage) error {
	const query = `
		INSERT INTO messages (from_acc_id, to_acc_id, content)
		VALUES ($1, $2, $3)
	`

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
		SELECT content, sent_at
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

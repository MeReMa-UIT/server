package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
)

func CheckConversationExists(ctx context.Context, accID1, accID2 int64) error {
	const query = `
		SELECT EXISTS (
			SELECT 1
			FROM conversations
			WHERE (acc_id_1 = $1 AND acc_id_2 = $2) OR (acc_id_1 = $2 AND acc_id_2 = $1)
		)
	`
	row := dbpool.QueryRow(ctx, query, accID1, accID2)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return errs.ErrConversationNotFound
	}

	return nil
}

func GetConversationList(ctx context.Context, accID string) ([]models.Conversation, error) {
	const query = `
		SELECT conversation_id, acc_id_1, acc_id_2, last_message_at
		FROM conversations
		WHERE acc_id_1 = $1 or acc_id_2 = $1
	`
	rows, _ := dbpool.Query(ctx, query, accID)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Conversation])

	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetConversationMessage(ctx context.Context, conversationID string) ([]models.Message, error) {
	const query = `
		SELECT *
		FROM messages
		WHERE conversation_id = $1
		ORDER BY sent_at
	`
	rows, _ := dbpool.Query(ctx, query, conversationID)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Message])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func StoreMessage(ctx context.Context, message models.NewMessage, senderAccID string) (models.Message, error) {
	const query = `
		INSERT INTO messages (conversation_id, sender_acc_id, content)
		VALUES ($1, $2, $3)
		RETURNING *
	`
	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return models.Message{}, err
	}
	defer tx.Rollback(ctx)

	rows, _ := tx.Query(ctx, query, message.ConversationID, senderAccID, message.Content)
	storedMessage, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.Message])

	if err != nil {
		return models.Message{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return models.Message{}, err
	}

	return storedMessage, nil
}

func UpdateConversationLastMessage(ctx context.Context, conversationID int64) error {
	const query = `
		UPDATE conversations
		SET last_message_at = NOW()
		WHERE conversation_id = $1
	`
	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, query, conversationID); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

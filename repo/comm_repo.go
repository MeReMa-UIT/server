package repo

import (
	"context"
	"time"

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
		SELECT
			conversation_id,
			CASE
				WHEN acc_id_1 = $1 THEN acc_id_2
				ELSE acc_id_1
			END AS partner_acc_id,
			last_message_at
		FROM conversations
		WHERE acc_id_1 = $1 OR acc_id_2 = $1;
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

func UpdateConversationLastMessage(ctx context.Context, conversationID int64, lastMessageTime time.Time) error {
	const query = `
		UPDATE conversations
		SET last_message_at = $1
		WHERE conversation_id = $2
	`
	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, query, lastMessageTime, conversationID); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func UpdateMessageSeenStatus(ctx context.Context, conversationID int64, partnerAccID int64, readTime time.Time) error {
	const query = `
		UPDATE messages
		SET is_seen = TRUE
		WHERE conversation_id = $1 AND sender_acc_id = $2 AND is_seen = FALSE AND sent_at <= $3 
	`
	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, query, conversationID, partnerAccID, readTime); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

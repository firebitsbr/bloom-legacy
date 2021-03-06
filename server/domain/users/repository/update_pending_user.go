package repository

import (
	"context"

	"gitlab.com/bloom42/bloom/server/db"
	"gitlab.com/bloom42/bloom/server/domain/users"
	"gitlab.com/bloom42/bloom/server/errors"
	"gitlab.com/bloom42/gobox/log"
)

// UpdatePendingUser update a pending user in the database
func (repo *UsersRepository) UpdatePendingUser(ctx context.Context, db db.Queryer, pendingUser users.PendingUser) error {
	query := `UPDATE pending_users SET
		updated_at = $1, verified_at = $2, failed_attempts = $3, code_hash = $4
		WHERE id = $5`
	_, err := db.Exec(ctx, query, pendingUser.UpdatedAt, pendingUser.VerifiedAt, pendingUser.FailedAttempts,
		pendingUser.CodeHash, pendingUser.ID)

	if err != nil {
		logger := log.FromCtx(ctx)
		const errMessage = "users.UpdatePendingUser: updating pending user"
		logger.Error(errMessage, log.Err("error", err), log.UUID("pending_user.id", pendingUser.ID))
		err = errors.Internal(errMessage, err)
	}

	return err
}

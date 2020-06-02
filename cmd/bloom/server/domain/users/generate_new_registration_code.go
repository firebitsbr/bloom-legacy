package users

import (
	"context"

	"time"

	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/gobox/crypto"
	"gitlab.com/bloom42/gobox/rz"
)

func generateNewRegistrationCode(ctx context.Context, tx *sqlx.Tx, pendingUser *PendingUser) (string, error) {
	logger := rz.FromCtx(ctx)
	var err error

	now := time.Now().UTC()
	verificationCode, err := crypto.RandAlphabet([]byte(USER_VERIFICATION_CODE_ALPHABET), 8)
	if err != nil {
		logger.Error("users.GenerateNewRegistrationCode: error generating verification code", rz.Err(err))
		return "", NewError(ErrorSendingNewRegistrationCode)
	}

	// TODO: update params
	codeHash, err := crypto.HashPassword(verificationCode, PENDING_USER_CODE_HASH_PARAMS)
	if err != nil {
		logger.Error("users.GenerateNewRegistrationCode: hashing verification code", rz.Err(err))
		return "", NewError(ErrorSendingNewRegistrationCode)
	}

	pendingUser.VerificationCodeHash = codeHash
	pendingUser.UpdatedAt = now

	queryUpdatePendingUser := "UPDATE pending_users SET verification_code_hash = $1, updated_at = $2 WHERE id = $3"
	_, err = tx.Exec(queryUpdatePendingUser, pendingUser.VerificationCodeHash, pendingUser.UpdatedAt, pendingUser.ID)
	if err != nil {
		logger.Error("users.GenerateNewRegistrationCode: updateing pending user", rz.Err(err), rz.String("pending_user.id", pendingUser.ID.String()))
		return "", NewError(ErrorSendingNewRegistrationCode)
	}
	return string(verificationCode), nil
}

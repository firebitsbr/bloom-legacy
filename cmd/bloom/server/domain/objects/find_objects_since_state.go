package objects

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/db"
	"gitlab.com/bloom42/lily/rz"
	"gitlab.com/bloom42/lily/uuid"
)

func FindObjectSinceState(ctx context.Context, tx *sqlx.Tx, sinceState int64, userID, groupID *uuid.UUID) ([]Object, error) {
	logger := rz.FromCtx(ctx)
	ret := []Object{}
	var err error
	var whereID uuid.UUID

	if userID == nil && groupID == nil {
		logger.Error("objects.FindObjectSinceState: userID and groupID can't be both nil")
		err = NewError(ErrorInternal)
		return ret, err
	}

	if userID != nil && groupID != nil {
		logger.Error("objects.FindObjectSinceState: userID and groupID can't be both non nil")
		err = NewError(ErrorInternal)
		return ret, err
	}

	queryFind := "SELECT * FROM objects WHERE updated_at_state > $1 AND"
	if userID != nil {
		queryFind += " user_id = $2"
		whereID = *userID
	} else {
		queryFind += " group_id = $2"
		whereID = *groupID
	}
	if tx == nil {
		err = db.DB.Select(&ret, queryFind, whereID)
	} else {
		err = tx.Select(&ret, queryFind, whereID)
	}
	if err != nil {
		logger.Error("objects.FindObjectSinceState: finding objects", rz.Err(err),
			rz.String("whereID", whereID.String()))
		return ret, NewError(ErrorInternal)
	}

	return ret, nil
}

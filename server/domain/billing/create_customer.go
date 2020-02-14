package billing

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/bloom/server/domain/users"
	"gitlab.com/bloom42/libs/rz-go"
)

// creates a customer only for us
func CreateCustomer(ctx context.Context, tx *sqlx.Tx, user *users.User, userID, groupID *string) (*Customer, error) {
	logger := rz.FromCtx(ctx)
	var err error
	var ret *Customer

	// validate params
	if user == nil {
		logger.Error("", rz.Err(NewError(ErrorUserIsNull)))
		return ret, NewError(ErrorCreatingCustomer)
	}

	// find default plan
	defaultPlan, err := FindDefaultPlan(ctx, tx)
	if err != nil {
		logger.Error("billing.CreateCustomer: finding default plan", rz.Err(err))
		return ret, err
	}

	// create customer
	now := time.Now().UTC()
	newUuid := uuid.New()
	ret = &Customer{
		ID:            newUuid.String(),
		CreatedAt:     now,
		UpdatedAt:     now,
		PlanID:        defaultPlan.ID,
		StripeID:      nil,
		Email:         user.Email,
		UsedStorage:   0,
		PlanUpdatedAt: now,
		UserID:        userID,
		GroupID:       groupID,
	}

	queryCreateCustomer := `INSERT INTO billing_customers
		(id, created_at, updated_at, plan_id, stripe_id, email, used_storage, plan_updated_at, user_id, group_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err = tx.Exec(queryCreateCustomer, ret.ID, ret.CreatedAt, ret.UpdatedAt, ret.PlanID,
		ret.StripeID, ret.Email, ret.UsedStorage, ret.PlanUpdatedAt, ret.UserID, ret.GroupID)
	if err != nil {
		logger.Error("billing.CreateCustomer: inserting new customer", rz.Err(err))
		return ret, NewError(ErrorCreatingCustomer)
	}

	return ret, nil
}

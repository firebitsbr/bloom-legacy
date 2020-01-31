package mutation

import (
	"context"
	"time"

	gqlerrors "gitlab.com/bloom42/bloom/server/api/graphql/errors"
	"gitlab.com/bloom42/bloom/server/api/graphql/gqlutil"
	"gitlab.com/bloom42/bloom/server/api/graphql/model"
	"gitlab.com/bloom42/bloom/server/db"
	"gitlab.com/bloom42/bloom/server/domain/users"
	"gitlab.com/bloom42/libs/crypto42-go/rand"
	"gitlab.com/bloom42/libs/rz-go"
)

func (resolver *Resolver) Register(ctx context.Context, input model.RegisterInput) (*model.RegistrationStarted, error) {
	logger := rz.FromCtx(ctx)
	var ret *model.RegistrationStarted
	currentUser := gqlutil.UserFromCtx(ctx)

	if currentUser != nil {
		return ret, gqlerrors.MustNotBeAuthenticated()
	}

	// sleep to prevent spam and bruteforce
	sleep, err := rand.Int64(500, 800)
	if err != nil {
		logger.Error("mutation.Register: generating random int", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	// create pending user
	tx, err := db.DB.Beginx()
	if err != nil {
		logger.Error("mutation.Register: Starting transaction", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	newPendingUser, verificationCode, err := users.CreatePendingUser(ctx, tx, input.DisplayName, input.Email)
	if err != nil {
		tx.Rollback()
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	err = users.SendUserVerificationCode(input.Email, input.DisplayName, verificationCode)
	if err != nil {
		tx.Rollback()
		logger.Error("mutation.Register: Sending confirmation email", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("mutation.Register: Committing transaction", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	ret = &model.RegistrationStarted{
		ID: newPendingUser.ID,
	}

	return ret, nil
}
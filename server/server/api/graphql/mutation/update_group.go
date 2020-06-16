package mutation

import (
	"context"

	"gitlab.com/bloom42/bloom/server/server/api/apiutil"
	"gitlab.com/bloom42/bloom/server/server/api/graphql/gqlerrors"
	"gitlab.com/bloom42/bloom/server/server/api/graphql/model"
	"gitlab.com/bloom42/bloom/server/server/domain/groups"
)

// UpdateGroup is used by groups' admins to update the group info
func (r *Resolver) UpdateGroup(ctx context.Context, input model.GroupInput) (ret *model.Group, err error) {
	currentUser := apiutil.UserFromCtx(ctx)

	if currentUser == nil {
		return ret, gqlerrors.AuthenticationRequired()
	}

	params := groups.UpdateGroupParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}
	group, err := groups.UpdateGroup(ctx, currentUser, params)
	if err != nil {
		err = gqlerrors.New(err)
		return
	}

	ret = &model.Group{
		ID:          &group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   &group.CreatedAt,
	}
	return ret, nil
}

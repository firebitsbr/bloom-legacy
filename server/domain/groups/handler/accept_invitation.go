package handler

import (
	"context"
	rpc "gitlab.com/bloom42/bloom/common/rpc/groups"
)

func (handler Handler) AcceptInvitation(ctx context.Context, params *rpc.AcceptInvitationParams) (*rpc.Empty, error) {
	ret := &rpc.Empty{}

	return ret, nil
}

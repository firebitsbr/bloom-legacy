package handler

import (
	"context"
	rpc "gitlab.com/bloom42/bloom/common/rpc/groups"
)

func (handler Handler) UpdateGroup(ctx context.Context, params *rpc.UpdateGroupParams) (*rpc.Group, error) {
	ret := &rpc.Group{}

	return ret, nil
}

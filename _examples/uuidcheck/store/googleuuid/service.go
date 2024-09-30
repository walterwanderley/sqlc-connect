// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect). DO NOT EDIT.

package googleuuid

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	pb "uuidcheck/api/googleuuid/v1"
	"uuidcheck/api/googleuuid/v1/v1connect"
	"uuidcheck/internal/validation"
)

type Service struct {
	v1connect.UnimplementedGoogleuuidServiceHandler
	querier *Queries
}

func (s *Service) GetProductsByIds(ctx context.Context, req *connect.Request[pb.GetProductsByIdsRequest]) (*connect.Response[pb.GetProductsByIdsResponse], error) {
	var dollar_1 []uuid.UUID
	dollar_1 = make([]uuid.UUID, len(req.Msg.GetDollar_1()))
	for i, s := range req.Msg.GetDollar_1() {
		if v, err := uuid.Parse(s); err != nil {
			err = fmt.Errorf("invalid Dollar_1: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		} else {
			dollar_1[i] = v
		}
	}

	result, err := s.querier.GetProductsByIds(ctx, dollar_1)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "GetProductsByIds")
		return nil, err
	}
	res := new(pb.GetProductsByIdsResponse)
	for _, r := range result {
		res.List = append(res.List, toProduct(r))
	}
	return connect.NewResponse(res), nil
}

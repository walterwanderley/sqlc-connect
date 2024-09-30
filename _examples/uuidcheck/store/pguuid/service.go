// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect). DO NOT EDIT.

package pguuid

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"

	pb "uuidcheck/api/pguuid/v1"
	"uuidcheck/api/pguuid/v1/v1connect"
)

type Service struct {
	v1connect.UnimplementedPguuidServiceHandler
	querier *Queries
}

func (s *Service) GetProductsByIds(ctx context.Context, req *connect.Request[pb.GetProductsByIdsRequest]) (*connect.Response[pb.GetProductsByIdsResponse], error) {
	var dollar_1 []pgtype.UUID
	dollar_1 = make([]pgtype.UUID, len(req.Msg.GetDollar_1()))
	for i, s := range req.Msg.GetDollar_1() {
		if err := dollar_1[i].Scan(s.Value); err != nil {
			return nil, fmt.Errorf("invalid UUID in Dollar_1 at index %d: %w", i, err)
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

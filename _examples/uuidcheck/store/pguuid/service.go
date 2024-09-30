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

	if uuidStr, err := result.MarshalJSON(); err != nil {
		return nil, fmt.Errorf("failed to convert UUID to string: %w", err)
	} else {
		return connect.NewResponse(&pb.CreateUserResponse{Value: wrapperspb.String(string(uuidStr))}), nil
	}
}

func (s *Service) CreateUserReturnAll(ctx context.Context, req *connect.Request[pb.CreateUserReturnAllRequest]) (*connect.Response[pb.CreateUserReturnAllResponse], error) {
	var arg CreateUserReturnAllParams
	if v := req.Msg.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.Msg.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUserReturnAll(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateUserReturnAll")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateUserReturnAllResponse{User: toUser(result)}), nil
}

func (s *Service) CreateUserReturnPartial(ctx context.Context, req *connect.Request[pb.CreateUserReturnPartialRequest]) (*connect.Response[pb.CreateUserReturnPartialResponse], error) {
	var arg CreateUserReturnPartialParams
	if v := req.Msg.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.Msg.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUserReturnPartial(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateUserReturnPartial")
		return nil, err

	}
	return connect.NewResponse(res), nil
}

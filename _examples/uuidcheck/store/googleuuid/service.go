// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect). DO NOT EDIT.

package googleuuid

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	pb "uuidcheck/api/googleuuid/v1"
	"uuidcheck/api/googleuuid/v1/v1connect"
	"uuidcheck/internal/validation"
)

type Service struct {
	v1connect.UnimplementedGoogleuuidServiceHandler
	querier *Queries
}

func (s *Service) CreateProduct(ctx context.Context, req *connect.Request[pb.CreateProductRequest]) (*connect.Response[pb.CreateProductResponse], error) {
	var arg CreateProductParams
	arg.ID = req.Msg.GetId()
	if v := req.Msg.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProduct(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateProduct")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateProductResponse{Value: result}), nil
}

func (s *Service) CreateProductReturnAll(ctx context.Context, req *connect.Request[pb.CreateProductReturnAllRequest]) (*connect.Response[pb.CreateProductReturnAllResponse], error) {
	var arg CreateProductReturnAllParams
	arg.ID = req.Msg.GetId()
	if v := req.Msg.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProductReturnAll(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateProductReturnAll")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateProductReturnAllResponse{Product: toProduct(result)}), nil
}

func (s *Service) CreateProductReturnPartial(ctx context.Context, req *connect.Request[pb.CreateProductReturnPartialRequest]) (*connect.Response[pb.CreateProductReturnPartialResponse], error) {
	var arg CreateProductReturnPartialParams
	arg.ID = req.Msg.GetId()
	if v := req.Msg.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProductReturnPartial(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateProductReturnPartial")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateProductReturnPartialResponse{CreateProductReturnPartialRow: toCreateProductReturnPartialRow(result)}), nil
}

func (s *Service) CreateUser(ctx context.Context, req *connect.Request[pb.CreateUserRequest]) (*connect.Response[pb.CreateUserResponse], error) {
	var arg CreateUserParams
	if v, err := uuid.Parse(req.Msg.GetId()); err != nil {
		err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
		return nil, err
	} else {
		arg.ID = v
	}
	if v := req.Msg.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUser(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateUser")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateUserResponse{Value: result.String()}), nil
}

func (s *Service) CreateUserReturnAll(ctx context.Context, req *connect.Request[pb.CreateUserReturnAllRequest]) (*connect.Response[pb.CreateUserReturnAllResponse], error) {
	var arg CreateUserReturnAllParams
	if v, err := uuid.Parse(req.Msg.GetId()); err != nil {
		err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
		return nil, err
	} else {
		arg.ID = v
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
	if v, err := uuid.Parse(req.Msg.GetId()); err != nil {
		err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
		return nil, err
	} else {
		arg.ID = v
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
	return connect.NewResponse(&pb.CreateUserReturnPartialResponse{CreateUserReturnPartialRow: toCreateUserReturnPartialRow(result)}), nil
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
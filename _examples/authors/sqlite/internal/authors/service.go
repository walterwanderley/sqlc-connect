// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect). DO NOT EDIT.

package authors

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/bufbuild/connect-go"

	pb "authors/api/authors/v1"
	"authors/api/authors/v1/v1connect"
)

type Service struct {
	v1connect.UnimplementedAuthorsServiceHandler
	querier *Queries
}

func (s *Service) CreateAuthor(ctx context.Context, req *connect.Request[pb.CreateAuthorRequest]) (*connect.Response[pb.CreateAuthorResponse], error) {
	var arg CreateAuthorParams
	arg.Name = req.Msg.GetName()
	if v := req.Msg.GetBio(); v != nil {
		arg.Bio = sql.NullString{Valid: true, String: v.Value}
	}

	result, err := s.querier.CreateAuthor(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateAuthor")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateAuthorResponse{Value: toExecResult(result)}), nil
}

func (s *Service) DeleteAuthor(ctx context.Context, req *connect.Request[pb.DeleteAuthorRequest]) (*connect.Response[pb.DeleteAuthorResponse], error) {
	id := req.Msg.GetId()

	err := s.querier.DeleteAuthor(ctx, id)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "DeleteAuthor")
		return nil, err
	}
	return connect.NewResponse(&pb.DeleteAuthorResponse{}), nil
}

func (s *Service) GetAuthor(ctx context.Context, req *connect.Request[pb.GetAuthorRequest]) (*connect.Response[pb.GetAuthorResponse], error) {
	id := req.Msg.GetId()

	result, err := s.querier.GetAuthor(ctx, id)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "GetAuthor")
		return nil, err
	}
	return connect.NewResponse(&pb.GetAuthorResponse{Author: toAuthor(result)}), nil
}

func (s *Service) ListAuthors(ctx context.Context, req *connect.Request[pb.ListAuthorsRequest]) (*connect.Response[pb.ListAuthorsResponse], error) {

	result, err := s.querier.ListAuthors(ctx)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "ListAuthors")
		return nil, err
	}
	res := new(pb.ListAuthorsResponse)
	for _, r := range result {
		res.List = append(res.List, toAuthor(r))
	}
	return connect.NewResponse(res), nil
}

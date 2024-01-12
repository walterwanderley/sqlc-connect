// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package authors

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, db DBTX, arg *CreateAuthorParams) (*Authors, error)
	DeleteAuthor(ctx context.Context, db DBTX, id int64) error
	GetAuthor(ctx context.Context, db DBTX, id int64) (*Authors, error)
	ListAuthors(ctx context.Context, db DBTX) ([]*Authors, error)
}

var _ Querier = (*Queries)(nil)
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package authors

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Authors struct {
	ID   int64       `json:"id"`
	Name string      `json:"name"`
	Bio  pgtype.Text `json:"bio"`
}

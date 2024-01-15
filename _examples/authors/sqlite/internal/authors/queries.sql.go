// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package authors

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: createAuthor :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ? 
)
`

type createAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) createAuthor(ctx context.Context, arg createAuthorParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAuthor, arg.Name, arg.Bio)
}

const deleteAuthor = `-- name: deleteAuthor :exec
DELETE FROM authors
WHERE id = ?
`

func (q *Queries) deleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: getAuthor :one
SELECT id, name, bio FROM authors
WHERE id = ? LIMIT 1
`

func (q *Queries) getAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthors = `-- name: listAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) listAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

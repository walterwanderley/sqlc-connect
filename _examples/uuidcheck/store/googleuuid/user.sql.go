// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package googleuuid

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createLocationTransactions = `-- name: CreateLocationTransactions :exec

INSERT INTO location_transactions
SELECT
    UNNEST($1::UUID[]) as location_id,
    UNNEST($2::UUID[]) as transaction_id
`

type CreateLocationTransactionsParams struct {
	Column1 []uuid.UUID `json:"column_1"`
	Column2 []uuid.UUID `json:"column_2"`
}

// ---------
func (q *Queries) CreateLocationTransactions(ctx context.Context, arg CreateLocationTransactionsParams) error {
	_, err := q.db.Exec(ctx, createLocationTransactions, arg.Column1, arg.Column2)
	return err
}

const createProduct = `-- name: CreateProduct :one

INSERT INTO products (id, category) VALUES ($1, $2) RETURNING id
`

type CreateProductParams struct {
	ID       int32       `json:"id"`
	Category pgtype.Int4 `json:"category"`
}

// ---------
func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (int32, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.ID, arg.Category)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createProductReturnAll = `-- name: CreateProductReturnAll :one
INSERT INTO products (id, category) VALUES ($1, $2) RETURNING id, category, name
`

type CreateProductReturnAllParams struct {
	ID       int32       `json:"id"`
	Category pgtype.Int4 `json:"category"`
}

func (q *Queries) CreateProductReturnAll(ctx context.Context, arg CreateProductReturnAllParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProductReturnAll, arg.ID, arg.Category)
	var i Product
	err := row.Scan(&i.ID, &i.Category, &i.Name)
	return i, err
}

const createProductReturnPartial = `-- name: CreateProductReturnPartial :one
INSERT INTO
    products (id, category)
VALUES ($1, $2) RETURNING id,
    name
`

type CreateProductReturnPartialParams struct {
	ID       int32       `json:"id"`
	Category pgtype.Int4 `json:"category"`
}

type CreateProductReturnPartialRow struct {
	ID   int32       `json:"id"`
	Name pgtype.Text `json:"name"`
}

func (q *Queries) CreateProductReturnPartial(ctx context.Context, arg CreateProductReturnPartialParams) (CreateProductReturnPartialRow, error) {
	row := q.db.QueryRow(ctx, createProductReturnPartial, arg.ID, arg.Category)
	var i CreateProductReturnPartialRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, location) VALUES ($1, $2) RETURNING id
`

type CreateUserParams struct {
	ID       uuid.UUID   `json:"id"`
	Location pgtype.UUID `json:"location"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createUser, arg.ID, arg.Location)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createUserReturnAll = `-- name: CreateUserReturnAll :one
INSERT INTO users (id, location) VALUES ($1, $2) RETURNING id, location, name
`

type CreateUserReturnAllParams struct {
	ID       uuid.UUID   `json:"id"`
	Location pgtype.UUID `json:"location"`
}

func (q *Queries) CreateUserReturnAll(ctx context.Context, arg CreateUserReturnAllParams) (User, error) {
	row := q.db.QueryRow(ctx, createUserReturnAll, arg.ID, arg.Location)
	var i User
	err := row.Scan(&i.ID, &i.Location, &i.Name)
	return i, err
}

const createUserReturnPartial = `-- name: CreateUserReturnPartial :one
INSERT INTO users (id, location) VALUES ($1, $2) RETURNING id, name
`

type CreateUserReturnPartialParams struct {
	ID       uuid.UUID   `json:"id"`
	Location pgtype.UUID `json:"location"`
}

type CreateUserReturnPartialRow struct {
	ID   uuid.UUID   `json:"id"`
	Name pgtype.Text `json:"name"`
}

func (q *Queries) CreateUserReturnPartial(ctx context.Context, arg CreateUserReturnPartialParams) (CreateUserReturnPartialRow, error) {
	row := q.db.QueryRow(ctx, createUserReturnPartial, arg.ID, arg.Location)
	var i CreateUserReturnPartialRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getProductsByIds = `-- name: GetProductsByIds :many
SELECT id, category, name FROM products WHERE id = ANY($1::uuid[])
`

func (q *Queries) GetProductsByIds(ctx context.Context, dollar_1 []uuid.UUID) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProductsByIds, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Category, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

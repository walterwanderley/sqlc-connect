// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect).

package authors

import (
	"authors/api/authors/v1/v1connect"
)

// NewService is a constructor of a v1.AuthorsServiceHandler implementation.
// Use this function to customize the server by adding middlewares to it.
func NewService(querier *Queries) v1connect.AuthorsServiceHandler {
	return &Service{querier: querier}
}

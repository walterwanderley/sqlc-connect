[![LiberaPay](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/walterwanderley/donate)
[![receives](https://img.shields.io/liberapay/receives/walterwanderley.svg?logo=liberapay)](https://liberapay.com/walterwanderley/donate)
[![patrons](https://img.shields.io/liberapay/patrons/walterwanderley.svg?logo=liberapay)](https://liberapay.com/walterwanderley/donate)

# sqlc-connect
Generate [connect-go](https://connect.build/) server from SQL. If you’re searching for a SQLC plugin, use [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server/).

### Requirements

- Go 1.21 or superior
- [sqlc](https://sqlc.dev/)
- [buf](https://buf.build/)

```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/bufbuild/buf/cmd/buf@latest
```

### Installation

```sh
go install github.com/walterwanderley/sqlc-connect@latest
```

### Example

1. Create a queries.sql file:

```sql
--queries.sql

CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text,
  created_at TIMESTAMP
);

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio, created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

```

2. Create a sqlc.yaml file

```yaml
version: "2"
sql:
- schema: "./queries.sql"
  queries: "./queries.sql"
  engine: "postgresql"
  gen:
    go:
      out: "internal/author"
      
```

3. Execute sqlc

```sh
sqlc generate
```

4. Execute sqlc-connect

```sh
sqlc-connect -m "authors"
```

5. Run the generated server

```sh
go run . -db [Database Connection URL] -dev
```

6. Enjoy!

[http://localhost:5000/swagger](http://localhost:5000/swagger)

or

```sh
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
grpcui -plaintext localhost:5000
```

### Editing the generated code

- It's safe to edit any generated code that doesn't have the `DO NOT EDIT` indication at the very first line.

- After modify a SQL file, execute these commands below:

```sh
sqlc generate
go generate
```

- After modify a *.proto file, execute `buf generate`.

### Similar Projects

- [sqlc-grpc](https://github.com/walterwanderley/sqlc-grpc)
- [xo-grpc](https://github.com/walterwanderley/xo-grpc)
# About

Author example taken from [sqlc][sqlc] Git repository [examples][sqlc-git].

[sqlc]: https://sqlc.dev
[sqlc-git]: https://github.com/sqlc-dev/sqlc/tree/main/examples/authors

## Running

```sh
./gen.sh
docker compose up
```

### Exploring with gRPC UI

```sh
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
grpcui -plaintext localhost:8080
```

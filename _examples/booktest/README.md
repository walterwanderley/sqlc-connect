# About

Booktest example taken from [sqlc][sqlc] Git repository [examples][sqlc-git].

[sqlc]: https://sqlc.dev
[sqlc-git]: https://github.com/kyleconroy/sqlc/tree/main/examples/booktest

## Running

```sh
./gen.sh
docker-compose up
```

### Exploring with gRPC UI

```sh
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
grpcui -plaintext localhost:8080
```


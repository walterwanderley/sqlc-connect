module github.com/walterwanderley/sqlc-connect

go 1.23

replace github.com/walterwanderley/sqlc-grpc => ../sqlc-grpc

require (
	github.com/walterwanderley/sqlc-grpc v0.19.9
	golang.org/x/mod v0.20.0
	golang.org/x/tools v0.24.0
)

require (
	github.com/emicklei/proto v1.13.2 // indirect
	golang.org/x/sync v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

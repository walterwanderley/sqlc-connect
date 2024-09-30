// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: googleuuid/v1/googleuuid.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
	v1 "uuidcheck/api/googleuuid/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GoogleuuidServiceName is the fully-qualified name of the GoogleuuidService service.
	GoogleuuidServiceName = "googleuuid.v1.GoogleuuidService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GoogleuuidServiceGetProductsByIdsProcedure is the fully-qualified name of the GoogleuuidService's
	// GetProductsByIds RPC.
	GoogleuuidServiceGetProductsByIdsProcedure = "/googleuuid.v1.GoogleuuidService/GetProductsByIds"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	googleuuidServiceServiceDescriptor                = v1.File_googleuuid_v1_googleuuid_proto.Services().ByName("GoogleuuidService")
	googleuuidServiceGetProductsByIdsMethodDescriptor = googleuuidServiceServiceDescriptor.Methods().ByName("GetProductsByIds")
)

// GoogleuuidServiceClient is a client for the googleuuid.v1.GoogleuuidService service.
type GoogleuuidServiceClient interface {
	GetProductsByIds(context.Context, *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error)
}

// NewGoogleuuidServiceClient constructs a client for the googleuuid.v1.GoogleuuidService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGoogleuuidServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GoogleuuidServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &googleuuidServiceClient{
		getProductsByIds: connect.NewClient[v1.GetProductsByIdsRequest, v1.GetProductsByIdsResponse](
			httpClient,
			baseURL+GoogleuuidServiceGetProductsByIdsProcedure,
			connect.WithSchema(googleuuidServiceGetProductsByIdsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// googleuuidServiceClient implements GoogleuuidServiceClient.
type googleuuidServiceClient struct {
	getProductsByIds *connect.Client[v1.GetProductsByIdsRequest, v1.GetProductsByIdsResponse]
}

// GetProductsByIds calls googleuuid.v1.GoogleuuidService.GetProductsByIds.
func (c *googleuuidServiceClient) GetProductsByIds(ctx context.Context, req *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error) {
	return c.getProductsByIds.CallUnary(ctx, req)
}

// GoogleuuidServiceHandler is an implementation of the googleuuid.v1.GoogleuuidService service.
type GoogleuuidServiceHandler interface {
	GetProductsByIds(context.Context, *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error)
}

// NewGoogleuuidServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGoogleuuidServiceHandler(svc GoogleuuidServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	googleuuidServiceGetProductsByIdsHandler := connect.NewUnaryHandler(
		GoogleuuidServiceGetProductsByIdsProcedure,
		svc.GetProductsByIds,
		connect.WithSchema(googleuuidServiceGetProductsByIdsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/googleuuid.v1.GoogleuuidService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GoogleuuidServiceGetProductsByIdsProcedure:
			googleuuidServiceGetProductsByIdsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGoogleuuidServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGoogleuuidServiceHandler struct{}

func (UnimplementedGoogleuuidServiceHandler) GetProductsByIds(context.Context, *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.GetProductsByIds is not implemented"))
}

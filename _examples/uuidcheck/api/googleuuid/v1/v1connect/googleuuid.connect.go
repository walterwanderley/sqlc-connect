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
	// GoogleuuidServiceCreateLocationTransactionsProcedure is the fully-qualified name of the
	// GoogleuuidService's CreateLocationTransactions RPC.
	GoogleuuidServiceCreateLocationTransactionsProcedure = "/googleuuid.v1.GoogleuuidService/CreateLocationTransactions"
	// GoogleuuidServiceCreateProductProcedure is the fully-qualified name of the GoogleuuidService's
	// CreateProduct RPC.
	GoogleuuidServiceCreateProductProcedure = "/googleuuid.v1.GoogleuuidService/CreateProduct"
	// GoogleuuidServiceCreateProductReturnAllProcedure is the fully-qualified name of the
	// GoogleuuidService's CreateProductReturnAll RPC.
	GoogleuuidServiceCreateProductReturnAllProcedure = "/googleuuid.v1.GoogleuuidService/CreateProductReturnAll"
	// GoogleuuidServiceCreateProductReturnPartialProcedure is the fully-qualified name of the
	// GoogleuuidService's CreateProductReturnPartial RPC.
	GoogleuuidServiceCreateProductReturnPartialProcedure = "/googleuuid.v1.GoogleuuidService/CreateProductReturnPartial"
	// GoogleuuidServiceCreateUserProcedure is the fully-qualified name of the GoogleuuidService's
	// CreateUser RPC.
	GoogleuuidServiceCreateUserProcedure = "/googleuuid.v1.GoogleuuidService/CreateUser"
	// GoogleuuidServiceCreateUserReturnAllProcedure is the fully-qualified name of the
	// GoogleuuidService's CreateUserReturnAll RPC.
	GoogleuuidServiceCreateUserReturnAllProcedure = "/googleuuid.v1.GoogleuuidService/CreateUserReturnAll"
	// GoogleuuidServiceCreateUserReturnPartialProcedure is the fully-qualified name of the
	// GoogleuuidService's CreateUserReturnPartial RPC.
	GoogleuuidServiceCreateUserReturnPartialProcedure = "/googleuuid.v1.GoogleuuidService/CreateUserReturnPartial"
	// GoogleuuidServiceGetProductsByIdsProcedure is the fully-qualified name of the GoogleuuidService's
	// GetProductsByIds RPC.
	GoogleuuidServiceGetProductsByIdsProcedure = "/googleuuid.v1.GoogleuuidService/GetProductsByIds"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	googleuuidServiceServiceDescriptor                          = v1.File_googleuuid_v1_googleuuid_proto.Services().ByName("GoogleuuidService")
	googleuuidServiceCreateLocationTransactionsMethodDescriptor = googleuuidServiceServiceDescriptor.Methods().ByName("CreateLocationTransactions")
	googleuuidServiceCreateProductMethodDescriptor              = googleuuidServiceServiceDescriptor.Methods().ByName("CreateProduct")
	googleuuidServiceCreateProductReturnAllMethodDescriptor     = googleuuidServiceServiceDescriptor.Methods().ByName("CreateProductReturnAll")
	googleuuidServiceCreateProductReturnPartialMethodDescriptor = googleuuidServiceServiceDescriptor.Methods().ByName("CreateProductReturnPartial")
	googleuuidServiceCreateUserMethodDescriptor                 = googleuuidServiceServiceDescriptor.Methods().ByName("CreateUser")
	googleuuidServiceCreateUserReturnAllMethodDescriptor        = googleuuidServiceServiceDescriptor.Methods().ByName("CreateUserReturnAll")
	googleuuidServiceCreateUserReturnPartialMethodDescriptor    = googleuuidServiceServiceDescriptor.Methods().ByName("CreateUserReturnPartial")
	googleuuidServiceGetProductsByIdsMethodDescriptor           = googleuuidServiceServiceDescriptor.Methods().ByName("GetProductsByIds")
)

// GoogleuuidServiceClient is a client for the googleuuid.v1.GoogleuuidService service.
type GoogleuuidServiceClient interface {
	CreateLocationTransactions(context.Context, *connect.Request[v1.CreateLocationTransactionsRequest]) (*connect.Response[v1.CreateLocationTransactionsResponse], error)
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error)
	CreateProductReturnAll(context.Context, *connect.Request[v1.CreateProductReturnAllRequest]) (*connect.Response[v1.CreateProductReturnAllResponse], error)
	CreateProductReturnPartial(context.Context, *connect.Request[v1.CreateProductReturnPartialRequest]) (*connect.Response[v1.CreateProductReturnPartialResponse], error)
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error)
	CreateUserReturnAll(context.Context, *connect.Request[v1.CreateUserReturnAllRequest]) (*connect.Response[v1.CreateUserReturnAllResponse], error)
	CreateUserReturnPartial(context.Context, *connect.Request[v1.CreateUserReturnPartialRequest]) (*connect.Response[v1.CreateUserReturnPartialResponse], error)
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
		createLocationTransactions: connect.NewClient[v1.CreateLocationTransactionsRequest, v1.CreateLocationTransactionsResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateLocationTransactionsProcedure,
			connect.WithSchema(googleuuidServiceCreateLocationTransactionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProduct: connect.NewClient[v1.CreateProductRequest, v1.CreateProductResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateProductProcedure,
			connect.WithSchema(googleuuidServiceCreateProductMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProductReturnAll: connect.NewClient[v1.CreateProductReturnAllRequest, v1.CreateProductReturnAllResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateProductReturnAllProcedure,
			connect.WithSchema(googleuuidServiceCreateProductReturnAllMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProductReturnPartial: connect.NewClient[v1.CreateProductReturnPartialRequest, v1.CreateProductReturnPartialResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateProductReturnPartialProcedure,
			connect.WithSchema(googleuuidServiceCreateProductReturnPartialMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createUser: connect.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateUserProcedure,
			connect.WithSchema(googleuuidServiceCreateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createUserReturnAll: connect.NewClient[v1.CreateUserReturnAllRequest, v1.CreateUserReturnAllResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateUserReturnAllProcedure,
			connect.WithSchema(googleuuidServiceCreateUserReturnAllMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createUserReturnPartial: connect.NewClient[v1.CreateUserReturnPartialRequest, v1.CreateUserReturnPartialResponse](
			httpClient,
			baseURL+GoogleuuidServiceCreateUserReturnPartialProcedure,
			connect.WithSchema(googleuuidServiceCreateUserReturnPartialMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
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
	createLocationTransactions *connect.Client[v1.CreateLocationTransactionsRequest, v1.CreateLocationTransactionsResponse]
	createProduct              *connect.Client[v1.CreateProductRequest, v1.CreateProductResponse]
	createProductReturnAll     *connect.Client[v1.CreateProductReturnAllRequest, v1.CreateProductReturnAllResponse]
	createProductReturnPartial *connect.Client[v1.CreateProductReturnPartialRequest, v1.CreateProductReturnPartialResponse]
	createUser                 *connect.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	createUserReturnAll        *connect.Client[v1.CreateUserReturnAllRequest, v1.CreateUserReturnAllResponse]
	createUserReturnPartial    *connect.Client[v1.CreateUserReturnPartialRequest, v1.CreateUserReturnPartialResponse]
	getProductsByIds           *connect.Client[v1.GetProductsByIdsRequest, v1.GetProductsByIdsResponse]
}

// CreateLocationTransactions calls googleuuid.v1.GoogleuuidService.CreateLocationTransactions.
func (c *googleuuidServiceClient) CreateLocationTransactions(ctx context.Context, req *connect.Request[v1.CreateLocationTransactionsRequest]) (*connect.Response[v1.CreateLocationTransactionsResponse], error) {
	return c.createLocationTransactions.CallUnary(ctx, req)
}

// CreateProduct calls googleuuid.v1.GoogleuuidService.CreateProduct.
func (c *googleuuidServiceClient) CreateProduct(ctx context.Context, req *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error) {
	return c.createProduct.CallUnary(ctx, req)
}

// CreateProductReturnAll calls googleuuid.v1.GoogleuuidService.CreateProductReturnAll.
func (c *googleuuidServiceClient) CreateProductReturnAll(ctx context.Context, req *connect.Request[v1.CreateProductReturnAllRequest]) (*connect.Response[v1.CreateProductReturnAllResponse], error) {
	return c.createProductReturnAll.CallUnary(ctx, req)
}

// CreateProductReturnPartial calls googleuuid.v1.GoogleuuidService.CreateProductReturnPartial.
func (c *googleuuidServiceClient) CreateProductReturnPartial(ctx context.Context, req *connect.Request[v1.CreateProductReturnPartialRequest]) (*connect.Response[v1.CreateProductReturnPartialResponse], error) {
	return c.createProductReturnPartial.CallUnary(ctx, req)
}

// CreateUser calls googleuuid.v1.GoogleuuidService.CreateUser.
func (c *googleuuidServiceClient) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// CreateUserReturnAll calls googleuuid.v1.GoogleuuidService.CreateUserReturnAll.
func (c *googleuuidServiceClient) CreateUserReturnAll(ctx context.Context, req *connect.Request[v1.CreateUserReturnAllRequest]) (*connect.Response[v1.CreateUserReturnAllResponse], error) {
	return c.createUserReturnAll.CallUnary(ctx, req)
}

// CreateUserReturnPartial calls googleuuid.v1.GoogleuuidService.CreateUserReturnPartial.
func (c *googleuuidServiceClient) CreateUserReturnPartial(ctx context.Context, req *connect.Request[v1.CreateUserReturnPartialRequest]) (*connect.Response[v1.CreateUserReturnPartialResponse], error) {
	return c.createUserReturnPartial.CallUnary(ctx, req)
}

// GetProductsByIds calls googleuuid.v1.GoogleuuidService.GetProductsByIds.
func (c *googleuuidServiceClient) GetProductsByIds(ctx context.Context, req *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error) {
	return c.getProductsByIds.CallUnary(ctx, req)
}

// GoogleuuidServiceHandler is an implementation of the googleuuid.v1.GoogleuuidService service.
type GoogleuuidServiceHandler interface {
	CreateLocationTransactions(context.Context, *connect.Request[v1.CreateLocationTransactionsRequest]) (*connect.Response[v1.CreateLocationTransactionsResponse], error)
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error)
	CreateProductReturnAll(context.Context, *connect.Request[v1.CreateProductReturnAllRequest]) (*connect.Response[v1.CreateProductReturnAllResponse], error)
	CreateProductReturnPartial(context.Context, *connect.Request[v1.CreateProductReturnPartialRequest]) (*connect.Response[v1.CreateProductReturnPartialResponse], error)
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error)
	CreateUserReturnAll(context.Context, *connect.Request[v1.CreateUserReturnAllRequest]) (*connect.Response[v1.CreateUserReturnAllResponse], error)
	CreateUserReturnPartial(context.Context, *connect.Request[v1.CreateUserReturnPartialRequest]) (*connect.Response[v1.CreateUserReturnPartialResponse], error)
	GetProductsByIds(context.Context, *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error)
}

// NewGoogleuuidServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGoogleuuidServiceHandler(svc GoogleuuidServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	googleuuidServiceCreateLocationTransactionsHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateLocationTransactionsProcedure,
		svc.CreateLocationTransactions,
		connect.WithSchema(googleuuidServiceCreateLocationTransactionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceCreateProductHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateProductProcedure,
		svc.CreateProduct,
		connect.WithSchema(googleuuidServiceCreateProductMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceCreateProductReturnAllHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateProductReturnAllProcedure,
		svc.CreateProductReturnAll,
		connect.WithSchema(googleuuidServiceCreateProductReturnAllMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceCreateProductReturnPartialHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateProductReturnPartialProcedure,
		svc.CreateProductReturnPartial,
		connect.WithSchema(googleuuidServiceCreateProductReturnPartialMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceCreateUserHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateUserProcedure,
		svc.CreateUser,
		connect.WithSchema(googleuuidServiceCreateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceCreateUserReturnAllHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateUserReturnAllProcedure,
		svc.CreateUserReturnAll,
		connect.WithSchema(googleuuidServiceCreateUserReturnAllMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceCreateUserReturnPartialHandler := connect.NewUnaryHandler(
		GoogleuuidServiceCreateUserReturnPartialProcedure,
		svc.CreateUserReturnPartial,
		connect.WithSchema(googleuuidServiceCreateUserReturnPartialMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	googleuuidServiceGetProductsByIdsHandler := connect.NewUnaryHandler(
		GoogleuuidServiceGetProductsByIdsProcedure,
		svc.GetProductsByIds,
		connect.WithSchema(googleuuidServiceGetProductsByIdsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/googleuuid.v1.GoogleuuidService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GoogleuuidServiceCreateLocationTransactionsProcedure:
			googleuuidServiceCreateLocationTransactionsHandler.ServeHTTP(w, r)
		case GoogleuuidServiceCreateProductProcedure:
			googleuuidServiceCreateProductHandler.ServeHTTP(w, r)
		case GoogleuuidServiceCreateProductReturnAllProcedure:
			googleuuidServiceCreateProductReturnAllHandler.ServeHTTP(w, r)
		case GoogleuuidServiceCreateProductReturnPartialProcedure:
			googleuuidServiceCreateProductReturnPartialHandler.ServeHTTP(w, r)
		case GoogleuuidServiceCreateUserProcedure:
			googleuuidServiceCreateUserHandler.ServeHTTP(w, r)
		case GoogleuuidServiceCreateUserReturnAllProcedure:
			googleuuidServiceCreateUserReturnAllHandler.ServeHTTP(w, r)
		case GoogleuuidServiceCreateUserReturnPartialProcedure:
			googleuuidServiceCreateUserReturnPartialHandler.ServeHTTP(w, r)
		case GoogleuuidServiceGetProductsByIdsProcedure:
			googleuuidServiceGetProductsByIdsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGoogleuuidServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGoogleuuidServiceHandler struct{}

func (UnimplementedGoogleuuidServiceHandler) CreateLocationTransactions(context.Context, *connect.Request[v1.CreateLocationTransactionsRequest]) (*connect.Response[v1.CreateLocationTransactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateLocationTransactions is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateProduct is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) CreateProductReturnAll(context.Context, *connect.Request[v1.CreateProductReturnAllRequest]) (*connect.Response[v1.CreateProductReturnAllResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateProductReturnAll is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) CreateProductReturnPartial(context.Context, *connect.Request[v1.CreateProductReturnPartialRequest]) (*connect.Response[v1.CreateProductReturnPartialResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateProductReturnPartial is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateUser is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) CreateUserReturnAll(context.Context, *connect.Request[v1.CreateUserReturnAllRequest]) (*connect.Response[v1.CreateUserReturnAllResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateUserReturnAll is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) CreateUserReturnPartial(context.Context, *connect.Request[v1.CreateUserReturnPartialRequest]) (*connect.Response[v1.CreateUserReturnPartialResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.CreateUserReturnPartial is not implemented"))
}

func (UnimplementedGoogleuuidServiceHandler) GetProductsByIds(context.Context, *connect.Request[v1.GetProductsByIdsRequest]) (*connect.Response[v1.GetProductsByIdsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("googleuuid.v1.GoogleuuidService.GetProductsByIds is not implemented"))
}

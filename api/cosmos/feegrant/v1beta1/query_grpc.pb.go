// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: cosmos/feegrant/v1beta1/query.proto

package feegrantv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Query_Allowance_FullMethodName           = "/cosmos.feegrant.v1beta1.Query/Allowance"
	Query_Allowances_FullMethodName          = "/cosmos.feegrant.v1beta1.Query/Allowances"
	Query_AllowancesByGranter_FullMethodName = "/cosmos.feegrant.v1beta1.Query/AllowancesByGranter"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service.
type QueryClient interface {
	// Allowance returns fee granted to the grantee by the granter.
	Allowance(ctx context.Context, in *QueryAllowanceRequest, opts ...grpc.CallOption) (*QueryAllowanceResponse, error)
	// Allowances returns all the grants for address.
	Allowances(ctx context.Context, in *QueryAllowancesRequest, opts ...grpc.CallOption) (*QueryAllowancesResponse, error)
	// AllowancesByGranter returns all the grants given by an address
	//
	// Since: cosmos-sdk 0.46
	AllowancesByGranter(ctx context.Context, in *QueryAllowancesByGranterRequest, opts ...grpc.CallOption) (*QueryAllowancesByGranterResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Allowance(ctx context.Context, in *QueryAllowanceRequest, opts ...grpc.CallOption) (*QueryAllowanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAllowanceResponse)
	err := c.cc.Invoke(ctx, Query_Allowance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Allowances(ctx context.Context, in *QueryAllowancesRequest, opts ...grpc.CallOption) (*QueryAllowancesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAllowancesResponse)
	err := c.cc.Invoke(ctx, Query_Allowances_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllowancesByGranter(ctx context.Context, in *QueryAllowancesByGranterRequest, opts ...grpc.CallOption) (*QueryAllowancesByGranterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAllowancesByGranterResponse)
	err := c.cc.Invoke(ctx, Query_AllowancesByGranter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
//
// Query defines the gRPC querier service.
type QueryServer interface {
	// Allowance returns fee granted to the grantee by the granter.
	Allowance(context.Context, *QueryAllowanceRequest) (*QueryAllowanceResponse, error)
	// Allowances returns all the grants for address.
	Allowances(context.Context, *QueryAllowancesRequest) (*QueryAllowancesResponse, error)
	// AllowancesByGranter returns all the grants given by an address
	//
	// Since: cosmos-sdk 0.46
	AllowancesByGranter(context.Context, *QueryAllowancesByGranterRequest) (*QueryAllowancesByGranterResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Allowance(context.Context, *QueryAllowanceRequest) (*QueryAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allowance not implemented")
}
func (UnimplementedQueryServer) Allowances(context.Context, *QueryAllowancesRequest) (*QueryAllowancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allowances not implemented")
}
func (UnimplementedQueryServer) AllowancesByGranter(context.Context, *QueryAllowancesByGranterRequest) (*QueryAllowancesByGranterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowancesByGranter not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Allowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Allowance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allowance(ctx, req.(*QueryAllowanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Allowances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allowances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Allowances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allowances(ctx, req.(*QueryAllowancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllowancesByGranter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowancesByGranterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllowancesByGranter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllowancesByGranter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllowancesByGranter(ctx, req.(*QueryAllowancesByGranterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.feegrant.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allowance",
			Handler:    _Query_Allowance_Handler,
		},
		{
			MethodName: "Allowances",
			Handler:    _Query_Allowances_Handler,
		},
		{
			MethodName: "AllowancesByGranter",
			Handler:    _Query_AllowancesByGranter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/feegrant/v1beta1/query.proto",
}

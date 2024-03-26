// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/discovered_cluster_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiscoveredClustersService_CountDiscoveredClusters_FullMethodName = "/v1.DiscoveredClustersService/CountDiscoveredClusters"
	DiscoveredClustersService_GetDiscoveredCluster_FullMethodName    = "/v1.DiscoveredClustersService/GetDiscoveredCluster"
	DiscoveredClustersService_ListDiscoveredClusters_FullMethodName  = "/v1.DiscoveredClustersService/ListDiscoveredClusters"
)

// DiscoveredClustersServiceClient is the client API for DiscoveredClustersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveredClustersServiceClient interface {
	// CountDiscoveredClusters returns the number of discovered clusters after filtering by requested fields.
	CountDiscoveredClusters(ctx context.Context, in *CountDiscoveredClustersRequest, opts ...grpc.CallOption) (*CountDiscoveredClustersResponse, error)
	// GetDiscoveredCluster retrieves a discovered cluster by ID.
	GetDiscoveredCluster(ctx context.Context, in *GetDiscoveredClusterRequest, opts ...grpc.CallOption) (*GetDiscoveredClusterResponse, error)
	// ListDiscoveredClusters returns the list of discovered clusters after filtered by requested fields.
	ListDiscoveredClusters(ctx context.Context, in *ListDiscoveredClustersRequest, opts ...grpc.CallOption) (*ListDiscoveredClustersResponse, error)
}

type discoveredClustersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveredClustersServiceClient(cc grpc.ClientConnInterface) DiscoveredClustersServiceClient {
	return &discoveredClustersServiceClient{cc}
}

func (c *discoveredClustersServiceClient) CountDiscoveredClusters(ctx context.Context, in *CountDiscoveredClustersRequest, opts ...grpc.CallOption) (*CountDiscoveredClustersResponse, error) {
	out := new(CountDiscoveredClustersResponse)
	err := c.cc.Invoke(ctx, DiscoveredClustersService_CountDiscoveredClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveredClustersServiceClient) GetDiscoveredCluster(ctx context.Context, in *GetDiscoveredClusterRequest, opts ...grpc.CallOption) (*GetDiscoveredClusterResponse, error) {
	out := new(GetDiscoveredClusterResponse)
	err := c.cc.Invoke(ctx, DiscoveredClustersService_GetDiscoveredCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveredClustersServiceClient) ListDiscoveredClusters(ctx context.Context, in *ListDiscoveredClustersRequest, opts ...grpc.CallOption) (*ListDiscoveredClustersResponse, error) {
	out := new(ListDiscoveredClustersResponse)
	err := c.cc.Invoke(ctx, DiscoveredClustersService_ListDiscoveredClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveredClustersServiceServer is the server API for DiscoveredClustersService service.
// All implementations should embed UnimplementedDiscoveredClustersServiceServer
// for forward compatibility
type DiscoveredClustersServiceServer interface {
	// CountDiscoveredClusters returns the number of discovered clusters after filtering by requested fields.
	CountDiscoveredClusters(context.Context, *CountDiscoveredClustersRequest) (*CountDiscoveredClustersResponse, error)
	// GetDiscoveredCluster retrieves a discovered cluster by ID.
	GetDiscoveredCluster(context.Context, *GetDiscoveredClusterRequest) (*GetDiscoveredClusterResponse, error)
	// ListDiscoveredClusters returns the list of discovered clusters after filtered by requested fields.
	ListDiscoveredClusters(context.Context, *ListDiscoveredClustersRequest) (*ListDiscoveredClustersResponse, error)
}

// UnimplementedDiscoveredClustersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDiscoveredClustersServiceServer struct {
}

func (UnimplementedDiscoveredClustersServiceServer) CountDiscoveredClusters(context.Context, *CountDiscoveredClustersRequest) (*CountDiscoveredClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountDiscoveredClusters not implemented")
}
func (UnimplementedDiscoveredClustersServiceServer) GetDiscoveredCluster(context.Context, *GetDiscoveredClusterRequest) (*GetDiscoveredClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscoveredCluster not implemented")
}
func (UnimplementedDiscoveredClustersServiceServer) ListDiscoveredClusters(context.Context, *ListDiscoveredClustersRequest) (*ListDiscoveredClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiscoveredClusters not implemented")
}

// UnsafeDiscoveredClustersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveredClustersServiceServer will
// result in compilation errors.
type UnsafeDiscoveredClustersServiceServer interface {
	mustEmbedUnimplementedDiscoveredClustersServiceServer()
}

func RegisterDiscoveredClustersServiceServer(s grpc.ServiceRegistrar, srv DiscoveredClustersServiceServer) {
	s.RegisterService(&DiscoveredClustersService_ServiceDesc, srv)
}

func _DiscoveredClustersService_CountDiscoveredClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountDiscoveredClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveredClustersServiceServer).CountDiscoveredClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveredClustersService_CountDiscoveredClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveredClustersServiceServer).CountDiscoveredClusters(ctx, req.(*CountDiscoveredClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveredClustersService_GetDiscoveredCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscoveredClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveredClustersServiceServer).GetDiscoveredCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveredClustersService_GetDiscoveredCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveredClustersServiceServer).GetDiscoveredCluster(ctx, req.(*GetDiscoveredClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveredClustersService_ListDiscoveredClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiscoveredClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveredClustersServiceServer).ListDiscoveredClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveredClustersService_ListDiscoveredClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveredClustersServiceServer).ListDiscoveredClusters(ctx, req.(*ListDiscoveredClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscoveredClustersService_ServiceDesc is the grpc.ServiceDesc for DiscoveredClustersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveredClustersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DiscoveredClustersService",
	HandlerType: (*DiscoveredClustersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountDiscoveredClusters",
			Handler:    _DiscoveredClustersService_CountDiscoveredClusters_Handler,
		},
		{
			MethodName: "GetDiscoveredCluster",
			Handler:    _DiscoveredClustersService_GetDiscoveredCluster_Handler,
		},
		{
			MethodName: "ListDiscoveredClusters",
			Handler:    _DiscoveredClustersService_ListDiscoveredClusters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/discovered_cluster_service.proto",
}
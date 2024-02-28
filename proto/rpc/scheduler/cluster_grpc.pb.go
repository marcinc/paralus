// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/rpc/scheduler/cluster.proto

package rpcv3

import (
	context "context"
	v31 "github.com/paralus/paralus/proto/types/commonpb/v3"
	v3 "github.com/paralus/paralus/proto/types/infrapb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClusterService_CreateCluster_FullMethodName       = "/paralus.dev.rpc.v3.ClusterService/CreateCluster"
	ClusterService_GetClusters_FullMethodName         = "/paralus.dev.rpc.v3.ClusterService/GetClusters"
	ClusterService_GetCluster_FullMethodName          = "/paralus.dev.rpc.v3.ClusterService/GetCluster"
	ClusterService_UpdateCluster_FullMethodName       = "/paralus.dev.rpc.v3.ClusterService/UpdateCluster"
	ClusterService_DeleteCluster_FullMethodName       = "/paralus.dev.rpc.v3.ClusterService/DeleteCluster"
	ClusterService_DownloadCluster_FullMethodName     = "/paralus.dev.rpc.v3.ClusterService/DownloadCluster"
	ClusterService_UpdateClusterStatus_FullMethodName = "/paralus.dev.rpc.v3.ClusterService/UpdateClusterStatus"
	ClusterService_GetClusterStatus_FullMethodName    = "/paralus.dev.rpc.v3.ClusterService/GetClusterStatus"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	CreateCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v3.Cluster, error)
	GetClusters(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.ClusterList, error)
	GetCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v3.Cluster, error)
	UpdateCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v3.Cluster, error)
	DeleteCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*DeleteClusterResponse, error)
	DownloadCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v31.HttpBody, error)
	UpdateClusterStatus(ctx context.Context, in *UpdateClusterStatusRequest, opts ...grpc.CallOption) (*UpdateClusterStatusResponse, error)
	GetClusterStatus(ctx context.Context, in *GetClusterStatusRequest, opts ...grpc.CallOption) (*GetClusterStatusResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) CreateCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v3.Cluster, error) {
	out := new(v3.Cluster)
	err := c.cc.Invoke(ctx, ClusterService_CreateCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetClusters(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.ClusterList, error) {
	out := new(v3.ClusterList)
	err := c.cc.Invoke(ctx, ClusterService_GetClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v3.Cluster, error) {
	out := new(v3.Cluster)
	err := c.cc.Invoke(ctx, ClusterService_GetCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v3.Cluster, error) {
	out := new(v3.Cluster)
	err := c.cc.Invoke(ctx, ClusterService_UpdateCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DeleteCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*DeleteClusterResponse, error) {
	out := new(DeleteClusterResponse)
	err := c.cc.Invoke(ctx, ClusterService_DeleteCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DownloadCluster(ctx context.Context, in *v3.Cluster, opts ...grpc.CallOption) (*v31.HttpBody, error) {
	out := new(v31.HttpBody)
	err := c.cc.Invoke(ctx, ClusterService_DownloadCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateClusterStatus(ctx context.Context, in *UpdateClusterStatusRequest, opts ...grpc.CallOption) (*UpdateClusterStatusResponse, error) {
	out := new(UpdateClusterStatusResponse)
	err := c.cc.Invoke(ctx, ClusterService_UpdateClusterStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetClusterStatus(ctx context.Context, in *GetClusterStatusRequest, opts ...grpc.CallOption) (*GetClusterStatusResponse, error) {
	out := new(GetClusterStatusResponse)
	err := c.cc.Invoke(ctx, ClusterService_GetClusterStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	CreateCluster(context.Context, *v3.Cluster) (*v3.Cluster, error)
	GetClusters(context.Context, *v31.QueryOptions) (*v3.ClusterList, error)
	GetCluster(context.Context, *v3.Cluster) (*v3.Cluster, error)
	UpdateCluster(context.Context, *v3.Cluster) (*v3.Cluster, error)
	DeleteCluster(context.Context, *v3.Cluster) (*DeleteClusterResponse, error)
	DownloadCluster(context.Context, *v3.Cluster) (*v31.HttpBody, error)
	UpdateClusterStatus(context.Context, *UpdateClusterStatusRequest) (*UpdateClusterStatusResponse, error)
	GetClusterStatus(context.Context, *GetClusterStatusRequest) (*GetClusterStatusResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) CreateCluster(context.Context, *v3.Cluster) (*v3.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedClusterServiceServer) GetClusters(context.Context, *v31.QueryOptions) (*v3.ClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusters not implemented")
}
func (UnimplementedClusterServiceServer) GetCluster(context.Context, *v3.Cluster) (*v3.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedClusterServiceServer) UpdateCluster(context.Context, *v3.Cluster) (*v3.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (UnimplementedClusterServiceServer) DeleteCluster(context.Context, *v3.Cluster) (*DeleteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedClusterServiceServer) DownloadCluster(context.Context, *v3.Cluster) (*v31.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadCluster not implemented")
}
func (UnimplementedClusterServiceServer) UpdateClusterStatus(context.Context, *UpdateClusterStatusRequest) (*UpdateClusterStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClusterStatus not implemented")
}
func (UnimplementedClusterServiceServer) GetClusterStatus(context.Context, *GetClusterStatusRequest) (*GetClusterStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_CreateCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CreateCluster(ctx, req.(*v3.Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v31.QueryOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetClusters(ctx, req.(*v31.QueryOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetCluster(ctx, req.(*v3.Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_UpdateCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, req.(*v3.Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_DeleteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DeleteCluster(ctx, req.(*v3.Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DownloadCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DownloadCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_DownloadCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DownloadCluster(ctx, req.(*v3.Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_UpdateClusterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateClusterStatus(ctx, req.(*UpdateClusterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetClusterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetClusterStatus(ctx, req.(*GetClusterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.rpc.v3.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _ClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "GetClusters",
			Handler:    _ClusterService_GetClusters_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _ClusterService_GetCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _ClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _ClusterService_DeleteCluster_Handler,
		},
		{
			MethodName: "DownloadCluster",
			Handler:    _ClusterService_DownloadCluster_Handler,
		},
		{
			MethodName: "UpdateClusterStatus",
			Handler:    _ClusterService_UpdateClusterStatus_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _ClusterService_GetClusterStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/scheduler/cluster.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: hash/hash.proto

package hashv1

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

// HashClient is the client API for Hash service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashClient interface {
	HashUrl(ctx context.Context, in *HashUrlRequest, opts ...grpc.CallOption) (*HashUrlResponse, error)
	UnHashUrl(ctx context.Context, in *UnHashUrlRequest, opts ...grpc.CallOption) (*UnHashUrlResponse, error)
}

type hashClient struct {
	cc grpc.ClientConnInterface
}

func NewHashClient(cc grpc.ClientConnInterface) HashClient {
	return &hashClient{cc}
}

func (c *hashClient) HashUrl(ctx context.Context, in *HashUrlRequest, opts ...grpc.CallOption) (*HashUrlResponse, error) {
	out := new(HashUrlResponse)
	err := c.cc.Invoke(ctx, "/hash.Hash/HashUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashClient) UnHashUrl(ctx context.Context, in *UnHashUrlRequest, opts ...grpc.CallOption) (*UnHashUrlResponse, error) {
	out := new(UnHashUrlResponse)
	err := c.cc.Invoke(ctx, "/hash.Hash/UnHashUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashServer is the server API for Hash service.
// All implementations must embed UnimplementedHashServer
// for forward compatibility
type HashServer interface {
	HashUrl(context.Context, *HashUrlRequest) (*HashUrlResponse, error)
	UnHashUrl(context.Context, *UnHashUrlRequest) (*UnHashUrlResponse, error)
	mustEmbedUnimplementedHashServer()
}

// UnimplementedHashServer must be embedded to have forward compatible implementations.
type UnimplementedHashServer struct {
}

func (UnimplementedHashServer) HashUrl(context.Context, *HashUrlRequest) (*HashUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashUrl not implemented")
}
func (UnimplementedHashServer) UnHashUrl(context.Context, *UnHashUrlRequest) (*UnHashUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnHashUrl not implemented")
}
func (UnimplementedHashServer) mustEmbedUnimplementedHashServer() {}

// UnsafeHashServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashServer will
// result in compilation errors.
type UnsafeHashServer interface {
	mustEmbedUnimplementedHashServer()
}

func RegisterHashServer(s grpc.ServiceRegistrar, srv HashServer) {
	s.RegisterService(&Hash_ServiceDesc, srv)
}

func _Hash_HashUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashServer).HashUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Hash/HashUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashServer).HashUrl(ctx, req.(*HashUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hash_UnHashUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnHashUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashServer).UnHashUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Hash/UnHashUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashServer).UnHashUrl(ctx, req.(*UnHashUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hash_ServiceDesc is the grpc.ServiceDesc for Hash service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hash_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hash.Hash",
	HandlerType: (*HashServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HashUrl",
			Handler:    _Hash_HashUrl_Handler,
		},
		{
			MethodName: "UnHashUrl",
			Handler:    _Hash_UnHashUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hash/hash.proto",
}

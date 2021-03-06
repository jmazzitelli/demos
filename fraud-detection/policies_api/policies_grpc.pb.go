// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package policies_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PoliciesServiceClient is the client API for PoliciesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoliciesServiceClient interface {
	ListPolicies(ctx context.Context, in *PolicyOwner, opts ...grpc.CallOption) (PoliciesService_ListPoliciesClient, error)
}

type policiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoliciesServiceClient(cc grpc.ClientConnInterface) PoliciesServiceClient {
	return &policiesServiceClient{cc}
}

func (c *policiesServiceClient) ListPolicies(ctx context.Context, in *PolicyOwner, opts ...grpc.CallOption) (PoliciesService_ListPoliciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PoliciesService_serviceDesc.Streams[0], "/policies_api.PoliciesService/ListPolicies", opts...)
	if err != nil {
		return nil, err
	}
	x := &policiesServiceListPoliciesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PoliciesService_ListPoliciesClient interface {
	Recv() (*Policy, error)
	grpc.ClientStream
}

type policiesServiceListPoliciesClient struct {
	grpc.ClientStream
}

func (x *policiesServiceListPoliciesClient) Recv() (*Policy, error) {
	m := new(Policy)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PoliciesServiceServer is the server API for PoliciesService service.
// All implementations must embed UnimplementedPoliciesServiceServer
// for forward compatibility
type PoliciesServiceServer interface {
	ListPolicies(*PolicyOwner, PoliciesService_ListPoliciesServer) error
	mustEmbedUnimplementedPoliciesServiceServer()
}

// UnimplementedPoliciesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoliciesServiceServer struct {
}

func (UnimplementedPoliciesServiceServer) ListPolicies(*PolicyOwner, PoliciesService_ListPoliciesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}
func (UnimplementedPoliciesServiceServer) mustEmbedUnimplementedPoliciesServiceServer() {}

// UnsafePoliciesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoliciesServiceServer will
// result in compilation errors.
type UnsafePoliciesServiceServer interface {
	mustEmbedUnimplementedPoliciesServiceServer()
}

func RegisterPoliciesServiceServer(s *grpc.Server, srv PoliciesServiceServer) {
	s.RegisterService(&_PoliciesService_serviceDesc, srv)
}

func _PoliciesService_ListPolicies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PolicyOwner)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PoliciesServiceServer).ListPolicies(m, &policiesServiceListPoliciesServer{stream})
}

type PoliciesService_ListPoliciesServer interface {
	Send(*Policy) error
	grpc.ServerStream
}

type policiesServiceListPoliciesServer struct {
	grpc.ServerStream
}

func (x *policiesServiceListPoliciesServer) Send(m *Policy) error {
	return x.ServerStream.SendMsg(m)
}

var _PoliciesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "policies_api.PoliciesService",
	HandlerType: (*PoliciesServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPolicies",
			Handler:       _PoliciesService_ListPolicies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "policies.proto",
}

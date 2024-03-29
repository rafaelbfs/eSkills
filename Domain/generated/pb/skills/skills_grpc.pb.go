// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: skills.proto

package skills

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

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	CreateNewPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
	SearchPerson(ctx context.Context, in *PersonSearchParams, opts ...grpc.CallOption) (PersonService_SearchPersonClient, error)
	UpdatePerson(ctx context.Context, in *PersonUpdate, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) CreateNewPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/eu.terranatal.skillz.entities.PersonService/createNewPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) SearchPerson(ctx context.Context, in *PersonSearchParams, opts ...grpc.CallOption) (PersonService_SearchPersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonService_ServiceDesc.Streams[0], "/eu.terranatal.skillz.entities.PersonService/searchPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceSearchPersonClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonService_SearchPersonClient interface {
	Recv() (*Person, error)
	grpc.ClientStream
}

type personServiceSearchPersonClient struct {
	grpc.ClientStream
}

func (x *personServiceSearchPersonClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personServiceClient) UpdatePerson(ctx context.Context, in *PersonUpdate, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/eu.terranatal.skillz.entities.PersonService/updatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	CreateNewPerson(context.Context, *Person) (*Person, error)
	SearchPerson(*PersonSearchParams, PersonService_SearchPersonServer) error
	UpdatePerson(context.Context, *PersonUpdate) (*UpdateResponse, error)
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) CreateNewPerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewPerson not implemented")
}
func (UnimplementedPersonServiceServer) SearchPerson(*PersonSearchParams, PersonService_SearchPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchPerson not implemented")
}
func (UnimplementedPersonServiceServer) UpdatePerson(context.Context, *PersonUpdate) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_CreateNewPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreateNewPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eu.terranatal.skillz.entities.PersonService/createNewPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreateNewPerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_SearchPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PersonSearchParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonServiceServer).SearchPerson(m, &personServiceSearchPersonServer{stream})
}

type PersonService_SearchPersonServer interface {
	Send(*Person) error
	grpc.ServerStream
}

type personServiceSearchPersonServer struct {
	grpc.ServerStream
}

func (x *personServiceSearchPersonServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

func _PersonService_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eu.terranatal.skillz.entities.PersonService/updatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).UpdatePerson(ctx, req.(*PersonUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eu.terranatal.skillz.entities.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNewPerson",
			Handler:    _PersonService_CreateNewPerson_Handler,
		},
		{
			MethodName: "updatePerson",
			Handler:    _PersonService_UpdatePerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchPerson",
			Handler:       _PersonService_SearchPerson_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "skills.proto",
}

// PublicationsServiceClient is the client API for PublicationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicationsServiceClient interface {
	CreateNewPublication(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*Publication, error)
	UpdatePublication(ctx context.Context, in *PublUpdate, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type publicationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicationsServiceClient(cc grpc.ClientConnInterface) PublicationsServiceClient {
	return &publicationsServiceClient{cc}
}

func (c *publicationsServiceClient) CreateNewPublication(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*Publication, error) {
	out := new(Publication)
	err := c.cc.Invoke(ctx, "/eu.terranatal.skillz.entities.PublicationsService/createNewPublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicationsServiceClient) UpdatePublication(ctx context.Context, in *PublUpdate, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/eu.terranatal.skillz.entities.PublicationsService/updatePublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicationsServiceServer is the server API for PublicationsService service.
// All implementations must embed UnimplementedPublicationsServiceServer
// for forward compatibility
type PublicationsServiceServer interface {
	CreateNewPublication(context.Context, *Publication) (*Publication, error)
	UpdatePublication(context.Context, *PublUpdate) (*UpdateResponse, error)
	mustEmbedUnimplementedPublicationsServiceServer()
}

// UnimplementedPublicationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicationsServiceServer struct {
}

func (UnimplementedPublicationsServiceServer) CreateNewPublication(context.Context, *Publication) (*Publication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewPublication not implemented")
}
func (UnimplementedPublicationsServiceServer) UpdatePublication(context.Context, *PublUpdate) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublication not implemented")
}
func (UnimplementedPublicationsServiceServer) mustEmbedUnimplementedPublicationsServiceServer() {}

// UnsafePublicationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicationsServiceServer will
// result in compilation errors.
type UnsafePublicationsServiceServer interface {
	mustEmbedUnimplementedPublicationsServiceServer()
}

func RegisterPublicationsServiceServer(s grpc.ServiceRegistrar, srv PublicationsServiceServer) {
	s.RegisterService(&PublicationsService_ServiceDesc, srv)
}

func _PublicationsService_CreateNewPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicationsServiceServer).CreateNewPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eu.terranatal.skillz.entities.PublicationsService/createNewPublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicationsServiceServer).CreateNewPublication(ctx, req.(*Publication))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicationsService_UpdatePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicationsServiceServer).UpdatePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eu.terranatal.skillz.entities.PublicationsService/updatePublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicationsServiceServer).UpdatePublication(ctx, req.(*PublUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicationsService_ServiceDesc is the grpc.ServiceDesc for PublicationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eu.terranatal.skillz.entities.PublicationsService",
	HandlerType: (*PublicationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNewPublication",
			Handler:    _PublicationsService_CreateNewPublication_Handler,
		},
		{
			MethodName: "updatePublication",
			Handler:    _PublicationsService_UpdatePublication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skills.proto",
}

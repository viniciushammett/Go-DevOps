package proto

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

// PetStoreClient is the client API for PetStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetStoreClient interface {
	// Adds pets to the pet store.
	AddPets(ctx context.Context, in *AddPetsReq, opts ...grpc.CallOption) (*AddPetsResp, error)
	// Updates pets entries in the store.
	UpdatePets(ctx context.Context, in *UpdatePetsReq, opts ...grpc.CallOption) (*UpdatePetsResp, error)
	// Deletes pets from the pet store.
	DeletePets(ctx context.Context, in *DeletePetsReq, opts ...grpc.CallOption) (*DeletePetsResp, error)
	// Finds pets in the pet store.
	SearchPets(ctx context.Context, in *SearchPetsReq, opts ...grpc.CallOption) (PetStore_SearchPetsClient, error)
	// Changes the OTEL sampling type.
	ChangeSampler(ctx context.Context, in *ChangeSamplerReq, opts ...grpc.CallOption) (*ChangeSamplerResp, error)
}

type petStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewPetStoreClient(cc grpc.ClientConnInterface) PetStoreClient {
	return &petStoreClient{cc}
}

func (c *petStoreClient) AddPets(ctx context.Context, in *AddPetsReq, opts ...grpc.CallOption) (*AddPetsResp, error) {
	out := new(AddPetsResp)
	err := c.cc.Invoke(ctx, "/petstore.PetStore/AddPets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) UpdatePets(ctx context.Context, in *UpdatePetsReq, opts ...grpc.CallOption) (*UpdatePetsResp, error) {
	out := new(UpdatePetsResp)
	err := c.cc.Invoke(ctx, "/petstore.PetStore/UpdatePets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) DeletePets(ctx context.Context, in *DeletePetsReq, opts ...grpc.CallOption) (*DeletePetsResp, error) {
	out := new(DeletePetsResp)
	err := c.cc.Invoke(ctx, "/petstore.PetStore/DeletePets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) SearchPets(ctx context.Context, in *SearchPetsReq, opts ...grpc.CallOption) (PetStore_SearchPetsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PetStore_ServiceDesc.Streams[0], "/petstore.PetStore/SearchPets", opts...)
	if err != nil {
		return nil, err
	}
	x := &petStoreSearchPetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PetStore_SearchPetsClient interface {
	Recv() (*Pet, error)
	grpc.ClientStream
}

type petStoreSearchPetsClient struct {
	grpc.ClientStream
}

func (x *petStoreSearchPetsClient) Recv() (*Pet, error) {
	m := new(Pet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *petStoreClient) ChangeSampler(ctx context.Context, in *ChangeSamplerReq, opts ...grpc.CallOption) (*ChangeSamplerResp, error) {
	out := new(ChangeSamplerResp)
	err := c.cc.Invoke(ctx, "/petstore.PetStore/ChangeSampler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetStoreServer is the server API for PetStore service.
// All implementations must embed UnimplementedPetStoreServer
// for forward compatibility
type PetStoreServer interface {
	// Adds pets to the pet store.
	AddPets(context.Context, *AddPetsReq) (*AddPetsResp, error)
	// Updates pets entries in the store.
	UpdatePets(context.Context, *UpdatePetsReq) (*UpdatePetsResp, error)
	// Deletes pets from the pet store.
	DeletePets(context.Context, *DeletePetsReq) (*DeletePetsResp, error)
	// Finds pets in the pet store.
	SearchPets(*SearchPetsReq, PetStore_SearchPetsServer) error
	// Changes the OTEL sampling type.
	ChangeSampler(context.Context, *ChangeSamplerReq) (*ChangeSamplerResp, error)
	mustEmbedUnimplementedPetStoreServer()
}

// UnimplementedPetStoreServer must be embedded to have forward compatible implementations.
type UnimplementedPetStoreServer struct {
}

func (UnimplementedPetStoreServer) AddPets(context.Context, *AddPetsReq) (*AddPetsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPets not implemented")
}
func (UnimplementedPetStoreServer) UpdatePets(context.Context, *UpdatePetsReq) (*UpdatePetsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePets not implemented")
}
func (UnimplementedPetStoreServer) DeletePets(context.Context, *DeletePetsReq) (*DeletePetsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePets not implemented")
}
func (UnimplementedPetStoreServer) SearchPets(*SearchPetsReq, PetStore_SearchPetsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchPets not implemented")
}
func (UnimplementedPetStoreServer) ChangeSampler(context.Context, *ChangeSamplerReq) (*ChangeSamplerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSampler not implemented")
}
func (UnimplementedPetStoreServer) mustEmbedUnimplementedPetStoreServer() {}

// UnsafePetStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetStoreServer will
// result in compilation errors.
type UnsafePetStoreServer interface {
	mustEmbedUnimplementedPetStoreServer()
}

func RegisterPetStoreServer(s grpc.ServiceRegistrar, srv PetStoreServer) {
	s.RegisterService(&PetStore_ServiceDesc, srv)
}

func _PetStore_AddPets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).AddPets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetStore/AddPets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).AddPets(ctx, req.(*AddPetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_UpdatePets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).UpdatePets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetStore/UpdatePets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).UpdatePets(ctx, req.(*UpdatePetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_DeletePets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).DeletePets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetStore/DeletePets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).DeletePets(ctx, req.(*DeletePetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_SearchPets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchPetsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PetStoreServer).SearchPets(m, &petStoreSearchPetsServer{stream})
}

type PetStore_SearchPetsServer interface {
	Send(*Pet) error
	grpc.ServerStream
}

type petStoreSearchPetsServer struct {
	grpc.ServerStream
}

func (x *petStoreSearchPetsServer) Send(m *Pet) error {
	return x.ServerStream.SendMsg(m)
}

func _PetStore_ChangeSampler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSamplerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).ChangeSampler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetStore/ChangeSampler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).ChangeSampler(ctx, req.(*ChangeSamplerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PetStore_ServiceDesc is the grpc.ServiceDesc for PetStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petstore.PetStore",
	HandlerType: (*PetStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPets",
			Handler:    _PetStore_AddPets_Handler,
		},
		{
			MethodName: "UpdatePets",
			Handler:    _PetStore_UpdatePets_Handler,
		},
		{
			MethodName: "DeletePets",
			Handler:    _PetStore_DeletePets_Handler,
		},
		{
			MethodName: "ChangeSampler",
			Handler:    _PetStore_ChangeSampler_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchPets",
			Handler:       _PetStore_SearchPets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "petstore.proto",
}
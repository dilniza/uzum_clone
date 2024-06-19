// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: order_products.proto

package order_service

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
	OrderProductsService_Create_FullMethodName  = "/genproto.OrderProductsService/Create"
	OrderProductsService_GetByID_FullMethodName = "/genproto.OrderProductsService/GetByID"
	OrderProductsService_GetALL_FullMethodName  = "/genproto.OrderProductsService/GetALL"
	OrderProductsService_Update_FullMethodName  = "/genproto.OrderProductsService/Update"
	OrderProductsService_Delete_FullMethodName  = "/genproto.OrderProductsService/Delete"
)

// OrderProductsServiceClient is the client API for OrderProductsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderProductsServiceClient interface {
	Create(ctx context.Context, in *CreateOrderProducts, opts ...grpc.CallOption) (*GetOrderProducts, error)
	GetByID(ctx context.Context, in *OrderProductsPrimaryKey, opts ...grpc.CallOption) (*GetOrderProducts, error)
	GetALL(ctx context.Context, in *GetAllOrderRequest, opts ...grpc.CallOption) (*GetAllOrderResponse, error)
	Update(ctx context.Context, in *UpdateOrderProducts, opts ...grpc.CallOption) (*GetOrderProducts, error)
	Delete(ctx context.Context, in *OrderProductsPrimaryKey, opts ...grpc.CallOption) (*EmptyOrder, error)
}

type orderProductsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderProductsServiceClient(cc grpc.ClientConnInterface) OrderProductsServiceClient {
	return &orderProductsServiceClient{cc}
}

func (c *orderProductsServiceClient) Create(ctx context.Context, in *CreateOrderProducts, opts ...grpc.CallOption) (*GetOrderProducts, error) {
	out := new(GetOrderProducts)
	err := c.cc.Invoke(ctx, OrderProductsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductsServiceClient) GetByID(ctx context.Context, in *OrderProductsPrimaryKey, opts ...grpc.CallOption) (*GetOrderProducts, error) {
	out := new(GetOrderProducts)
	err := c.cc.Invoke(ctx, OrderProductsService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductsServiceClient) GetALL(ctx context.Context, in *GetAllOrderRequest, opts ...grpc.CallOption) (*GetAllOrderResponse, error) {
	out := new(GetAllOrderResponse)
	err := c.cc.Invoke(ctx, OrderProductsService_GetALL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductsServiceClient) Update(ctx context.Context, in *UpdateOrderProducts, opts ...grpc.CallOption) (*GetOrderProducts, error) {
	out := new(GetOrderProducts)
	err := c.cc.Invoke(ctx, OrderProductsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductsServiceClient) Delete(ctx context.Context, in *OrderProductsPrimaryKey, opts ...grpc.CallOption) (*EmptyOrder, error) {
	out := new(EmptyOrder)
	err := c.cc.Invoke(ctx, OrderProductsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderProductsServiceServer is the server API for OrderProductsService service.
// All implementations should embed UnimplementedOrderProductsServiceServer
// for forward compatibility
type OrderProductsServiceServer interface {
	Create(context.Context, *CreateOrderProducts) (*GetOrderProducts, error)
	GetByID(context.Context, *OrderProductsPrimaryKey) (*GetOrderProducts, error)
	GetALL(context.Context, *GetAllOrderRequest) (*GetAllOrderResponse, error)
	Update(context.Context, *UpdateOrderProducts) (*GetOrderProducts, error)
	Delete(context.Context, *OrderProductsPrimaryKey) (*EmptyOrder, error)
}

// UnimplementedOrderProductsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderProductsServiceServer struct {
}

func (UnimplementedOrderProductsServiceServer) Create(context.Context, *CreateOrderProducts) (*GetOrderProducts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderProductsServiceServer) GetByID(context.Context, *OrderProductsPrimaryKey) (*GetOrderProducts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedOrderProductsServiceServer) GetALL(context.Context, *GetAllOrderRequest) (*GetAllOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetALL not implemented")
}
func (UnimplementedOrderProductsServiceServer) Update(context.Context, *UpdateOrderProducts) (*GetOrderProducts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderProductsServiceServer) Delete(context.Context, *OrderProductsPrimaryKey) (*EmptyOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeOrderProductsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderProductsServiceServer will
// result in compilation errors.
type UnsafeOrderProductsServiceServer interface {
	mustEmbedUnimplementedOrderProductsServiceServer()
}

func RegisterOrderProductsServiceServer(s grpc.ServiceRegistrar, srv OrderProductsServiceServer) {
	s.RegisterService(&OrderProductsService_ServiceDesc, srv)
}

func _OrderProductsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderProducts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductsServiceServer).Create(ctx, req.(*CreateOrderProducts))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductsService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProductsPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductsServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductsService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductsServiceServer).GetByID(ctx, req.(*OrderProductsPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductsService_GetALL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductsServiceServer).GetALL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductsService_GetALL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductsServiceServer).GetALL(ctx, req.(*GetAllOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderProducts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductsServiceServer).Update(ctx, req.(*UpdateOrderProducts))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProductsPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductsServiceServer).Delete(ctx, req.(*OrderProductsPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderProductsService_ServiceDesc is the grpc.ServiceDesc for OrderProductsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderProductsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.OrderProductsService",
	HandlerType: (*OrderProductsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderProductsService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _OrderProductsService_GetByID_Handler,
		},
		{
			MethodName: "GetALL",
			Handler:    _OrderProductsService_GetALL_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderProductsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrderProductsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_products.proto",
}

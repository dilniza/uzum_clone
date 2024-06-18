// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: product_review.proto

package catalog_service

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
	ProductReviewService_Create_FullMethodName  = "/genproto.ProductReviewService/Create"
	ProductReviewService_GetByID_FullMethodName = "/genproto.ProductReviewService/GetByID"
	ProductReviewService_GetAll_FullMethodName  = "/genproto.ProductReviewService/GetAll"
	ProductReviewService_Update_FullMethodName  = "/genproto.ProductReviewService/Update"
	ProductReviewService_Delete_FullMethodName  = "/genproto.ProductReviewService/Delete"
)

// ProductReviewServiceClient is the client API for ProductReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductReviewServiceClient interface {
	Create(ctx context.Context, in *CreateProductReview, opts ...grpc.CallOption) (*ProductReview, error)
	GetByID(ctx context.Context, in *ProductReviewPrimaryKey, opts ...grpc.CallOption) (*ProductReview, error)
	GetAll(ctx context.Context, in *GetAllProductReviewRequest, opts ...grpc.CallOption) (*GetAllProductReviewResponse, error)
	Update(ctx context.Context, in *UpdateProductReview, opts ...grpc.CallOption) (*ProductReview, error)
	Delete(ctx context.Context, in *ProductReviewPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type productReviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductReviewServiceClient(cc grpc.ClientConnInterface) ProductReviewServiceClient {
	return &productReviewServiceClient{cc}
}

func (c *productReviewServiceClient) Create(ctx context.Context, in *CreateProductReview, opts ...grpc.CallOption) (*ProductReview, error) {
	out := new(ProductReview)
	err := c.cc.Invoke(ctx, ProductReviewService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productReviewServiceClient) GetByID(ctx context.Context, in *ProductReviewPrimaryKey, opts ...grpc.CallOption) (*ProductReview, error) {
	out := new(ProductReview)
	err := c.cc.Invoke(ctx, ProductReviewService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productReviewServiceClient) GetAll(ctx context.Context, in *GetAllProductReviewRequest, opts ...grpc.CallOption) (*GetAllProductReviewResponse, error) {
	out := new(GetAllProductReviewResponse)
	err := c.cc.Invoke(ctx, ProductReviewService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productReviewServiceClient) Update(ctx context.Context, in *UpdateProductReview, opts ...grpc.CallOption) (*ProductReview, error) {
	out := new(ProductReview)
	err := c.cc.Invoke(ctx, ProductReviewService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productReviewServiceClient) Delete(ctx context.Context, in *ProductReviewPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ProductReviewService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductReviewServiceServer is the server API for ProductReviewService service.
// All implementations must embed UnimplementedProductReviewServiceServer
// for forward compatibility
type ProductReviewServiceServer interface {
	Create(context.Context, *CreateProductReview) (*ProductReview, error)
	GetByID(context.Context, *ProductReviewPrimaryKey) (*ProductReview, error)
	GetAll(context.Context, *GetAllProductReviewRequest) (*GetAllProductReviewResponse, error)
	Update(context.Context, *UpdateProductReview) (*ProductReview, error)
	Delete(context.Context, *ProductReviewPrimaryKey) (*Empty, error)
	mustEmbedUnimplementedProductReviewServiceServer()
}

// UnimplementedProductReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductReviewServiceServer struct {
}

func (UnimplementedProductReviewServiceServer) Create(context.Context, *CreateProductReview) (*ProductReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductReviewServiceServer) GetByID(context.Context, *ProductReviewPrimaryKey) (*ProductReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedProductReviewServiceServer) GetAll(context.Context, *GetAllProductReviewRequest) (*GetAllProductReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedProductReviewServiceServer) Update(context.Context, *UpdateProductReview) (*ProductReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductReviewServiceServer) Delete(context.Context, *ProductReviewPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductReviewServiceServer) mustEmbedUnimplementedProductReviewServiceServer() {}

// UnsafeProductReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductReviewServiceServer will
// result in compilation errors.
type UnsafeProductReviewServiceServer interface {
	mustEmbedUnimplementedProductReviewServiceServer()
}

func RegisterProductReviewServiceServer(s grpc.ServiceRegistrar, srv ProductReviewServiceServer) {
	s.RegisterService(&ProductReviewService_ServiceDesc, srv)
}

func _ProductReviewService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductReviewServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductReviewService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductReviewServiceServer).Create(ctx, req.(*CreateProductReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductReviewService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReviewPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductReviewServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductReviewService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductReviewServiceServer).GetByID(ctx, req.(*ProductReviewPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductReviewService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductReviewServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductReviewService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductReviewServiceServer).GetAll(ctx, req.(*GetAllProductReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductReviewService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductReviewServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductReviewService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductReviewServiceServer).Update(ctx, req.(*UpdateProductReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductReviewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReviewPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductReviewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductReviewService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductReviewServiceServer).Delete(ctx, req.(*ProductReviewPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductReviewService_ServiceDesc is the grpc.ServiceDesc for ProductReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ProductReviewService",
	HandlerType: (*ProductReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductReviewService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _ProductReviewService_GetByID_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProductReviewService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductReviewService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductReviewService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_review.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: product_category.proto

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
	ProductCategoryService_Create_FullMethodName  = "/genproto.ProductCategoryService/Create"
	ProductCategoryService_GetByID_FullMethodName = "/genproto.ProductCategoryService/GetByID"
	ProductCategoryService_GetAll_FullMethodName  = "/genproto.ProductCategoryService/GetAll"
	ProductCategoryService_Update_FullMethodName  = "/genproto.ProductCategoryService/Update"
	ProductCategoryService_Delete_FullMethodName  = "/genproto.ProductCategoryService/Delete"
)

// ProductCategoryServiceClient is the client API for ProductCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCategoryServiceClient interface {
	Create(ctx context.Context, in *CreateProductCategory, opts ...grpc.CallOption) (*ProductCategory, error)
	GetByID(ctx context.Context, in *ProductCategoryPrimaryKey, opts ...grpc.CallOption) (*ProductCategory, error)
	GetAll(ctx context.Context, in *GetAllProductCategoryRequest, opts ...grpc.CallOption) (*GetAllProductCategoryResponse, error)
	Update(ctx context.Context, in *UpdateProductCategory, opts ...grpc.CallOption) (*ProductCategory, error)
	Delete(ctx context.Context, in *ProductCategoryPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type productCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCategoryServiceClient(cc grpc.ClientConnInterface) ProductCategoryServiceClient {
	return &productCategoryServiceClient{cc}
}

func (c *productCategoryServiceClient) Create(ctx context.Context, in *CreateProductCategory, opts ...grpc.CallOption) (*ProductCategory, error) {
	out := new(ProductCategory)
	err := c.cc.Invoke(ctx, ProductCategoryService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) GetByID(ctx context.Context, in *ProductCategoryPrimaryKey, opts ...grpc.CallOption) (*ProductCategory, error) {
	out := new(ProductCategory)
	err := c.cc.Invoke(ctx, ProductCategoryService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) GetAll(ctx context.Context, in *GetAllProductCategoryRequest, opts ...grpc.CallOption) (*GetAllProductCategoryResponse, error) {
	out := new(GetAllProductCategoryResponse)
	err := c.cc.Invoke(ctx, ProductCategoryService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) Update(ctx context.Context, in *UpdateProductCategory, opts ...grpc.CallOption) (*ProductCategory, error) {
	out := new(ProductCategory)
	err := c.cc.Invoke(ctx, ProductCategoryService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) Delete(ctx context.Context, in *ProductCategoryPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ProductCategoryService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCategoryServiceServer is the server API for ProductCategoryService service.
// All implementations should embed UnimplementedProductCategoryServiceServer
// for forward compatibility
type ProductCategoryServiceServer interface {
	Create(context.Context, *CreateProductCategory) (*ProductCategory, error)
	GetByID(context.Context, *ProductCategoryPrimaryKey) (*ProductCategory, error)
	GetAll(context.Context, *GetAllProductCategoryRequest) (*GetAllProductCategoryResponse, error)
	Update(context.Context, *UpdateProductCategory) (*ProductCategory, error)
	Delete(context.Context, *ProductCategoryPrimaryKey) (*Empty, error)
}

// UnimplementedProductCategoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductCategoryServiceServer struct {
}

func (UnimplementedProductCategoryServiceServer) Create(context.Context, *CreateProductCategory) (*ProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductCategoryServiceServer) GetByID(context.Context, *ProductCategoryPrimaryKey) (*ProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedProductCategoryServiceServer) GetAll(context.Context, *GetAllProductCategoryRequest) (*GetAllProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedProductCategoryServiceServer) Update(context.Context, *UpdateProductCategory) (*ProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductCategoryServiceServer) Delete(context.Context, *ProductCategoryPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeProductCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCategoryServiceServer will
// result in compilation errors.
type UnsafeProductCategoryServiceServer interface {
	mustEmbedUnimplementedProductCategoryServiceServer()
}

func RegisterProductCategoryServiceServer(s grpc.ServiceRegistrar, srv ProductCategoryServiceServer) {
	s.RegisterService(&ProductCategoryService_ServiceDesc, srv)
}

func _ProductCategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).Create(ctx, req.(*CreateProductCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCategoryPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).GetByID(ctx, req.(*ProductCategoryPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).GetAll(ctx, req.(*GetAllProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).Update(ctx, req.(*UpdateProductCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCategoryPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).Delete(ctx, req.(*ProductCategoryPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCategoryService_ServiceDesc is the grpc.ServiceDesc for ProductCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ProductCategoryService",
	HandlerType: (*ProductCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductCategoryService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _ProductCategoryService_GetByID_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProductCategoryService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductCategoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductCategoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_category.proto",
}

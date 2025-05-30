// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: paygateway.proto

package paygateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PayGatewayService_CreatePayChannel_FullMethodName       = "/paygateway.PayGatewayService/CreatePayChannel"
	PayGatewayService_GetPayChannel_FullMethodName          = "/paygateway.PayGatewayService/GetPayChannel"
	PayGatewayService_UpdatePayChannel_FullMethodName       = "/paygateway.PayGatewayService/UpdatePayChannel"
	PayGatewayService_ListPayChannel_FullMethodName         = "/paygateway.PayGatewayService/ListPayChannel"
	PayGatewayService_ActivatePayChannel_FullMethodName     = "/paygateway.PayGatewayService/ActivatePayChannel"
	PayGatewayService_InactivatePayChannel_FullMethodName   = "/paygateway.PayGatewayService/InactivatePayChannel"
	PayGatewayService_CreatePayChannelParams_FullMethodName = "/paygateway.PayGatewayService/CreatePayChannelParams"
	PayGatewayService_GetPayChannelParams_FullMethodName    = "/paygateway.PayGatewayService/GetPayChannelParams"
	PayGatewayService_UpdatePayChannelParams_FullMethodName = "/paygateway.PayGatewayService/UpdatePayChannelParams"
	PayGatewayService_CreatePayChannelOrder_FullMethodName  = "/paygateway.PayGatewayService/CreatePayChannelOrder"
	PayGatewayService_GetPayChannelOrder_FullMethodName     = "/paygateway.PayGatewayService/GetPayChannelOrder"
	PayGatewayService_UpdatePayChannelOrder_FullMethodName  = "/paygateway.PayGatewayService/UpdatePayChannelOrder"
	PayGatewayService_CreatePayChannelQuota_FullMethodName  = "/paygateway.PayGatewayService/CreatePayChannelQuota"
	PayGatewayService_GetPayChannelQuota_FullMethodName     = "/paygateway.PayGatewayService/GetPayChannelQuota"
	PayGatewayService_UpdatePayChannelQuota_FullMethodName  = "/paygateway.PayGatewayService/UpdatePayChannelQuota"
	PayGatewayService_CreatePayChannelFee_FullMethodName    = "/paygateway.PayGatewayService/CreatePayChannelFee"
	PayGatewayService_GetPayChannelFee_FullMethodName       = "/paygateway.PayGatewayService/GetPayChannelFee"
	PayGatewayService_UpdatePayChannelFee_FullMethodName    = "/paygateway.PayGatewayService/UpdatePayChannelFee"
)

// PayGatewayServiceClient is the client API for PayGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayGatewayServiceClient interface {
	// 创建支付渠道
	CreatePayChannel(ctx context.Context, in *CreatePayChannelRequest, opts ...grpc.CallOption) (*CreatePayChannelResponse, error)
	// 获取支付渠道
	GetPayChannel(ctx context.Context, in *GetPayChannelRequest, opts ...grpc.CallOption) (*GetPayChannelResponse, error)
	// 更新支付渠道
	UpdatePayChannel(ctx context.Context, in *UpdatePayChannelRequest, opts ...grpc.CallOption) (*UpdatePayChannelResponse, error)
	// 获取支付渠道列表
	ListPayChannel(ctx context.Context, in *ListPayChannelRequest, opts ...grpc.CallOption) (*ListPayChannelResponse, error)
	// 激活支付渠道
	ActivatePayChannel(ctx context.Context, in *ActivatePayChannelRequest, opts ...grpc.CallOption) (*ActivatePayChannelResponse, error)
	// 禁用支付渠道
	InactivatePayChannel(ctx context.Context, in *InactivatePayChannelRequest, opts ...grpc.CallOption) (*InactivatePayChannelResponse, error)
	// 创建支付渠道参数
	CreatePayChannelParams(ctx context.Context, in *CreatePayChannelParamsRequest, opts ...grpc.CallOption) (*CreatePayChannelParamsResponse, error)
	// 获取支付渠道参数
	GetPayChannelParams(ctx context.Context, in *GetPayChannelParamsRequest, opts ...grpc.CallOption) (*GetPayChannelParamsResponse, error)
	// 更新支付渠道参数
	UpdatePayChannelParams(ctx context.Context, in *UpdatePayChannelParamsRequest, opts ...grpc.CallOption) (*UpdatePayChannelParamsResponse, error)
	// 创建支付渠道订单
	CreatePayChannelOrder(ctx context.Context, in *CreatePayChannelOrderRequest, opts ...grpc.CallOption) (*CreatePayChannelOrderResponse, error)
	// 获取支付渠道订单
	GetPayChannelOrder(ctx context.Context, in *GetPayChannelOrderRequest, opts ...grpc.CallOption) (*GetPayChannelOrderResponse, error)
	// 更新支付渠道订单
	UpdatePayChannelOrder(ctx context.Context, in *UpdatePayChannelOrderRequest, opts ...grpc.CallOption) (*UpdatePayChannelOrderResponse, error)
	// 创建支付渠道配额
	CreatePayChannelQuota(ctx context.Context, in *CreatePayChannelQuotaRequest, opts ...grpc.CallOption) (*CreatePayChannelQuotaResponse, error)
	// 获取支付渠道配额
	GetPayChannelQuota(ctx context.Context, in *GetPayChannelQuotaRequest, opts ...grpc.CallOption) (*GetPayChannelQuotaResponse, error)
	// 更新支付渠道配额
	UpdatePayChannelQuota(ctx context.Context, in *UpdatePayChannelQuotaRequest, opts ...grpc.CallOption) (*UpdatePayChannelQuotaResponse, error)
	// 创建支付渠道费率
	CreatePayChannelFee(ctx context.Context, in *CreatePayChannelFeeRequest, opts ...grpc.CallOption) (*CreatePayChannelFeeResponse, error)
	// 获取支付渠道费率
	GetPayChannelFee(ctx context.Context, in *GetPayChannelFeeRequest, opts ...grpc.CallOption) (*GetPayChannelFeeResponse, error)
	// 更新支付渠道费率
	UpdatePayChannelFee(ctx context.Context, in *UpdatePayChannelFeeRequest, opts ...grpc.CallOption) (*UpdatePayChannelFeeResponse, error)
}

type payGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayGatewayServiceClient(cc grpc.ClientConnInterface) PayGatewayServiceClient {
	return &payGatewayServiceClient{cc}
}

func (c *payGatewayServiceClient) CreatePayChannel(ctx context.Context, in *CreatePayChannelRequest, opts ...grpc.CallOption) (*CreatePayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePayChannelResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_CreatePayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) GetPayChannel(ctx context.Context, in *GetPayChannelRequest, opts ...grpc.CallOption) (*GetPayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPayChannelResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_GetPayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) UpdatePayChannel(ctx context.Context, in *UpdatePayChannelRequest, opts ...grpc.CallOption) (*UpdatePayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePayChannelResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_UpdatePayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) ListPayChannel(ctx context.Context, in *ListPayChannelRequest, opts ...grpc.CallOption) (*ListPayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPayChannelResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_ListPayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) ActivatePayChannel(ctx context.Context, in *ActivatePayChannelRequest, opts ...grpc.CallOption) (*ActivatePayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivatePayChannelResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_ActivatePayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) InactivatePayChannel(ctx context.Context, in *InactivatePayChannelRequest, opts ...grpc.CallOption) (*InactivatePayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InactivatePayChannelResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_InactivatePayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) CreatePayChannelParams(ctx context.Context, in *CreatePayChannelParamsRequest, opts ...grpc.CallOption) (*CreatePayChannelParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePayChannelParamsResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_CreatePayChannelParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) GetPayChannelParams(ctx context.Context, in *GetPayChannelParamsRequest, opts ...grpc.CallOption) (*GetPayChannelParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPayChannelParamsResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_GetPayChannelParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) UpdatePayChannelParams(ctx context.Context, in *UpdatePayChannelParamsRequest, opts ...grpc.CallOption) (*UpdatePayChannelParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePayChannelParamsResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_UpdatePayChannelParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) CreatePayChannelOrder(ctx context.Context, in *CreatePayChannelOrderRequest, opts ...grpc.CallOption) (*CreatePayChannelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePayChannelOrderResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_CreatePayChannelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) GetPayChannelOrder(ctx context.Context, in *GetPayChannelOrderRequest, opts ...grpc.CallOption) (*GetPayChannelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPayChannelOrderResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_GetPayChannelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) UpdatePayChannelOrder(ctx context.Context, in *UpdatePayChannelOrderRequest, opts ...grpc.CallOption) (*UpdatePayChannelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePayChannelOrderResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_UpdatePayChannelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) CreatePayChannelQuota(ctx context.Context, in *CreatePayChannelQuotaRequest, opts ...grpc.CallOption) (*CreatePayChannelQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePayChannelQuotaResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_CreatePayChannelQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) GetPayChannelQuota(ctx context.Context, in *GetPayChannelQuotaRequest, opts ...grpc.CallOption) (*GetPayChannelQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPayChannelQuotaResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_GetPayChannelQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) UpdatePayChannelQuota(ctx context.Context, in *UpdatePayChannelQuotaRequest, opts ...grpc.CallOption) (*UpdatePayChannelQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePayChannelQuotaResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_UpdatePayChannelQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) CreatePayChannelFee(ctx context.Context, in *CreatePayChannelFeeRequest, opts ...grpc.CallOption) (*CreatePayChannelFeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePayChannelFeeResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_CreatePayChannelFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) GetPayChannelFee(ctx context.Context, in *GetPayChannelFeeRequest, opts ...grpc.CallOption) (*GetPayChannelFeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPayChannelFeeResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_GetPayChannelFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payGatewayServiceClient) UpdatePayChannelFee(ctx context.Context, in *UpdatePayChannelFeeRequest, opts ...grpc.CallOption) (*UpdatePayChannelFeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePayChannelFeeResponse)
	err := c.cc.Invoke(ctx, PayGatewayService_UpdatePayChannelFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayGatewayServiceServer is the server API for PayGatewayService service.
// All implementations must embed UnimplementedPayGatewayServiceServer
// for forward compatibility.
type PayGatewayServiceServer interface {
	// 创建支付渠道
	CreatePayChannel(context.Context, *CreatePayChannelRequest) (*CreatePayChannelResponse, error)
	// 获取支付渠道
	GetPayChannel(context.Context, *GetPayChannelRequest) (*GetPayChannelResponse, error)
	// 更新支付渠道
	UpdatePayChannel(context.Context, *UpdatePayChannelRequest) (*UpdatePayChannelResponse, error)
	// 获取支付渠道列表
	ListPayChannel(context.Context, *ListPayChannelRequest) (*ListPayChannelResponse, error)
	// 激活支付渠道
	ActivatePayChannel(context.Context, *ActivatePayChannelRequest) (*ActivatePayChannelResponse, error)
	// 禁用支付渠道
	InactivatePayChannel(context.Context, *InactivatePayChannelRequest) (*InactivatePayChannelResponse, error)
	// 创建支付渠道参数
	CreatePayChannelParams(context.Context, *CreatePayChannelParamsRequest) (*CreatePayChannelParamsResponse, error)
	// 获取支付渠道参数
	GetPayChannelParams(context.Context, *GetPayChannelParamsRequest) (*GetPayChannelParamsResponse, error)
	// 更新支付渠道参数
	UpdatePayChannelParams(context.Context, *UpdatePayChannelParamsRequest) (*UpdatePayChannelParamsResponse, error)
	// 创建支付渠道订单
	CreatePayChannelOrder(context.Context, *CreatePayChannelOrderRequest) (*CreatePayChannelOrderResponse, error)
	// 获取支付渠道订单
	GetPayChannelOrder(context.Context, *GetPayChannelOrderRequest) (*GetPayChannelOrderResponse, error)
	// 更新支付渠道订单
	UpdatePayChannelOrder(context.Context, *UpdatePayChannelOrderRequest) (*UpdatePayChannelOrderResponse, error)
	// 创建支付渠道配额
	CreatePayChannelQuota(context.Context, *CreatePayChannelQuotaRequest) (*CreatePayChannelQuotaResponse, error)
	// 获取支付渠道配额
	GetPayChannelQuota(context.Context, *GetPayChannelQuotaRequest) (*GetPayChannelQuotaResponse, error)
	// 更新支付渠道配额
	UpdatePayChannelQuota(context.Context, *UpdatePayChannelQuotaRequest) (*UpdatePayChannelQuotaResponse, error)
	// 创建支付渠道费率
	CreatePayChannelFee(context.Context, *CreatePayChannelFeeRequest) (*CreatePayChannelFeeResponse, error)
	// 获取支付渠道费率
	GetPayChannelFee(context.Context, *GetPayChannelFeeRequest) (*GetPayChannelFeeResponse, error)
	// 更新支付渠道费率
	UpdatePayChannelFee(context.Context, *UpdatePayChannelFeeRequest) (*UpdatePayChannelFeeResponse, error)
	mustEmbedUnimplementedPayGatewayServiceServer()
}

// UnimplementedPayGatewayServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPayGatewayServiceServer struct{}

func (UnimplementedPayGatewayServiceServer) CreatePayChannel(context.Context, *CreatePayChannelRequest) (*CreatePayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayChannel not implemented")
}
func (UnimplementedPayGatewayServiceServer) GetPayChannel(context.Context, *GetPayChannelRequest) (*GetPayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChannel not implemented")
}
func (UnimplementedPayGatewayServiceServer) UpdatePayChannel(context.Context, *UpdatePayChannelRequest) (*UpdatePayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayChannel not implemented")
}
func (UnimplementedPayGatewayServiceServer) ListPayChannel(context.Context, *ListPayChannelRequest) (*ListPayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayChannel not implemented")
}
func (UnimplementedPayGatewayServiceServer) ActivatePayChannel(context.Context, *ActivatePayChannelRequest) (*ActivatePayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivatePayChannel not implemented")
}
func (UnimplementedPayGatewayServiceServer) InactivatePayChannel(context.Context, *InactivatePayChannelRequest) (*InactivatePayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InactivatePayChannel not implemented")
}
func (UnimplementedPayGatewayServiceServer) CreatePayChannelParams(context.Context, *CreatePayChannelParamsRequest) (*CreatePayChannelParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayChannelParams not implemented")
}
func (UnimplementedPayGatewayServiceServer) GetPayChannelParams(context.Context, *GetPayChannelParamsRequest) (*GetPayChannelParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChannelParams not implemented")
}
func (UnimplementedPayGatewayServiceServer) UpdatePayChannelParams(context.Context, *UpdatePayChannelParamsRequest) (*UpdatePayChannelParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayChannelParams not implemented")
}
func (UnimplementedPayGatewayServiceServer) CreatePayChannelOrder(context.Context, *CreatePayChannelOrderRequest) (*CreatePayChannelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayChannelOrder not implemented")
}
func (UnimplementedPayGatewayServiceServer) GetPayChannelOrder(context.Context, *GetPayChannelOrderRequest) (*GetPayChannelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChannelOrder not implemented")
}
func (UnimplementedPayGatewayServiceServer) UpdatePayChannelOrder(context.Context, *UpdatePayChannelOrderRequest) (*UpdatePayChannelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayChannelOrder not implemented")
}
func (UnimplementedPayGatewayServiceServer) CreatePayChannelQuota(context.Context, *CreatePayChannelQuotaRequest) (*CreatePayChannelQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayChannelQuota not implemented")
}
func (UnimplementedPayGatewayServiceServer) GetPayChannelQuota(context.Context, *GetPayChannelQuotaRequest) (*GetPayChannelQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChannelQuota not implemented")
}
func (UnimplementedPayGatewayServiceServer) UpdatePayChannelQuota(context.Context, *UpdatePayChannelQuotaRequest) (*UpdatePayChannelQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayChannelQuota not implemented")
}
func (UnimplementedPayGatewayServiceServer) CreatePayChannelFee(context.Context, *CreatePayChannelFeeRequest) (*CreatePayChannelFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayChannelFee not implemented")
}
func (UnimplementedPayGatewayServiceServer) GetPayChannelFee(context.Context, *GetPayChannelFeeRequest) (*GetPayChannelFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChannelFee not implemented")
}
func (UnimplementedPayGatewayServiceServer) UpdatePayChannelFee(context.Context, *UpdatePayChannelFeeRequest) (*UpdatePayChannelFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayChannelFee not implemented")
}
func (UnimplementedPayGatewayServiceServer) mustEmbedUnimplementedPayGatewayServiceServer() {}
func (UnimplementedPayGatewayServiceServer) testEmbeddedByValue()                           {}

// UnsafePayGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayGatewayServiceServer will
// result in compilation errors.
type UnsafePayGatewayServiceServer interface {
	mustEmbedUnimplementedPayGatewayServiceServer()
}

func RegisterPayGatewayServiceServer(s grpc.ServiceRegistrar, srv PayGatewayServiceServer) {
	// If the following call pancis, it indicates UnimplementedPayGatewayServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PayGatewayService_ServiceDesc, srv)
}

func _PayGatewayService_CreatePayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).CreatePayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_CreatePayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).CreatePayChannel(ctx, req.(*CreatePayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_GetPayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).GetPayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_GetPayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).GetPayChannel(ctx, req.(*GetPayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_UpdatePayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).UpdatePayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_UpdatePayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).UpdatePayChannel(ctx, req.(*UpdatePayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_ListPayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).ListPayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_ListPayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).ListPayChannel(ctx, req.(*ListPayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_ActivatePayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivatePayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).ActivatePayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_ActivatePayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).ActivatePayChannel(ctx, req.(*ActivatePayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_InactivatePayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InactivatePayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).InactivatePayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_InactivatePayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).InactivatePayChannel(ctx, req.(*InactivatePayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_CreatePayChannelParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayChannelParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).CreatePayChannelParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_CreatePayChannelParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).CreatePayChannelParams(ctx, req.(*CreatePayChannelParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_GetPayChannelParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChannelParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).GetPayChannelParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_GetPayChannelParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).GetPayChannelParams(ctx, req.(*GetPayChannelParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_UpdatePayChannelParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayChannelParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).UpdatePayChannelParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_UpdatePayChannelParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).UpdatePayChannelParams(ctx, req.(*UpdatePayChannelParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_CreatePayChannelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayChannelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).CreatePayChannelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_CreatePayChannelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).CreatePayChannelOrder(ctx, req.(*CreatePayChannelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_GetPayChannelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChannelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).GetPayChannelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_GetPayChannelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).GetPayChannelOrder(ctx, req.(*GetPayChannelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_UpdatePayChannelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayChannelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).UpdatePayChannelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_UpdatePayChannelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).UpdatePayChannelOrder(ctx, req.(*UpdatePayChannelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_CreatePayChannelQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayChannelQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).CreatePayChannelQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_CreatePayChannelQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).CreatePayChannelQuota(ctx, req.(*CreatePayChannelQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_GetPayChannelQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChannelQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).GetPayChannelQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_GetPayChannelQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).GetPayChannelQuota(ctx, req.(*GetPayChannelQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_UpdatePayChannelQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayChannelQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).UpdatePayChannelQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_UpdatePayChannelQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).UpdatePayChannelQuota(ctx, req.(*UpdatePayChannelQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_CreatePayChannelFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayChannelFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).CreatePayChannelFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_CreatePayChannelFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).CreatePayChannelFee(ctx, req.(*CreatePayChannelFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_GetPayChannelFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChannelFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).GetPayChannelFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_GetPayChannelFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).GetPayChannelFee(ctx, req.(*GetPayChannelFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayGatewayService_UpdatePayChannelFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayChannelFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayGatewayServiceServer).UpdatePayChannelFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayGatewayService_UpdatePayChannelFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayGatewayServiceServer).UpdatePayChannelFee(ctx, req.(*UpdatePayChannelFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayGatewayService_ServiceDesc is the grpc.ServiceDesc for PayGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paygateway.PayGatewayService",
	HandlerType: (*PayGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayChannel",
			Handler:    _PayGatewayService_CreatePayChannel_Handler,
		},
		{
			MethodName: "GetPayChannel",
			Handler:    _PayGatewayService_GetPayChannel_Handler,
		},
		{
			MethodName: "UpdatePayChannel",
			Handler:    _PayGatewayService_UpdatePayChannel_Handler,
		},
		{
			MethodName: "ListPayChannel",
			Handler:    _PayGatewayService_ListPayChannel_Handler,
		},
		{
			MethodName: "ActivatePayChannel",
			Handler:    _PayGatewayService_ActivatePayChannel_Handler,
		},
		{
			MethodName: "InactivatePayChannel",
			Handler:    _PayGatewayService_InactivatePayChannel_Handler,
		},
		{
			MethodName: "CreatePayChannelParams",
			Handler:    _PayGatewayService_CreatePayChannelParams_Handler,
		},
		{
			MethodName: "GetPayChannelParams",
			Handler:    _PayGatewayService_GetPayChannelParams_Handler,
		},
		{
			MethodName: "UpdatePayChannelParams",
			Handler:    _PayGatewayService_UpdatePayChannelParams_Handler,
		},
		{
			MethodName: "CreatePayChannelOrder",
			Handler:    _PayGatewayService_CreatePayChannelOrder_Handler,
		},
		{
			MethodName: "GetPayChannelOrder",
			Handler:    _PayGatewayService_GetPayChannelOrder_Handler,
		},
		{
			MethodName: "UpdatePayChannelOrder",
			Handler:    _PayGatewayService_UpdatePayChannelOrder_Handler,
		},
		{
			MethodName: "CreatePayChannelQuota",
			Handler:    _PayGatewayService_CreatePayChannelQuota_Handler,
		},
		{
			MethodName: "GetPayChannelQuota",
			Handler:    _PayGatewayService_GetPayChannelQuota_Handler,
		},
		{
			MethodName: "UpdatePayChannelQuota",
			Handler:    _PayGatewayService_UpdatePayChannelQuota_Handler,
		},
		{
			MethodName: "CreatePayChannelFee",
			Handler:    _PayGatewayService_CreatePayChannelFee_Handler,
		},
		{
			MethodName: "GetPayChannelFee",
			Handler:    _PayGatewayService_GetPayChannelFee_Handler,
		},
		{
			MethodName: "UpdatePayChannelFee",
			Handler:    _PayGatewayService_UpdatePayChannelFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paygateway.proto",
}

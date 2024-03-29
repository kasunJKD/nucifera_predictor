// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/membership.proto

package membership

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	SignUp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	PasswordSignIn(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetModelDataByBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponseList, error)
	GetPredictedValuesByModelId(ctx context.Context, in *PredictedRequest, opts ...grpc.CallOption) (*PredictedResponseList, error)
	GetOriginalData(ctx context.Context, in *OriginalDataRequest, opts ...grpc.CallOption) (*OriginalDataList, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) SignUp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/membership.DataService/signUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) PasswordSignIn(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/membership.DataService/passwordSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetModelDataByBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponseList, error) {
	out := new(BatchResponseList)
	err := c.cc.Invoke(ctx, "/membership.DataService/getModelDataByBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetPredictedValuesByModelId(ctx context.Context, in *PredictedRequest, opts ...grpc.CallOption) (*PredictedResponseList, error) {
	out := new(PredictedResponseList)
	err := c.cc.Invoke(ctx, "/membership.DataService/getPredictedValuesByModelId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetOriginalData(ctx context.Context, in *OriginalDataRequest, opts ...grpc.CallOption) (*OriginalDataList, error) {
	out := new(OriginalDataList)
	err := c.cc.Invoke(ctx, "/membership.DataService/getOriginalData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	SignUp(context.Context, *Request) (*Response, error)
	PasswordSignIn(context.Context, *Request) (*Response, error)
	GetModelDataByBatch(context.Context, *BatchRequest) (*BatchResponseList, error)
	GetPredictedValuesByModelId(context.Context, *PredictedRequest) (*PredictedResponseList, error)
	GetOriginalData(context.Context, *OriginalDataRequest) (*OriginalDataList, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) SignUp(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedDataServiceServer) PasswordSignIn(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordSignIn not implemented")
}
func (UnimplementedDataServiceServer) GetModelDataByBatch(context.Context, *BatchRequest) (*BatchResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelDataByBatch not implemented")
}
func (UnimplementedDataServiceServer) GetPredictedValuesByModelId(context.Context, *PredictedRequest) (*PredictedResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredictedValuesByModelId not implemented")
}
func (UnimplementedDataServiceServer) GetOriginalData(context.Context, *OriginalDataRequest) (*OriginalDataList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalData not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.DataService/signUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).SignUp(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_PasswordSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).PasswordSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.DataService/passwordSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).PasswordSignIn(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetModelDataByBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetModelDataByBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.DataService/getModelDataByBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetModelDataByBatch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetPredictedValuesByModelId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetPredictedValuesByModelId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.DataService/getPredictedValuesByModelId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetPredictedValuesByModelId(ctx, req.(*PredictedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetOriginalData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginalDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetOriginalData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.DataService/getOriginalData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetOriginalData(ctx, req.(*OriginalDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "membership.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "signUp",
			Handler:    _DataService_SignUp_Handler,
		},
		{
			MethodName: "passwordSignIn",
			Handler:    _DataService_PasswordSignIn_Handler,
		},
		{
			MethodName: "getModelDataByBatch",
			Handler:    _DataService_GetModelDataByBatch_Handler,
		},
		{
			MethodName: "getPredictedValuesByModelId",
			Handler:    _DataService_GetPredictedValuesByModelId_Handler,
		},
		{
			MethodName: "getOriginalData",
			Handler:    _DataService_GetOriginalData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/membership.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: questionnaire/v1/questionnaire.proto

package v1

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
	Questionnaire_CreateQuestionnaire_FullMethodName = "/questionnaire.v1.Questionnaire/CreateQuestionnaire"
	Questionnaire_GetQuestionnaire_FullMethodName    = "/questionnaire.v1.Questionnaire/GetQuestionnaire"
	Questionnaire_ListQuestionnaire_FullMethodName   = "/questionnaire.v1.Questionnaire/ListQuestionnaire"
	Questionnaire_UpdateQuestionnaire_FullMethodName = "/questionnaire.v1.Questionnaire/UpdateQuestionnaire"
	Questionnaire_DeleteQuestionnaire_FullMethodName = "/questionnaire.v1.Questionnaire/DeleteQuestionnaire"
	Questionnaire_CreateQuestion_FullMethodName      = "/questionnaire.v1.Questionnaire/CreateQuestion"
	Questionnaire_GetQuestions_FullMethodName        = "/questionnaire.v1.Questionnaire/GetQuestions"
	Questionnaire_ListQuestion_FullMethodName        = "/questionnaire.v1.Questionnaire/ListQuestion"
	Questionnaire_UpdateQuestion_FullMethodName      = "/questionnaire.v1.Questionnaire/UpdateQuestion"
	Questionnaire_DeleteQuestion_FullMethodName      = "/questionnaire.v1.Questionnaire/DeleteQuestion"
	Questionnaire_SubmitAnswer_FullMethodName        = "/questionnaire.v1.Questionnaire/SubmitAnswer"
	Questionnaire_SubmitAnswerBulk_FullMethodName    = "/questionnaire.v1.Questionnaire/SubmitAnswerBulk"
)

// QuestionnaireClient is the client API for Questionnaire service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionnaireClient interface {
	CreateQuestionnaire(ctx context.Context, in *CreateQuestionnaireRequest, opts ...grpc.CallOption) (*CreateQuestionnaireReply, error)
	GetQuestionnaire(ctx context.Context, in *GetQuestionnaireRequest, opts ...grpc.CallOption) (*GetQuestionnaireReply, error)
	ListQuestionnaire(ctx context.Context, in *ListQuestionnaireRequest, opts ...grpc.CallOption) (*ListQuestionnaireReply, error)
	UpdateQuestionnaire(ctx context.Context, in *UpdateQuestionnaireRequest, opts ...grpc.CallOption) (*UpdateQuestionnaireReply, error)
	DeleteQuestionnaire(ctx context.Context, in *DeleteQuestionnaireRequest, opts ...grpc.CallOption) (*DeleteQuestionnaireReply, error)
	CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionReply, error)
	GetQuestions(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionReply, error)
	ListQuestion(ctx context.Context, in *ListQuestionRequest, opts ...grpc.CallOption) (*ListQuestionReply, error)
	UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*UpdateQuestionReply, error)
	DeleteQuestion(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*DeleteQuestionReply, error)
	SubmitAnswer(ctx context.Context, in *SubmitAnswerSingleRequest, opts ...grpc.CallOption) (*SubmitAnswerSingleReply, error)
	SubmitAnswerBulk(ctx context.Context, in *SubmitAnswerBulkRequest, opts ...grpc.CallOption) (*SubmitAnswerBulkReply, error)
}

type questionnaireClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionnaireClient(cc grpc.ClientConnInterface) QuestionnaireClient {
	return &questionnaireClient{cc}
}

func (c *questionnaireClient) CreateQuestionnaire(ctx context.Context, in *CreateQuestionnaireRequest, opts ...grpc.CallOption) (*CreateQuestionnaireReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateQuestionnaireReply)
	err := c.cc.Invoke(ctx, Questionnaire_CreateQuestionnaire_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) GetQuestionnaire(ctx context.Context, in *GetQuestionnaireRequest, opts ...grpc.CallOption) (*GetQuestionnaireReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQuestionnaireReply)
	err := c.cc.Invoke(ctx, Questionnaire_GetQuestionnaire_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) ListQuestionnaire(ctx context.Context, in *ListQuestionnaireRequest, opts ...grpc.CallOption) (*ListQuestionnaireReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListQuestionnaireReply)
	err := c.cc.Invoke(ctx, Questionnaire_ListQuestionnaire_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) UpdateQuestionnaire(ctx context.Context, in *UpdateQuestionnaireRequest, opts ...grpc.CallOption) (*UpdateQuestionnaireReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateQuestionnaireReply)
	err := c.cc.Invoke(ctx, Questionnaire_UpdateQuestionnaire_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) DeleteQuestionnaire(ctx context.Context, in *DeleteQuestionnaireRequest, opts ...grpc.CallOption) (*DeleteQuestionnaireReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteQuestionnaireReply)
	err := c.cc.Invoke(ctx, Questionnaire_DeleteQuestionnaire_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateQuestionReply)
	err := c.cc.Invoke(ctx, Questionnaire_CreateQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) GetQuestions(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQuestionReply)
	err := c.cc.Invoke(ctx, Questionnaire_GetQuestions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) ListQuestion(ctx context.Context, in *ListQuestionRequest, opts ...grpc.CallOption) (*ListQuestionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListQuestionReply)
	err := c.cc.Invoke(ctx, Questionnaire_ListQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*UpdateQuestionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateQuestionReply)
	err := c.cc.Invoke(ctx, Questionnaire_UpdateQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) DeleteQuestion(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*DeleteQuestionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteQuestionReply)
	err := c.cc.Invoke(ctx, Questionnaire_DeleteQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) SubmitAnswer(ctx context.Context, in *SubmitAnswerSingleRequest, opts ...grpc.CallOption) (*SubmitAnswerSingleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAnswerSingleReply)
	err := c.cc.Invoke(ctx, Questionnaire_SubmitAnswer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) SubmitAnswerBulk(ctx context.Context, in *SubmitAnswerBulkRequest, opts ...grpc.CallOption) (*SubmitAnswerBulkReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAnswerBulkReply)
	err := c.cc.Invoke(ctx, Questionnaire_SubmitAnswerBulk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionnaireServer is the server API for Questionnaire service.
// All implementations must embed UnimplementedQuestionnaireServer
// for forward compatibility.
type QuestionnaireServer interface {
	CreateQuestionnaire(context.Context, *CreateQuestionnaireRequest) (*CreateQuestionnaireReply, error)
	GetQuestionnaire(context.Context, *GetQuestionnaireRequest) (*GetQuestionnaireReply, error)
	ListQuestionnaire(context.Context, *ListQuestionnaireRequest) (*ListQuestionnaireReply, error)
	UpdateQuestionnaire(context.Context, *UpdateQuestionnaireRequest) (*UpdateQuestionnaireReply, error)
	DeleteQuestionnaire(context.Context, *DeleteQuestionnaireRequest) (*DeleteQuestionnaireReply, error)
	CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionReply, error)
	GetQuestions(context.Context, *GetQuestionRequest) (*GetQuestionReply, error)
	ListQuestion(context.Context, *ListQuestionRequest) (*ListQuestionReply, error)
	UpdateQuestion(context.Context, *UpdateQuestionRequest) (*UpdateQuestionReply, error)
	DeleteQuestion(context.Context, *DeleteQuestionRequest) (*DeleteQuestionReply, error)
	SubmitAnswer(context.Context, *SubmitAnswerSingleRequest) (*SubmitAnswerSingleReply, error)
	SubmitAnswerBulk(context.Context, *SubmitAnswerBulkRequest) (*SubmitAnswerBulkReply, error)
	mustEmbedUnimplementedQuestionnaireServer()
}

// UnimplementedQuestionnaireServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuestionnaireServer struct{}

func (UnimplementedQuestionnaireServer) CreateQuestionnaire(context.Context, *CreateQuestionnaireRequest) (*CreateQuestionnaireReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServer) GetQuestionnaire(context.Context, *GetQuestionnaireRequest) (*GetQuestionnaireReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServer) ListQuestionnaire(context.Context, *ListQuestionnaireRequest) (*ListQuestionnaireReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServer) UpdateQuestionnaire(context.Context, *UpdateQuestionnaireRequest) (*UpdateQuestionnaireReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServer) DeleteQuestionnaire(context.Context, *DeleteQuestionnaireRequest) (*DeleteQuestionnaireReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestionnaire not implemented")
}
func (UnimplementedQuestionnaireServer) CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedQuestionnaireServer) GetQuestions(context.Context, *GetQuestionRequest) (*GetQuestionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestions not implemented")
}
func (UnimplementedQuestionnaireServer) ListQuestion(context.Context, *ListQuestionRequest) (*ListQuestionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestion not implemented")
}
func (UnimplementedQuestionnaireServer) UpdateQuestion(context.Context, *UpdateQuestionRequest) (*UpdateQuestionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}
func (UnimplementedQuestionnaireServer) DeleteQuestion(context.Context, *DeleteQuestionRequest) (*DeleteQuestionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
func (UnimplementedQuestionnaireServer) SubmitAnswer(context.Context, *SubmitAnswerSingleRequest) (*SubmitAnswerSingleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswer not implemented")
}
func (UnimplementedQuestionnaireServer) SubmitAnswerBulk(context.Context, *SubmitAnswerBulkRequest) (*SubmitAnswerBulkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswerBulk not implemented")
}
func (UnimplementedQuestionnaireServer) mustEmbedUnimplementedQuestionnaireServer() {}
func (UnimplementedQuestionnaireServer) testEmbeddedByValue()                       {}

// UnsafeQuestionnaireServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionnaireServer will
// result in compilation errors.
type UnsafeQuestionnaireServer interface {
	mustEmbedUnimplementedQuestionnaireServer()
}

func RegisterQuestionnaireServer(s grpc.ServiceRegistrar, srv QuestionnaireServer) {
	// If the following call pancis, it indicates UnimplementedQuestionnaireServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Questionnaire_ServiceDesc, srv)
}

func _Questionnaire_CreateQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).CreateQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_CreateQuestionnaire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).CreateQuestionnaire(ctx, req.(*CreateQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_GetQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).GetQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_GetQuestionnaire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).GetQuestionnaire(ctx, req.(*GetQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_ListQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).ListQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_ListQuestionnaire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).ListQuestionnaire(ctx, req.(*ListQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_UpdateQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).UpdateQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_UpdateQuestionnaire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).UpdateQuestionnaire(ctx, req.(*UpdateQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_DeleteQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).DeleteQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_DeleteQuestionnaire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).DeleteQuestionnaire(ctx, req.(*DeleteQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_CreateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).CreateQuestion(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_GetQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).GetQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_GetQuestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).GetQuestions(ctx, req.(*GetQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_ListQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).ListQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_ListQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).ListQuestion(ctx, req.(*ListQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_UpdateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).UpdateQuestion(ctx, req.(*UpdateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_DeleteQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).DeleteQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_DeleteQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).DeleteQuestion(ctx, req.(*DeleteQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_SubmitAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswerSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).SubmitAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_SubmitAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).SubmitAnswer(ctx, req.(*SubmitAnswerSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_SubmitAnswerBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswerBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).SubmitAnswerBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_SubmitAnswerBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).SubmitAnswerBulk(ctx, req.(*SubmitAnswerBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Questionnaire_ServiceDesc is the grpc.ServiceDesc for Questionnaire service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Questionnaire_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "questionnaire.v1.Questionnaire",
	HandlerType: (*QuestionnaireServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestionnaire",
			Handler:    _Questionnaire_CreateQuestionnaire_Handler,
		},
		{
			MethodName: "GetQuestionnaire",
			Handler:    _Questionnaire_GetQuestionnaire_Handler,
		},
		{
			MethodName: "ListQuestionnaire",
			Handler:    _Questionnaire_ListQuestionnaire_Handler,
		},
		{
			MethodName: "UpdateQuestionnaire",
			Handler:    _Questionnaire_UpdateQuestionnaire_Handler,
		},
		{
			MethodName: "DeleteQuestionnaire",
			Handler:    _Questionnaire_DeleteQuestionnaire_Handler,
		},
		{
			MethodName: "CreateQuestion",
			Handler:    _Questionnaire_CreateQuestion_Handler,
		},
		{
			MethodName: "GetQuestions",
			Handler:    _Questionnaire_GetQuestions_Handler,
		},
		{
			MethodName: "ListQuestion",
			Handler:    _Questionnaire_ListQuestion_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _Questionnaire_UpdateQuestion_Handler,
		},
		{
			MethodName: "DeleteQuestion",
			Handler:    _Questionnaire_DeleteQuestion_Handler,
		},
		{
			MethodName: "SubmitAnswer",
			Handler:    _Questionnaire_SubmitAnswer_Handler,
		},
		{
			MethodName: "SubmitAnswerBulk",
			Handler:    _Questionnaire_SubmitAnswerBulk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "questionnaire/v1/questionnaire.proto",
}

// Code generated by microgen. DO NOT EDIT.

package grpc

import (
	"context"
	pb "github.com/devimteam/microgen/examples/usersvc/pb"
	transport "github.com/devimteam/microgen/examples/usersvc/pkg/usersvc/transport"
	empty "github.com/golang/protobuf/ptypes/empty"
)

//========================================= Request Encoders =========================================//

func _Encode_CreateUser_Request(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*transport.CreateUser_Request)
	return __CreateUser_Request_ToProtobuf(req)
}

func _Encode_UpdateUser_Request(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*transport.UpdateUser_Request)
	return __UpdateUser_Request_ToProtobuf(req)
}

func _Encode_GetUser_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func _Encode_FindUsers_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_CreateComment_Request(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*transport.CreateComment_Request)
	return __CreateComment_Request_ToProtobuf(req)
}

func _Encode_GetComment_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func _Encode_GetUserComments_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

//========================================= Request Decoders =========================================//

func _Decode_CreateUser_Request(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUser_Request)
	return __CreateUser_Request_FromProtobuf(req)
}

func _Decode_UpdateUser_Request(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUser_Request)
	return __UpdateUser_Request_FromProtobuf(req)
}

func _Decode_GetUser_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func _Decode_FindUsers_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func _Decode_CreateComment_Request(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateComment_Request)
	return __CreateComment_Request_FromProtobuf(req)
}

func _Decode_GetComment_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func _Decode_GetUserComments_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

//========================================= Response Encoders =========================================//

func _Encode_CreateUser_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func _Encode_UpdateUser_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_GetUser_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*transport.GetUser_Response)
	return __GetUser_Response_ToProtobuf(resp)
}

func _Encode_FindUsers_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*transport.FindUsers_Response)
	return __FindUsers_Response_ToProtobuf(resp)
}

func _Encode_CreateComment_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func _Encode_GetComment_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*transport.GetComment_Response)
	return __GetComment_Response_ToProtobuf(resp)
}

func _Encode_GetUserComments_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*transport.GetUserComments_Response)
	return __GetUserComments_Response_ToProtobuf(resp)
}

//========================================= Response Decoders =========================================//

func _Decode_CreateUser_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func _Decode_UpdateUser_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func _Decode_GetUser_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUser_Response)
	return __GetUser_Response_FromProtobuf(resp)
}

func _Decode_FindUsers_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.FindUsers_Response)
	return __FindUsers_Response_FromProtobuf(resp)
}

func _Decode_CreateComment_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func _Decode_GetComment_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetComment_Response)
	return __GetComment_Response_FromProtobuf(resp)
}

func _Decode_GetUserComments_Response(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUserComments_Response)
	return __GetUserComments_Response_FromProtobuf(resp)
}
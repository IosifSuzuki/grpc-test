package service

import (
	"context"
	"gRPC/internal/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	helloMethodName = "Hello"
)

type TestService struct {
	pb.UnimplementedServiceServer
}

func NewTestService() *TestService {
	return &TestService{}
}

func (t *TestService) Test(_ context.Context, request *pb.GetRequestTest) (*pb.ResponseTest, error) {
	switch request.GetText() {
	case helloMethodName:
		return t.helloHandlers()
	default:
		return nil, status.Error(codes.NotFound, "unknown method")
	}
}

func (t *TestService) helloHandlers() (*pb.ResponseTest, error) {
	response := pb.ResponseTest{
		Result: "Word",
	}
	return &response, nil
}

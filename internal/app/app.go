package app

import (
	"gRPC/internal/config"
	"gRPC/internal/proto/pb"
	"gRPC/internal/router"
	"gRPC/internal/service"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func RunHttpServer() error {
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(net.JoinHostPort(conf.GRPCHost, conf.GRPCPort), grpc.WithInsecure())
	if err != nil {
		return err
	}
	grpcClient := pb.NewServiceClient(conn)
	rout := router.NewRouter(grpcClient)
	return http.ListenAndServe(net.JoinHostPort(conf.PublicHost, conf.HTTPPort), rout.PrepareRouter())
}

func RunGrpcServer() error {
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", net.JoinHostPort(conf.PublicHost, conf.GRPCPort))
	if err != nil {
		return err
	}
	grpcService := service.NewTestService()
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, grpcService)
	return server.Serve(listener)
}

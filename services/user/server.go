package backend

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/anvh2/be-blog/common"
	pb "github.com/anvh2/be-blog/grpc-gen/user"
)

// Server ...
type Server struct {
	logger *zap.Logger
}

// NewServer ...
func NewServer() *Server {
	logger, err := common.NewLogger(viper.GetString("users.log_path"))
	if err != nil {
		log.Fatal("failed to new logger production")
	}

	return &Server{
		logger: logger,
	}
}

// Run ...
func (s *Server) Run() error {
	port := viper.GetInt("users.grpc_port")

	server, err := common.NewGrpcServer(port, func(server *grpc.Server) {
		pb.RegisterUserServiceServer(server, s)
	})
	if err != nil {
		s.logger.Fatal("Can't new grpc server", zap.Error(err))
	}

	// server.EnableHTTP(pb.RegisterUserServiceHandlerFromEndpoint, "")

	return server.Run()
}

// Login ...
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Data: &pb.LoginResponse_Data{
			Token: "",
		},
		Error: &pb.Error{
			Code:    1,
			Message: "OK",
		},
	}, nil
}

// Register ...
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		Data: &pb.RegisterResponse_Data{
			Success: true,
		},
		Error: &pb.Error{
			Code:    1,
			Message: "OK",
		},
	}, nil
}

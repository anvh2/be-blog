package backend

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/anvh2/be-blog/common"
	pb "github.com/anvh2/be-blog/grpc-gen/user"
)

// Server ...
type Server struct {
	logger *zap.Logger
	userDb UserDb
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

	server.EnableHTTP(pb.RegisterUserServiceHandlerFromEndpoint, "")
	server.AddShutdownHook(func() {
		s.userDb.Close()
	})
	server.WithHTTPAuthFunc(s.authen, []string{""})

	return server.Run()
}

func (s *Server) authen(r *http.Request) {

}

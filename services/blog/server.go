package backend

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/plugins/storages"

	"google.golang.org/grpc"

	"github.com/anvh2/be-blog/plugins/storages/sqlite"

	"github.com/jinzhu/gorm"
	// include gorm sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	blogDb storages.BlogDb
	proxy  *ReverseProxy
	logger *zap.Logger
}

// NewServer ...
func NewServer() *Server {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{
		viper.GetString("blog.log_path"),
	}
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.MessageKey = "message"

	logger, err := config.Build()
	if err != nil {
		fmt.Println("failed to new logger production")
		// return err
	}

	proxy := NewReverseProxy(logger)

	db, err := gorm.Open("sqlite3", viper.GetString("blog.data_path"))
	if err != nil {
		logger.Error("failed to connection database", zap.Error(err))
		// return err
	}

	blogDb := sqlite.NewBlogDb(db, logger)
	return &Server{
		logger: logger,
		blogDb: blogDb,
		proxy:  proxy,
	}
}

// Run ...
func (s *Server) Run() error {
	port := viper.GetInt("blog.grpc_port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		s.logger.Error("[Run] failed to listen tcp", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, s)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			s.logger.Fatal("[Run] failed to start server", zap.Error(err))
		}
	}()

	go func() {
		err := s.proxy.Run(context.Background())
		if err != nil {
			s.logger.Fatal("[Run] failed to start proxy", zap.Error(err))
		}
	}()

	s.logger.Info("[Run] start listen", zap.Int("port", port))

	sig := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("Shuting down...")
		close(done)
	}()

	fmt.Println("Server is listening\nCtr-c to interup...")
	<-done
	fmt.Println("Shutdown")
	return nil
}

// Login ...
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

// List Blog
func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	if req.Offset < 0 || req.Limit <= 0 {
		s.logger.Error("[List] invalid offset and limit", zap.Int64("offset", req.Offset), zap.Int64("limit", req.Limit))
		return &pb.ListResponse{
			Code:    -1,
			Message: "invalid offset and limit",
		}, nil
	}

	data, err := s.blogDb.List(ctx, req.Offset, req.Limit)
	if err != nil {
		s.logger.Error("[List] failed to get blogs", zap.Error(err))
		return &pb.ListResponse{
			Code:    -1,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[List] get list blogs", zap.Int("count", len(data)))
	return &pb.ListResponse{
		Code:    1,
		Message: "OK",
		Blog:    data,
	}, nil
}

// Get Blog
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.BlogID < 0 {
		s.logger.Error("[Get] failed to get blog", zap.Int64("id", req.BlogID))
		return &pb.GetResponse{
			Code:    -3,
			Message: "invalid id",
		}, nil
	}

	data, err := s.blogDb.Get(ctx, req.BlogID)
	if err != nil {
		s.logger.Error("[Get] failed to get blog", zap.Int64("id", req.BlogID))
		return &pb.GetResponse{
			Code:    -3,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Get] get blog", zap.String("blog", data.String()))
	return &pb.GetResponse{
		Code:    1,
		Message: "OK",
		Blog:    data,
	}, nil
}

// Update Blog
func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	err := s.blogDb.Update(ctx, &pb.BlogData{})
	if err != nil {
		s.logger.Error("[Update] failed to update blog", zap.String("blog", req.String()))
		return &pb.UpdateResponse{
			Code:    -4,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Update] update blog", zap.String("blog", req.String()))
	return &pb.UpdateResponse{
		Code:    1,
		Message: "OK",
	}, nil
}

// Delete Blog
func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if req.BlogID < 0 {
		s.logger.Error("[Delete] failed to delete blog", zap.Int64("id", req.BlogID))
		return &pb.DeleteResponse{
			Code:    -5,
			Message: "invalid id",
		}, nil
	}
	err := s.blogDb.Delete(ctx, req.BlogID)
	if err != nil {
		s.logger.Error("[Delete] failed to delete blog", zap.Int64("id", req.BlogID))
		return &pb.DeleteResponse{
			Code:    -5,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Delete] delete blog", zap.Int64("id", req.BlogID))
	return &pb.DeleteResponse{
		Code:    1,
		Message: "OK",
	}, nil
}

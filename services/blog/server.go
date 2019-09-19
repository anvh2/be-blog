package blog

import (
	"context"
	"fmt"
	"net"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
	pb "github.com/anvh2/z-blogs/grpc-gen/blog"
	"github.com/anvh2/z-blogs/storages"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	blogDb storages.BlogDb
	logger *zap.Logger
}

// NewServer ...
func NewServer(blogDb storages.BlogDb, logger *zap.Logger) *Server {
	return &Server{
		blogDb: blogDb,
		logger: logger,
	}
}

// Run ...
func (s *Server) Run() error {
	port := viper.GetInt("blogs.grpc_port")
	// create a listener on TCP
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		s.logger.Error("[Run] failed to listen tcp", zap.Error(err))
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	pb.RegisterBlogServiceServer(grpcServer, s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		s.logger.Error("[Run] failed to start server", zap.Error(err))
	}
	defer s.logger.Info("[Run] start listen", zap.Int("port", port))

	return nil
}

// List Blog
func (s *Server) List(ctx context.Context, req *blog.ListRequest) (*blog.ListResponse, error) {
	if req.Offset < 0 || req.Limit <= 0 {
		s.logger.Error("[List] invalid offset and limit", zap.Int64("offset", req.Offset), zap.Int64("limit", req.Limit))
		return &blog.ListResponse{
			Code:    -1,
			Message: "invalid offset and limit",
		}, nil
	}

	data, err := s.blogDb.List(ctx, req.Offset, req.Limit)
	if err != nil {
		s.logger.Error("[List] failed to get blogs", zap.Error(err))
		return &blog.ListResponse{
			Code:    -1,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[List] get list blogs", zap.Int("count", len(data)))
	return &blog.ListResponse{
		Code:    1,
		Message: "OK",
		Blogs:   data,
	}, nil
}

// Create Blog
func (s *Server) Create(ctx context.Context, req *blog.BlogData) (*blog.BlogResponse, error) {
	err := s.blogDb.Create(ctx, req)
	if err != nil {
		s.logger.Error("[Create] failed to create blog", zap.String("blog", req.String()))
		return &blog.BlogResponse{
			Code:    -2,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Create] create blog", zap.String("blog", req.String()))
	return &blog.BlogResponse{
		Code:    1,
		Message: "OK",
		Blog:    req,
	}, nil
}

// Get Blog
func (s *Server) Get(ctx context.Context, req *blog.GetRequest) (*blog.BlogResponse, error) {
	if req.Id < 0 {
		s.logger.Error("[Get] failed to get blog", zap.Int64("id", req.Id))
		return &blog.BlogResponse{
			Code:    -3,
			Message: "invalid id",
		}, nil
	}

	data, err := s.blogDb.Get(ctx, req.Id)
	if err != nil {
		s.logger.Error("[Get] failed to get blog", zap.Int64("id", req.Id))
		return &blog.BlogResponse{
			Code:    -3,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Get] get blog", zap.String("blog", data.String()))
	return &blog.BlogResponse{
		Code:    1,
		Message: "OK",
		Blog:    data,
	}, nil
}

// Update Blog
func (s *Server) Update(ctx context.Context, req *blog.BlogData) (*blog.BlogResponse, error) {
	err := s.blogDb.Update(ctx, req)
	if err != nil {
		s.logger.Error("[Update] failed to update blog", zap.String("blog", req.String()))
		return &blog.BlogResponse{
			Code:    -4,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Update] update blog", zap.String("blog", req.String()))
	return &blog.BlogResponse{
		Code:    1,
		Message: "OK",
		Blog:    req,
	}, nil
}

// Delete Blog
func (s *Server) Delete(ctx context.Context, req *blog.DeleteRequest) (*blog.BlogResponse, error) {
	if req.Id < 0 {
		s.logger.Error("[Delete] failed to delete blog", zap.Int64("id", req.Id))
		return &blog.BlogResponse{
			Code:    -5,
			Message: "invalid id",
		}, nil
	}
	err := s.blogDb.Delete(ctx, req.Id)
	if err != nil {
		s.logger.Error("[Delete] failed to delete blog", zap.Int64("id", req.Id))
		return &blog.BlogResponse{
			Code:    -5,
			Message: err.Error(),
		}, nil
	}

	defer s.logger.Info("[Delete] delete blog", zap.Int64("id", req.Id))
	return &blog.BlogResponse{
		Code:    1,
		Message: "OK",
	}, nil
}

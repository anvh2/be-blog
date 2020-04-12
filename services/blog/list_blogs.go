package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"go.uber.org/zap"
)

// List Blog
func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	if req.Offset < 0 || req.Limit <= 0 {
		s.logger.Error("[List] invalid offset and limit", zap.Int64("offset", req.Offset), zap.Int64("limit", req.Limit))
		return &pb.ListResponse{}, nil
	}

	return &pb.ListResponse{}, nil
}

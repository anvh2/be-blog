package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
)

// Remove ...
func (s *Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	if req.BlogID == "" {
		s.logger.Error("[Remove] empty blogID")
		return &pb.RemoveResponse{}, nil
	}

	return &pb.RemoveResponse{}, nil
}

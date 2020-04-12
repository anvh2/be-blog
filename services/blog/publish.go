package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
)

// Publish ...
func (s *Server) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	return &pb.PublishResponse{}, nil
}

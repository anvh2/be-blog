package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
)

// Update ...
func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	return &pb.UpdateResponse{}, nil
}

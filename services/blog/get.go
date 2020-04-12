package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/plugins/errors"
	"go.uber.org/zap"
)

// Get Blog
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.BlogID == "" {
		s.logger.Error("[Get] empty blogID")
		return &pb.GetResponse{
			Error: &pb.Error{
				Code:    errors.EmptyBlogID,
				Message: errors.GetMessage(errors.EmptyBlogID),
			},
		}, nil
	}

	blog, err := s.blogDb.Read(ctx, req.BlogID)
	if err != nil {
		s.logger.Error("[Get] failed to get blog", zap.String("blog_id", req.BlogID), zap.Error(err))
		return &pb.GetResponse{
			Error: &pb.Error{
				Code:    errors.FailedGetBlog,
				Message: errors.GetMessage(errors.FailedGetBlog),
			},
		}, nil
	}

	return &pb.GetResponse{
		Data: &pb.GetResponse_Data{
			Blog: blog,
		},
		Error: &pb.Error{
			Code:    errors.Success,
			Message: errors.GetMessage(errors.Success),
		},
	}, nil
}

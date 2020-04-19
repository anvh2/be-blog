package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/plugins/errors"
	"go.uber.org/zap"
)

// View ...
func (s *Server) View(ctx context.Context, req *pb.ViewRequest) (*pb.ViewResponse, error) {
	if req.BlogID == "" {
		s.logger.Error("[View] empty blogID")
		return &pb.ViewResponse{
			Error: &pb.Error{
				Code:    errors.EmptyBlogID,
				Message: errors.GetMessage(errors.EmptyBlogID),
				Detail:  errors.GetDetail(errors.EmptyBlogID),
			},
		}, nil
	}

	blog, err := s.blogDB.Read(ctx, req.BlogID)
	if err != nil {
		s.logger.Error("[View] failed to get blog", zap.String("blog_id", req.BlogID), zap.Error(err))
		return &pb.ViewResponse{
			Error: &pb.Error{
				Code:    errors.FailedViewBlog,
				Message: errors.GetMessage(errors.FailedViewBlog),
				Detail:  errors.GetDetail(errors.FailedViewBlog),
			},
		}, nil
	}

	if blog.Type == pb.Type_PRIVATE {
		s.logger.Error("[View] permission denied", zap.String("blog_id", req.BlogID), zap.Any("blog", blog))
		return &pb.ViewResponse{
			Error: &pb.Error{
				Code:    errors.ErrorPermissionDenied,
				Message: errors.GetMessage(errors.ErrorPermissionDenied),
				Detail:  errors.GetDetail(errors.ErrorPermissionDenied),
			},
		}, nil
	}

	return &pb.ViewResponse{
		Data: &pb.ViewResponse_Data{
			Blog: blog,
		},
		Error: &pb.Error{
			Code:    errors.Success,
			Message: errors.GetMessage(errors.Success),
			Detail:  errors.GetDetail(errors.Success),
		},
	}, nil
}

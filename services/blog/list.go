package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/plugins/errors"
	"go.uber.org/zap"
)

// List ...
func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	if req.Offset < 0 || req.Limit < 0 || req.Limit > 30 {
		s.logger.Error("[List] invalid offset and limit", zap.Int32("offset", req.Offset), zap.Int32("limit", req.Limit))
		return &pb.ListResponse{
			Error: &pb.Error{
				Code:    errors.InvalidOffsetLimit,
				Message: errors.GetMessage(errors.InvalidOffsetLimit),
				Detail:  errors.GetDetail(errors.InvalidOffsetLimit),
			},
		}, nil
	}

	// default limit
	if req.Limit == 0 {
		req.Limit = 10
	}

	s.logger.Info("[List] list request", zap.Any("req", req))

	blogs, err := s.blogDB.List(ctx, req.Offset, req.Limit)
	if err != nil {
		s.logger.Error("[List] failed to get list blog", zap.Error(err))
		return &pb.ListResponse{
			Error: &pb.Error{
				Code:    errors.FailedGetBlog,
				Message: errors.GetMessage(errors.FailedGetBlog),
				Detail:  errors.GetDetail(errors.FailedGetBlog),
			},
		}, nil
	}

	// for test
	for _, blog := range blogs {
		blog.UserDName = "Hoang An"
	}

	num, err := s.blogDB.GetNumOfBlogs(ctx)
	if err != nil {
		s.logger.Error("[List] failed to get number of blogs", zap.Error(err))
		return &pb.ListResponse{
			Error: &pb.Error{
				Code:    errors.FailedGetNumOfBlogs,
				Message: errors.GetMessage(errors.FailedGetNumOfBlogs),
				Detail:  errors.GetDetail(errors.FailedGetNumOfBlogs),
			},
		}, nil
	}

	return &pb.ListResponse{
		Data: &pb.ListResponse_Data{
			Blogs:      blogs,
			TotalBlogs: num,
		},
		Error: &pb.Error{
			Code:    errors.Success,
			Message: errors.GetMessage(errors.Success),
			Detail:  errors.GetDetail(errors.Success),
		},
	}, nil
}

package backend

import (
	"context"
	"time"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"go.uber.org/zap"
)

// Create Blog
func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if req.Header == "" {
		return &pb.CreateResponse{
			Code:    -1,
			Message: "HEADER_EMPTY",
		}, nil
	} else if req.Subtitle == "" {
		return &pb.CreateResponse{
			Code:    -1,
			Message: "SUBTITLE_EMPTY",
		}, nil
	} else if req.Background == "" {
		return &pb.CreateResponse{
			Code:    -1,
			Message: "BACKGROUND_EMPTY",
		}, nil
	} else if req.Content == "" {
		return &pb.CreateResponse{
			Code:    -1,
			Message: "CONTENT_EMPTY",
		}, nil
	}

	// userID, err := utils.GetUserIDFromCtx(ctx)
	// if err != nil {
	// 	s.logger.Error("[Create] authen failed", zap.Any("ctx", ctx), zap.Error(err))
	// 	return &pb.CreateResponse{
	// 		Code:    -1,
	// 		Message: "AUTH_FAILED",
	// 	}, nil
	// }

	// s.logger.Info("[Create] request create", zap.String("userID", userID), zap.Any("req", req))

	blog := &pb.BlogData{
		Header:     req.Header,
		Subtitle:   req.Subtitle,
		Background: req.Background,
		Content:    req.Content,
		// UserID:     userID,
		Status:     pb.Status_DRAFT,
		CreateTime: time.Now().UnixNano() / 1e6,
		ReadTime:   req.ReadTime,
	}

	err := s.blogDb.Create(ctx, blog)
	if err != nil {
		s.logger.Error("[Create] failed to create blog", zap.String("blog", blog.String()))
		return &pb.CreateResponse{
			Code:    -1,
			Message: "CREATE_FAILED",
		}, nil
	}

	s.logger.Info("[Create] created blog", zap.String("blog", blog.String()))

	return &pb.CreateResponse{
		Code:    1,
		Message: "OK",
	}, nil
}

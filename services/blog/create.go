package backend

import (
	"context"
	"time"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/plugins/errors"
	"github.com/anvh2/be-blog/utils"
	"go.uber.org/zap"
)

// Create ...
func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	/*if req.UserID == "" {
		s.logger.Error("[Create] empty userID")
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.EmptyUsername,
				Message: errors.GetMessage(errors.EmptyUsername),
				Detail: errors.GetDetail(errors.EmptyUsername),
			},
		}, nil
	} else*/if req.Header == "" {
		s.logger.Error("[Create] empty blog header", zap.String("userID", req.UserID))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.EmptyBlogHeader,
				Message: errors.GetMessage(errors.EmptyBlogHeader),
				Detail:  errors.GetDetail(errors.EmptyBlogHeader),
			},
		}, nil
	} else if req.Subtitle == "" {
		s.logger.Error("[Create] empty blog subtitle", zap.String("userID", req.UserID))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.EmptyBlogSubtitle,
				Message: errors.GetMessage(errors.EmptyBlogSubtitle),
				Detail:  errors.GetDetail(errors.EmptyBlogSubtitle),
			},
		}, nil
	} else if req.Background == "" {
		s.logger.Error("[Create] empty blog background", zap.String("userID", req.UserID))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.EmptyBlogBackground,
				Message: errors.GetMessage(errors.EmptyBlogBackground),
				Detail:  errors.GetDetail(errors.EmptyBlogBackground),
			},
		}, nil
	} else if req.Content == "" {
		s.logger.Error("[Create] empty blog content", zap.String("userID", req.UserID))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.EmptyBlogContent,
				Message: errors.GetMessage(errors.EmptyBlogContent),
				Detail:  errors.GetDetail(errors.EmptyBlogContent),
			},
		}, nil
	} else if req.ReadTime <= 0 {
		s.logger.Error("[Create] empty blog read time", zap.String("userID", req.UserID), zap.Int32("read_time", req.ReadTime))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.InvalidReadTime,
				Message: errors.GetMessage(errors.InvalidReadTime),
				Detail:  errors.GetDetail(errors.InvalidReadTime),
			},
		}, nil
	}

	s.logger.Info("[Create] request create", zap.String("userID", req.UserID), zap.Any("req", req))

	blogID, err := s.blogDB.NextBlogID(ctx, utils.TimeToMs(time.Now()))
	if err != nil {
		s.logger.Error("[Create] failed to create blog", zap.String("userID", req.UserID), zap.Error(err))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.FailedGenBlogID,
				Message: errors.GetMessage(errors.FailedGenBlogID),
				Detail:  errors.GetDetail(errors.FailedGenBlogID),
			},
		}, nil
	}

	blog := &pb.BlogData{
		BlogID:     blogID,
		Header:     req.Header,
		Subtitle:   req.Subtitle,
		Background: req.Background,
		Content:    req.Content,
		UserID:     req.UserID,
		Status:     pb.Status_DRAFT,
		Type:       pb.Type_PUBLIC,
		CreateTime: time.Now().UnixNano() / 1e6,
		ReadTime:   req.ReadTime,
	}

	s.logger.Info("[Create] created blog", zap.String("blog", blog.String()))

	err = s.blogDB.Create(ctx, blog)
	if err != nil {
		s.logger.Error("[Create] failed to create blog", zap.String("userID", req.UserID), zap.Error(err))
		return &pb.CreateResponse{
			Error: &pb.Error{
				Code:    errors.FailedCreateBlog,
				Message: errors.GetMessage(errors.FailedCreateBlog),
				Detail:  errors.GetDetail(errors.FailedCreateBlog),
			},
		}, nil
	}

	return &pb.CreateResponse{
		Data: &pb.CreateResponse_Data{
			BlogID: blogID,
		},
		Error: &pb.Error{
			Code:    errors.Success,
			Message: errors.GetMessage(errors.Success),
			Detail:  errors.GetDetail(errors.Success),
		},
	}, nil
}

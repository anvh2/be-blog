package backend

import (
	"context"
	"time"

	pb "github.com/anvh2/be-blog/grpc-gen/user"
	"github.com/anvh2/be-blog/plugins/encode"
	"github.com/anvh2/be-blog/plugins/errors"
	"github.com/anvh2/be-blog/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Register ...
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Username == "" {
		s.logger.Error("[Login] empty username")
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.EmptyUsername,
				Message: errors.GetMessage(errors.EmptyUsername),
			},
		}, nil
	} else if req.Password == "" || req.ConfirmPassword == "" {
		s.logger.Error("[Login] empty password", zap.String("username", req.Username))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.EmptyPassword,
				Message: errors.GetMessage(errors.EmptyPassword),
			},
		}, nil
	} else if req.DName == "" {
		s.logger.Error("[Login] empty display name", zap.String("username", req.Username))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.EmptyDName,
				Message: errors.GetMessage(errors.EmptyDName),
			},
		}, nil
	} else if req.Email == "" {
		s.logger.Error("[Login] empty email", zap.String("username", req.Username))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.EmptyEmail,
				Message: errors.GetMessage(errors.EmptyEmail),
			},
		}, nil
	}

	s.logger.Info("[Login] register request", zap.Any("req", req))

	if req.Password != req.ConfirmPassword {
		s.logger.Error("[Login] password and confirm password not match", zap.String("username", req.Username))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.ErrorPasswordNotMatch,
				Message: errors.GetMessage(errors.ErrorPasswordNotMatch),
			},
		}, nil
	}

	userID, err := s.userDB.NextUserID(ctx, time.Now().UnixNano()/1e6)
	if err != nil {
		s.logger.Error("[Login] failed to generate userID", zap.String("username", req.Username), zap.Error(err))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.FailedGenUserID,
				Message: errors.GetMessage(errors.FailedGenUserID),
			},
		}, nil
	}

	encryptpwd, err := encode.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("[Login] failed to hash password", zap.String("username", req.Username), zap.Error(err))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.FailedHashPassword,
				Message: errors.GetMessage(errors.FailedHashPassword),
			},
		}, nil
	}

	user := &pb.UserData{
		UserID:         userID,
		Username:       req.Username,
		Password:       encryptpwd,
		DName:          req.DName,
		Email:          req.Email,
		Avatar:         req.Avatar,
		Role:           pb.Role_UNKNOWN,
		RegisteredTime: utils.TimeToMs(time.Now()),
		RegisterSource: pb.Source_ORIGIN,
	}

	if req.Avatar == "" {
		user.Avatar = viper.GetString("user_service.default_avatar")
	}

	err = s.userDB.Create(ctx, user)
	if err != nil {
		s.logger.Error("[Login] failed to create new user", zap.String("username", req.Username))
		return &pb.RegisterResponse{
			Error: &pb.Error{
				Code:    errors.FailedRegisterUser,
				Message: errors.GetMessage(errors.FailedRegisterUser),
			},
		}, nil
	}

	return &pb.RegisterResponse{
		Data: &pb.RegisterResponse_Data{
			Success: true,
		},
		Error: &pb.Error{
			Code:    1,
			Message: "OK",
		},
	}, nil
}

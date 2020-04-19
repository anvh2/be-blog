package backend

import (
	"context"
	"time"

	"github.com/anvh2/be-blog/plugins/encode"
	"github.com/anvh2/be-blog/plugins/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	pb "github.com/anvh2/be-blog/grpc-gen/user"
)

// Claims ...
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login ...
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Username == "" {
		s.logger.Error("[Login] Empty username")
		return &pb.LoginResponse{
			Error: &pb.Error{
				Code:    errors.EmptyUsername,
				Message: errors.GetMessage(errors.EmptyUsername),
			},
		}, nil
	} else if req.Password == "" {
		s.logger.Error("[Login] Empty password", zap.String("username", req.Username))
		return &pb.LoginResponse{
			Error: &pb.Error{
				Code:    errors.EmptyPassword,
				Message: errors.GetMessage(errors.EmptyPassword),
			},
		}, nil
	}

	s.logger.Info("[Login] login request", zap.Any("req", req))

	user, err := s.userDB.ReadViaUsername(ctx, req.Username)
	if err != nil {
		s.logger.Error("[Login] failed to read user via username", zap.String("username", req.Username))
		return &pb.LoginResponse{
			Error: &pb.Error{
				Code:    errors.FailedLogin,
				Message: errors.GetMessage(errors.FailedLogin),
			},
		}, nil
	}

	if ok, err := encode.VerifyPassword(req.Password, user.Password); !ok || err != nil {
		s.logger.Error("[Login] invalid password", zap.String("username", req.Username), zap.String("invalid_pass", req.Password),
			zap.Error(err))
		return &pb.LoginResponse{
			Error: &pb.Error{
				Code:    errors.InvalidPassword,
				Message: errors.GetMessage(errors.InvalidPassword),
			},
		}, nil
	}

	token, err := s.signToken(req.Username)
	if err != nil || token == "" {
		s.logger.Error("[Login] failed to sign token", zap.String("username", req.Username), zap.Error(err))
		return &pb.LoginResponse{
			Error: &pb.Error{
				Code:    errors.ErrorSignToken,
				Message: errors.GetMessage(errors.ErrorSignToken),
			},
		}, nil
	}

	return &pb.LoginResponse{
		Data: &pb.LoginResponse_Data{
			UserID:   user.UserID,
			UserName: user.Username,
			DName:    user.DName,
			Avatar:   user.Avatar,
			Role:     pb.Role(user.Role),
			Token:    token,
		},
		Error: &pb.Error{
			Code:    1,
			Message: "OK",
		},
	}, nil
}

func (s *Server) signToken(username string) (string, error) {
	expiration := time.Now().Add(viper.GetDuration("user_service.expire_token_time") * time.Second)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("user_service.jwt_key")))
}

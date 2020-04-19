package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/user"
)

// UserDB ...
type UserDB interface {
	Create(ctx context.Context, user *pb.UserData) error
	Read(ctx context.Context, userID string) (*pb.UserData, error)
	ReadViaUsername(ctx context.Context, username string) (*pb.UserData, error)
	NextUserID(ctx context.Context, createTime int64) (string, error)
	Close()
}

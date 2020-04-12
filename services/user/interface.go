package backend

import (
	"context"

	"github.com/anvh2/be-blog/plugins/storages/mysql"
)

// UserDb ...
type UserDb interface {
	Create(ctx context.Context, user *mysql.UserData) error
	Read(ctx context.Context, userID string) (*mysql.UserData, error)
	ReadViaUsername(ctx context.Context, username string) (*mysql.UserData, error)
	NextUserID(ctx context.Context, createTime int64) (string, error)
	Close()
}

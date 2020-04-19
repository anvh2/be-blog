package database

import (
	"context"

	"github.com/anvh2/be-blog/plugins/storages/mysql"
	"github.com/anvh2/be-blog/plugins/storages/redis"
	"go.uber.org/zap"

	pb "github.com/anvh2/be-blog/grpc-gen/user"
	goredis "github.com/go-redis/redis"
)

// UserDB ...
type UserDB struct {
	logger *zap.Logger
	db     *mysql.UserDB
	cache  *redis.UserDB
}

// NewUserDB ...
func NewUserDB(db *mysql.UserDB, cache *redis.UserDB, logger *zap.Logger) *UserDB {
	return &UserDB{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

// Create ...
func (db *UserDB) Create(ctx context.Context, user *pb.UserData) error {
	db.cache.Set(ctx, user)
	return db.db.Create(ctx, user)
}

// Read ...
func (db *UserDB) Read(ctx context.Context, userID string) (*pb.UserData, error) {
	user, err := db.cache.Get(ctx, userID)
	if err != nil && err == goredis.Nil {
		user, err := db.db.Read(ctx, userID)
		if err == nil {
			db.cache.Set(ctx, user)
		}
		return user, nil
	}
	return user, err
}

// ReadViaUsername ...
func (db *UserDB) ReadViaUsername(ctx context.Context, username string) (*pb.UserData, error) {
	user, err := db.cache.GetViaUsername(ctx, username)
	if err != nil && err == goredis.Nil {
		user, err := db.db.ReadViaUsername(ctx, username)
		if err == nil {
			db.cache.Set(ctx, user)
		}
		return user, nil
	}
	return user, err
}

// NextUserID ...
func (db *UserDB) NextUserID(ctx context.Context, createTime int64) (string, error) {
	return db.db.NextUserID(ctx, createTime)
}

// Close ...
func (db *UserDB) Close() {
	db.db.Close()
	db.cache.Close()
}

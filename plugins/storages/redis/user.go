package redis

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/anvh2/be-blog/grpc-gen/user"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const (
	keyUserData = "blog.user.data.id.%s"
	keyUserID   = "blog.user.data.username.%s"
)

// UserDB ...
type UserDB struct {
	logger *zap.Logger
	db     *redis.Client
}

// NewUserDB ...
func NewUserDB(db *redis.Client, logger *zap.Logger) *UserDB {
	return &UserDB{
		logger: logger,
		db:     db,
	}
}

// Set ...
func (db *UserDB) Set(ctx context.Context, user *pb.UserData) error {
	item, err := json.Marshal(user)
	if err != nil {
		db.logger.Error("[Set] failed to marshal user", zap.Any("user", user), zap.Error(err))
		return err
	}

	key := fmt.Sprintf(keyUserData, user.UserID)
	err = db.db.Set(key, item, 0).Err()
	if err != nil {
		db.logger.Error("[Set] failed to set item", zap.String("user", string(item)), zap.Error(err))
		return err
	}

	keyID := fmt.Sprintf(keyUserID, user.Username)
	err = db.db.Set(keyID, user.UserID, 0).Err()
	if err != nil {
		db.logger.Error("[Set] failed to set userid via username", zap.String("user", string(item)), zap.Error(err))
		return err
	}

	db.logger.Info("[Set] set item", zap.String("user", string(item)))
	return nil
}

// Get ...
func (db *UserDB) Get(ctx context.Context, userID string) (*pb.UserData, error) {
	user := &pb.UserData{}

	key := fmt.Sprintf(keyData, userID)
	item, err := db.db.Get(key).Result()
	if err != nil {
		db.logger.Error("[Get] failed to get user", zap.String("user_id", userID), zap.Error(err))
		return user, err
	}

	err = json.Unmarshal([]byte(item), user)
	if err != nil {
		db.logger.Error("[Get] failed to parse user", zap.String("user_id", userID), zap.Error(err))
		return user, err
	}

	db.logger.Info("[Get] get item", zap.String("user_id", userID), zap.Any("user", user))
	return user, nil
}

// GetViaUsername ...
func (db *UserDB) GetViaUsername(ctx context.Context, username string) (*pb.UserData, error) {
	user := &pb.UserData{}

	key := fmt.Sprintf(keyUserID, username)
	userID, err := db.db.Get(key).Result()
	if err != nil {
		db.logger.Error("[Get] failed to get user id via username", zap.String("username", username), zap.Error(err))
		return user, err
	}

	return db.Get(ctx, userID)
}

// Close ...
func (db *UserDB) Close() {
	db.db.Close()
}

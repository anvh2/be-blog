package mysql

import (
	"context"
	"errors"
	"fmt"

	"github.com/anvh2/be-blog/common"
	pb "github.com/anvh2/be-blog/grpc-gen/user"
	"github.com/anvh2/be-blog/utils"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// UserDB ...
type UserDB struct {
	logger *common.WrappedLogger
	db     *gorm.DB
}

// NewUserDB ...
func NewUserDB(db *gorm.DB, logger *common.WrappedLogger) (*UserDB, error) {
	err := db.AutoMigrate(&pb.UserData{}).Error
	if err != nil {
		return nil, err
	}

	return &UserDB{
		logger: logger,
		db:     db,
	}, nil
}

// Create ...
func (db *UserDB) Create(ctx context.Context, item *pb.UserData) error {
	db.logger.Info("[Create] create item", zap.Any("user", item))
	return db.db.Create(item).Error
}

// Read ...
func (db *UserDB) Read(ctx context.Context, userID string) (*pb.UserData, error) {
	user := pb.UserData{}

	err := db.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		db.logger.Error("[Read] failed to read user", zap.String("user_id", userID), zap.Error(err))
		return nil, err
	}

	db.logger.Info("[Read] read user", zap.String("user_id", userID), zap.Any("user", user))
	return &user, nil
}

// ReadViaUsername ...
func (db *UserDB) ReadViaUsername(ctx context.Context, username string) (*pb.UserData, error) {
	user := pb.UserData{}

	err := db.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		db.logger.Error("[Read] failed to read user", zap.String("username", username), zap.Error(err))
		return nil, err
	}

	db.logger.Info("[Read] read user", zap.String("username", username), zap.Any("user", user))
	return &user, nil
}

// Update ...
func (db *UserDB) Update(ctx context.Context, item *pb.UserData) error {
	db.logger.Info("[Update] update item", zap.Any("user", item))
	return db.db.Save(item).Error
}

// NextUserID ...
func (db *UserDB) NextUserID(ctx context.Context, createTime int64) (string, error) {
	res := struct {
		Counter int64 `gorm:"column:counter"`
	}{}

	format := utils.MsToTime(createTime).Format("0601")
	counterName := "user_" + format

	err := db.db.Raw("CALL nextID(?)", counterName).Scan(&res).Error
	if err != nil {
		db.logger.Error("[NextUserID] failed to get next user counter", zap.String("counter_name", counterName), zap.Error(err))
		return "", err
	}

	if res.Counter == 0 {
		return "", errors.New("failed to get next user counter")
	}

	userID := fmt.Sprintf(patternID, format, res.Counter)

	db.logger.Info("[NextUserID] generate userID", zap.String("user_id", userID))
	return userID, nil
}

// Close ...
func (db *UserDB) Close() {
	db.db.Close()
}

package mysql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/utils"
)

const (
	patternID = "%s%08d"
)

// BlogDB ...
type BlogDB struct {
	logger *zap.Logger
	db     *gorm.DB
}

// NewBlogDB ...
func NewBlogDB(db *gorm.DB, logger *zap.Logger) (*BlogDB, error) {
	err := db.AutoMigrate(&pb.BlogData{}).Error
	if err != nil {
		return nil, err
	}

	return &BlogDB{
		db:     db,
		logger: logger,
	}, nil
}

// Create ...
func (db *BlogDB) Create(ctx context.Context, item *pb.BlogData) error {
	db.logger.Info("[Create] create item", zap.Any("blog", item))
	return db.db.Create(item).Error
}

// Read ...
func (db *BlogDB) Read(ctx context.Context, blogID string) (*pb.BlogData, error) {
	blog := pb.BlogData{}

	err := db.db.Where("blog_id = ?", blogID).First(&blog).Error
	if err != nil {
		db.logger.Error("[Read] failed to read blog", zap.String("blog_id", blogID), zap.Error(err))
		return nil, err
	}

	db.logger.Info("[Read] read blog", zap.String("blog_id", blogID), zap.Any("blog", blog))
	return &blog, nil
}

// Update ...
func (db *BlogDB) Update(ctx context.Context, item *pb.BlogData) error {
	db.logger.Info("[Update] update item", zap.Any("blog", item))
	return db.db.Save(item).Error
}

// List ...
func (db *BlogDB) List(ctx context.Context, offset, limit int32) ([]*pb.BlogData, error) {
	blogs := []*pb.BlogData{}

	err := db.db.Offset(offset).Limit(limit).Find(&blogs).Error
	if err != nil {
		db.logger.Error("[List] failed to list", zap.Error(err))
		return blogs, err
	}

	db.logger.Info("[List] list", zap.Int("num", len(blogs)))
	return blogs, nil
}

// NextBlogID ...
func (db *BlogDB) NextBlogID(ctx context.Context, createTime int64) (string, error) {
	res := struct {
		Counter int64 `gorm:"column:counter"`
	}{}

	format := utils.MsToTime(createTime).Format("0601")
	counterName := "blog_" + format

	err := db.db.Raw("CALL nextID(?)", counterName).Scan(&res).Error
	if err != nil {
		db.logger.Error("[NextBlogID] failed to get next blog counter", zap.String("counter_name", counterName), zap.Error(err))
		return "", err
	}

	if res.Counter == 0 {
		return "", errors.New("failed to get next blog counter")
	}

	blogID := fmt.Sprintf(patternID, format, res.Counter)

	db.logger.Info("[NextBlogID] generate blogID", zap.String("blog_id", blogID))
	return blogID, nil
}

// Close ...
func (db *BlogDB) Close() {
	db.db.Close()
}

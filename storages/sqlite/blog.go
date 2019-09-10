package sqlite

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
	"github.com/jinzhu/gorm"

	// include gorm sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.uber.org/zap"
)

// BlogDb ...
type BlogDb struct {
	db        *gorm.DB
	logger    *zap.Logger
	cacheList []*blog.BlogData
	cacheMap  map[int64]*blog.BlogData
}

// NewBlogDb ...
func NewBlogDb(db *gorm.DB, logger *zap.Logger) *BlogDb {
	db.AutoMigrate(&blog.BlogData{})
	return &BlogDb{
		db:        db,
		logger:    logger,
		cacheList: nil,
		cacheMap:  make(map[int64]*blog.BlogData),
	}
}

// List Blog
func (db *BlogDb) List(ctx context.Context, offset, limit int64) ([]*blog.BlogData, error) {
	if db.cacheList != nil {
		return db.cacheList, nil
	}

	var items []*blog.BlogData
	if err := db.db.Find(&items).Error; err != nil {
		return nil, err
	}

	for _, item := range items {
		fillData(item)
	}
	defer db.logger.Info("get items", zap.Int("count", len(items)))
	defer db.updateCache(items)
	return items, nil
}

// Create Blog
func (db *BlogDb) Create(ctx context.Context, item *blog.BlogData) error {
	defer db.logger.Info("create item", zap.String("item", item.String()))
	defer db.clearCache()
	return db.db.Create(fillData(item)).Error
}

// Get Blog
func (db *BlogDb) Get(ctx context.Context, id int64) (*blog.BlogData, error) {
	if db.cacheList != nil {
		if v, ok := db.cacheMap[id]; ok {
			return v, nil
		}
		return nil, fmt.Errorf("Blog is not found: %d", id)
	}
	var item blog.BlogData
	if err := db.db.First(&item, id).Error; err != nil {
		return nil, err
	}

	defer db.logger.Info("get item", zap.String("item", item.String()))
	return fillData(&item), nil
}

// Update Blog
func (db *BlogDb) Update(ctx context.Context, item *blog.BlogData) error {
	defer db.logger.Info("update item", zap.String("item", item.String()))
	defer db.clearCache()
	return db.db.Save(fillData(item)).Error
}

// Delete Blog
func (db *BlogDb) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("Invalid id: %d", id)
	}
	defer db.logger.Info("delete item", zap.Int64("id", id))
	db.clearCache()
	return db.db.Delete(&blog.BlogData{Id: id}).Error
}

// Close ...
func (db *BlogDb) Close() {
	db.db.Close()
	db.clearCache()
}

func (db *BlogDb) updateCache(items []*blog.BlogData) {
	cacheMap := make(map[int64]*blog.BlogData)
	for _, item := range items {
		cacheMap[item.Id] = item
	}
	db.cacheMap = cacheMap
	db.cacheList = items
}

func (db *BlogDb) clearCache() {
	db.cacheList = nil
}

func fillData(data *blog.BlogData) *blog.BlogData {
	if data.Comments != nil {
		if comments, err := json.Marshal(data.Comments); err == nil {
			data.CommentStr = string(comments)
		}
	} else {
		data.Comments = []*blog.Comment{}
		if data.CommentStr != "" {
			json.Unmarshal([]byte(data.CommentStr), &data.Comments)
		}
	}

	if data.Tags != nil {
		if tags, err := json.Marshal(data.Tags); err == nil {
			data.TagStr = string(tags)
		}
	} else {
		data.Tags = []string{}
		if data.TagStr != "" {
			json.Unmarshal([]byte(data.TagStr), &data.Tags)
		}
	}

	if data.Images != nil {
		if images, err := json.Marshal(data.Images); err == nil {
			data.ImagesStr = string(images)
		}
	} else {
		data.Images = []string{}
		if data.ImagesStr != "" {
			json.Unmarshal([]byte(data.ImagesStr), &data.Images)
		}
	}
	return data
}

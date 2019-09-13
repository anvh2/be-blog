package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
	"github.com/anvh2/z-blogs/plugins/poller"
	"github.com/jinzhu/gorm"

	// include gorm sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.uber.org/zap"
)

// DURATION ...
var DURATION time.Duration = 1

// Item ...
type Item struct {
	blog       *blog.BlogData
	expiration int64
}

func (i *Item) onExpired() bool {
	if i.expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > i.expiration
}

// BlogDb ...
type BlogDb struct {
	db     *gorm.DB
	logger *zap.Logger
	cache  map[int64]Item
	mutex  *sync.RWMutex
}

// NewBlogDb ...
func NewBlogDb(db *gorm.DB, logger *zap.Logger) *BlogDb {
	db.AutoMigrate(&blog.BlogData{})

	return &BlogDb{
		db:     db,
		logger: logger,
		cache:  make(map[int64]Item),
		mutex:  &sync.RWMutex{},
	}
}

func (db *BlogDb) roundRobin() {
	for key, item := range db.cache {
		if item.onExpired() {
			db.DeleteCache(key)
		}
	}
}

// SyncCache ...
func (db *BlogDb) SyncCache(duration time.Duration) {
	poller := poller.NewPoller(db.roundRobin, duration)
	go poller.Run()
}

// SetCache ...
func (db *BlogDb) SetCache(key int64, blog *blog.BlogData, duration time.Duration) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.cache[key] = Item{
		blog:       blog,
		expiration: time.Now().Add(duration).UnixNano(),
	}
}

// GetCache ...
func (db *BlogDb) GetCache(key int64) *blog.BlogData {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if item, ok := db.cache[key]; ok {
		return item.blog
	}
	return nil
}

// DeleteCache ...
func (db *BlogDb) DeleteCache(key int64) {
	delete(db.cache, key)
}

// UpdateCache ...
func (db *BlogDb) UpdateCache(items []*blog.BlogData) {

}

// ClearCache ...
func (db *BlogDb) ClearCache() {

}

// List Blog
func (db *BlogDb) List(ctx context.Context, offset, limit int64) ([]*blog.BlogData, error) {
	// TODO: get cache first

	var items []*blog.BlogData
	if err := db.db.Where("status <> ?", blog.Status_REMOVE).Find(&items).Error; err != nil {
		return nil, err
	}

	for _, item := range items {
		fillData(item)
	}
	defer db.logger.Info("get items", zap.Int("count", len(items)))
	defer db.UpdateCache(items)
	return items, nil
}

// Create Blog
func (db *BlogDb) Create(ctx context.Context, item *blog.BlogData) error {
	defer db.logger.Info("create item", zap.String("item", item.String()))
	defer db.SetCache(item.Id, item, DURATION)
	return db.db.Create(fillData(item)).Error
}

// Get Blog
func (db *BlogDb) Get(ctx context.Context, id int64) (*blog.BlogData, error) {
	data := db.GetCache(id)
	if data != nil {
		return data, nil
	}
	var item blog.BlogData
	if err := db.db.Where("status <> ?", blog.Status_REMOVE).First(&item, id).Error; err != nil {
		return nil, err
	}

	defer db.logger.Info("get item", zap.String("item", item.String()))
	return fillData(&item), nil
}

// Update Blog
func (db *BlogDb) Update(ctx context.Context, item *blog.BlogData) error {
	defer db.logger.Info("update item", zap.String("item", item.String()))
	defer db.SetCache(item.Id, item, DURATION)
	return db.db.Save(fillData(item)).Error
}

// Delete Blog
func (db *BlogDb) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("Invalid id: %d", id)
	}
	defer db.logger.Info("delete item", zap.Int64("id", id))

	db.DeleteCache(id)

	data, err := db.Get(ctx, id)
	if err != nil {
		return err
	}

	data.Status = blog.Status_REMOVE
	return db.db.Save(data).Error
}

// Close ...
func (db *BlogDb) Close() {
	db.db.Close()
	db.ClearCache()
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

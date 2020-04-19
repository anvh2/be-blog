package sqlite

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/anvh2/be-blog/grpc-gen/blog"
// 	"github.com/jinzhu/gorm"

// 	// include gorm sqlite
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// 	"go.uber.org/zap"
// )

// // BlogDB ...
// type BlogDB struct {
// 	db     *gorm.DB
// 	logger *zap.Logger
// }

// // NewBlogDB ...
// func NewBlogDB(db *gorm.DB, logger *zap.Logger) *BlogDB {
// 	db.AutoMigrate(&blog.BlogData{})

// 	return &BlogDB{
// 		db:     db,
// 		logger: logger,
// 	}
// }

// // List Blog
// func (db *BlogDB) List(ctx context.Context, offset, limit int64) ([]*blog.BlogData, error) {
// 	// TODO: get cache first

// 	var items []*blog.BlogData
// 	if err := db.db.Where("status <> ?", blog.Status_REMOVE).Find(&items).Error; err != nil {
// 		return nil, err
// 	}

// 	for _, item := range items {
// 		fillData(item)
// 	}
// 	defer db.logger.Info("get items", zap.Int("count", len(items)))
// 	return items, nil
// }

// // Create Blog
// func (db *BlogDB) Create(ctx context.Context, item *blog.BlogData) error {
// 	defer db.logger.Info("[Create]create item", zap.String("item", item.String()))
// 	return db.db.Create(fillData(item)).Error
// }

// // Get Blog
// func (db *BlogDB) Get(ctx context.Context, id int64) (*blog.BlogData, error) {
// 	var item blog.BlogData
// 	if err := db.db.Where("status <> ?", blog.Status_REMOVE).First(&item, id).Error; err != nil {
// 		return nil, err
// 	}

// 	defer db.logger.Info("get item", zap.String("item", item.String()))
// 	return fillData(&item), nil
// }

// // Update Blog
// func (db *BlogDB) Update(ctx context.Context, item *blog.BlogData) error {
// 	defer db.logger.Info("update item", zap.String("item", item.String()))
// 	return db.db.Save(fillData(item)).Error
// }

// // Delete Blog
// func (db *BlogDB) Delete(ctx context.Context, id int64) error {
// 	if id <= 0 {
// 		return fmt.Errorf("Invalid id: %d", id)
// 	}
// 	defer db.logger.Info("delete item", zap.Int64("id", id))

// 	data, err := db.Get(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	data.Status = blog.Status_REMOVE
// 	return db.db.Save(data).Error
// }

// // Close ...
// func (db *BlogDB) Close() {
// 	db.db.Close()
// }

// func fillData(data *blog.BlogData) *blog.BlogData {
// 	if data.Comments != nil {
// 		if comments, err := json.Marshal(data.Comments); err == nil {
// 			data.CommentStr = string(comments)
// 		}
// 	} else {
// 		data.Comments = []*blog.Comment{}
// 		if data.CommentStr != "" {
// 			json.Unmarshal([]byte(data.CommentStr), &data.Comments)
// 		}
// 	}

// 	if data.Tags != nil {
// 		if tags, err := json.Marshal(data.Tags); err == nil {
// 			data.TagStr = string(tags)
// 		}
// 	} else {
// 		data.Tags = []string{}
// 		if data.TagStr != "" {
// 			json.Unmarshal([]byte(data.TagStr), &data.Tags)
// 		}
// 	}

// 	if data.Images != nil {
// 		if images, err := json.Marshal(data.Images); err == nil {
// 			data.ImagesStr = string(images)
// 		}
// 	} else {
// 		data.Images = []string{}
// 		if data.ImagesStr != "" {
// 			json.Unmarshal([]byte(data.ImagesStr), &data.Images)
// 		}
// 	}
// 	return data
// }

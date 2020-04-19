package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
)

var (
	keyData       = "blog.data.%s"
	keyCreateTime = "blog.create_time" // ensure 30 newest blog will in redis
	keyNumOfBlog  = "blog.num.db"
)

// BlogDB ...
type BlogDB struct {
	logger *zap.Logger
	db     *redis.Client
}

// NewBlogDB ...
func NewBlogDB(db *redis.Client, logger *zap.Logger) *BlogDB {
	return &BlogDB{
		db:     db,
		logger: logger,
	}
}

// Set ...
func (db *BlogDB) Set(ctx context.Context, blog *pb.BlogData) error {
	item, err := json.Marshal(blog)
	if err != nil {
		db.logger.Error("[Set] failed to marshal blog", zap.Any("blog", blog), zap.Error(err))
		return err
	}

	key := fmt.Sprintf(keyData, blog.BlogID)
	err = db.db.Set(key, item, 0).Err()
	if err != nil {
		db.logger.Error("[Set] failed to set item", zap.String("blog", string(item)), zap.Error(err))
		return err
	}

	db.logger.Info("[Set] set item", zap.String("blog", string(item)))
	return nil
}

// Get ...
func (db *BlogDB) Get(ctx context.Context, blogID string) (*pb.BlogData, error) {
	blog := &pb.BlogData{}

	key := fmt.Sprintf(keyData, blogID)
	item, err := db.db.Get(key).Result()
	if err != nil {
		db.logger.Error("[Get] failed to get blog", zap.String("blog_id", blogID), zap.Error(err))
		return blog, err
	}

	err = json.Unmarshal([]byte(item), blog)
	if err != nil {
		db.logger.Error("[Get] failed to parse blog", zap.String("blog_id", blogID), zap.Error(err))
		return blog, err
	}

	db.logger.Info("[Get] get item", zap.String("blog_id", blogID), zap.Any("blog", blog))
	return blog, nil
}

// AddToSortedSet ...
func (db *BlogDB) AddToSortedSet(ctx context.Context, blog *pb.BlogData) (int64, error) {
	num, err := db.db.ZAdd(keyCreateTime, redis.Z{
		Score:  float64(blog.CreateTime),
		Member: blog.BlogID,
	}).Result()
	if err != nil {
		db.logger.Error("[AddToSortedSet] failed to add", zap.String("blog_id", blog.BlogID), zap.Error(err))
		return 0, err
	}

	defer db.logger.Info("[AddToSortedSet] add", zap.String("blog_id", blog.BlogID), zap.Int64("create_time", blog.CreateTime),
		zap.Int64("num", num))
	return num, nil
}

// List ...
func (db *BlogDB) List(ctx context.Context, start, stop int64) ([]*pb.BlogData, []string, error) {
	blogs := []*pb.BlogData{}
	miss := []string{}

	blogIDs, err := db.db.ZRange(keyCreateTime, start, stop).Result()
	if err != nil {
		db.logger.Error("[List] failed to get list blog id", zap.Error(err))
		return blogs, miss, err
	}

	for _, blogID := range blogIDs {
		blog, err := db.Get(ctx, blogID)
		if err != nil {
			miss = append(miss, blogID)
		}

		blogs = append(blogs, blog)
	}

	db.logger.Info("[List] list blog", zap.Int("num", len(blogs)), zap.Any("miss", miss))
	return blogs, miss, nil
}

// CompareNumOfBlogs ...
func (db *BlogDB) CompareNumOfBlogs(ctx context.Context) (bool, error) {
	numInCache, err := db.db.ZCard(keyCreateTime).Result()
	if err != nil {
		db.logger.Error("[CompareNumOfBlogs] failed to get number blog in cache", zap.Error(err))
		return false, err
	}

	numInDB, err := db.db.Get(keyNumOfBlog).Int64()
	if err != nil {
		db.logger.Error("[CompareNumOfBlogs] failed to get number blog in db", zap.Error(err))
		return false, err
	}

	db.logger.Info("[CompareNumOfBlogs] compare", zap.Int64("cache", numInCache), zap.Int64("db", numInDB))
	return numInCache == numInDB, nil
}

// SetNumOfBlogInDB ...
func (db *BlogDB) SetNumOfBlogInDB(ctx context.Context, num int64) error {
	err := db.db.Set(keyNumOfBlog, num, 0).Err()
	if err != nil {
		db.logger.Error("[SetNumOfBlogInDB] failed to set num of blog", zap.Int64("num", num), zap.Error(err))
		return err
	}

	db.logger.Info("[SetNumOfBlogInDB] set", zap.Int64("num", num))
	return nil
}

// IncrNumOfBlogInDB ...
func (db *BlogDB) IncrNumOfBlogInDB(ctx context.Context) (int64, error) {
	num, err := db.db.Incr(keyNumOfBlog).Result()
	if err != nil {
		db.logger.Error("[IncrNumOfBlogInDB] failed to incr", zap.Error(err))
		return 0, err
	}

	defer db.logger.Info("[IncrNumOfBlogInDB] incr", zap.Int64("num", num))
	return num, nil
}

// Close ...
func (db *BlogDB) Close() {
	db.db.Close()
}

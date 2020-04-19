package database

import (
	"context"
	"sort"

	"github.com/anvh2/be-blog/plugins/storages/mysql"
	"github.com/anvh2/be-blog/plugins/storages/redis"
	"go.uber.org/zap"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	goredis "github.com/go-redis/redis"
)

// BlogDB ...
type BlogDB struct {
	logger *zap.Logger
	db     *mysql.BlogDB
	cache  *redis.BlogDB
}

// NewBlogDB ...
func NewBlogDB(db *mysql.BlogDB, cache *redis.BlogDB, logger *zap.Logger) *BlogDB {
	return &BlogDB{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

// Create ..
func (db *BlogDB) Create(ctx context.Context, item *pb.BlogData) error {
	db.cache.Set(ctx, item)
	db.cache.AddToSortedSet(ctx, item)
	return db.db.Create(ctx, item)
}

// Read ...
func (db *BlogDB) Read(ctx context.Context, blogID string) (*pb.BlogData, error) {
	blog, err := db.cache.Get(ctx, blogID)
	if err == goredis.Nil {
		blog, err := db.db.Read(ctx, blogID)
		if err == nil {
			db.cache.Set(ctx, blog)
			db.cache.AddToSortedSet(ctx, blog)
		}

		return blog, err
	}

	return blog, err
}

// Update ...
func (db *BlogDB) Update(ctx context.Context, blog *pb.BlogData) error {
	db.cache.Set(ctx, blog)
	db.cache.AddToSortedSet(ctx, blog)
	return db.db.Update(ctx, blog)
}

// List ...
// TODO: should check blog in cache and db
func (db *BlogDB) List(ctx context.Context, offset, limit int32) ([]*pb.BlogData, error) {
	blogs := []*pb.BlogData{}
	miss := []string{}
	var err error

	blogs, miss, err = db.cache.List(ctx, int64(offset), int64(limit))
	if err != nil {
		return db.db.List(ctx, offset, limit)
	}

	if len(miss) > 0 {
		for _, blogID := range miss {
			blog, err := db.db.Read(ctx, blogID)
			if err == nil {
				db.cache.Set(ctx, blog)
				db.cache.AddToSortedSet(ctx, blog)
			}

			blogs = append(blogs, blog)
		}

		sort.Slice(blogs, func(i, j int) bool {
			return blogs[j].CreateTime < blogs[i].CreateTime
		})
	}

	return blogs, nil
}

// GetNumOfBlogs ...
func (db *BlogDB) GetNumOfBlogs(ctx context.Context) (int32, error) {
	return 0, nil
}

// NextBlogID ...
func (db *BlogDB) NextBlogID(ctx context.Context, createTime int64) (string, error) {
	return db.db.NextBlogID(ctx, createTime)
}

// Close ...
func (db *BlogDB) Close() {
	db.cache.Close()
	db.db.Close()
}

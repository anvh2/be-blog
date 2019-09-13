package storages

import (
	"context"
	"time"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
)

// BlogDb ...
type BlogDb interface {
	// SyncCache
	SyncCache(duration time.Duration)
	// List Blog
	List(ctx context.Context, offset, limit int64) ([]*blog.BlogData, error)
	// Create Blog
	Create(ctx context.Context, item *blog.BlogData) error
	// Get Blog
	Get(ctx context.Context, id int64) (*blog.BlogData, error)
	// Update Blog
	Update(ctx context.Context, item *blog.BlogData) error
	// Delete Blog
	Delete(ctx context.Context, id int64) error
	// Close
	Close()
}

package storages

import (
	"context"

	"github.com/anvh2/be-blog/grpc-gen/blog"
)

// BlogDb ...
type BlogDb interface {
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

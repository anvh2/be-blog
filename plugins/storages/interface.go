package storages

import (
	"context"

	"github.com/anvh2/be-blog/grpc-gen/blog"
)

// BlogDB ...
type BlogDB interface {
	List(ctx context.Context, offset, limit int64) ([]*blog.BlogData, error)
	Create(ctx context.Context, item *blog.BlogData) error
	Get(ctx context.Context, id int64) (*blog.BlogData, error)
	Update(ctx context.Context, item *blog.BlogData) error
	Delete(ctx context.Context, id int64) error
	Close()
}

package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
)

// BlogDb ...
type BlogDb interface {
	Create(ctx context.Context, item *pb.BlogData) error
	Read(ctx context.Context, blogID string) (*pb.BlogData, error)
	NextBlogID(ctx context.Context, createTime int64) (string, error)
	Close()
}

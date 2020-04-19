package backend

import (
	"context"

	pb "github.com/anvh2/be-blog/grpc-gen/blog"
)

// BlogDB ...
type BlogDB interface {
	Create(ctx context.Context, item *pb.BlogData) error
	Read(ctx context.Context, blogID string) (*pb.BlogData, error)
	List(ctx context.Context, offset, limit int32) ([]*pb.BlogData, error)
	GetNumOfBlogs(ctx context.Context) (int32, error)
	NextBlogID(ctx context.Context, createTime int64) (string, error)
	Close()
}

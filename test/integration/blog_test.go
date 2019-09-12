package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var blogClient blog.BlogServiceClient

func GetConnDev(port string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":"+port, opts...)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return conn
}

func TestMain(m *testing.M) {
	// read config
	viper.SetConfigName("z-blogs.dev")
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	port := viper.GetString("blogs.grpc_port")
	blogClient = blog.NewBlogServiceClient(GetConnDev(port))
	os.Exit(m.Run())
}

func TestAPI(t *testing.T) {
	item := &blog.BlogData{
		Comments: []*blog.Comment{
			&blog.Comment{
				Author:  "anvh2",
				Content: "Greate",
			},
		},
		Tags:   []string{"Tech"},
		Images: []string{"/app/demo.png"},
	}

	ctx := context.Background()

	// api create
	c, err := blogClient.Create(ctx, item)
	assert.Nil(t, err)

	// api get
	g, err := blogClient.Get(ctx, &blog.GetRequest{
		Id: c.Blog.Id,
	})
	assert.Nil(t, err)
	assert.Equal(t, c.Blog, g.Blog)

	// api delete
	d, err := blogClient.Delete(ctx, &blog.DeleteRequest{
		Id: c.Blog.Id,
	})
	assert.Nil(t, err)
	assert.Nil(t, d.Blog)

	//api update
	c.Blog.Status = blog.Status_PUBLISH
	u, err := blogClient.Update(ctx, c.Blog)
	assert.Nil(t, err)
	assert.Equal(t, c.Blog.Status, u.Blog.Status)
}

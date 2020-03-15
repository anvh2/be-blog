package integration

import (
	"fmt"
	"os"
	"testing"

	"github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var blogClient blog.BlogServiceClient

func getConnDev(port string) *grpc.ClientConn {
	return nil
}

func getConnLocal(port string) *grpc.ClientConn {
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
	viper.SetConfigName("config.dev")
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	port := viper.GetString("blog.grpc_port")
	blogClient = blog.NewBlogServiceClient(getConnLocal(port))
	os.Exit(m.Run())
}

func TestAPI(t *testing.T) {
	// item := &blog.BlogData{
	// 	Comments: []*blog.Comment{
	// 		&blog.Comment{
	// 			Author:  "anvh2",
	// 			Content: "Greate",
	// 		},
	// 	},
	// 	Tags:   []string{"Tech"},
	// 	Images: []string{"/app/demo.png"},
	// }

	// ctx := context.Background()

	// // api create
	// c, err := blogClient.Create(ctx, item)
	// assert.Nil(t, err)

	// // api get
	// g, err := blogClient.Get(ctx, &blog.GetRequest{
	// 	Id: c.Blog.Id,
	// })
	// assert.Nil(t, err)
	// assert.Equal(t, c.Blog, g.Blog)

	// // api delete
	// d, err := blogClient.Delete(ctx, &blog.DeleteRequest{
	// 	Id: c.Blog.Id,
	// })
	// assert.Nil(t, err)
	// assert.Nil(t, d.Blog)

	// //api update
	// c.Blog.Status = blog.Status_PUBLISH
	// u, err := blogClient.Update(ctx, c.Blog)
	// assert.Nil(t, err)
	// assert.Equal(t, c.Blog.Status, u.Blog.Status)
}

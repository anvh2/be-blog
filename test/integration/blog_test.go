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

}

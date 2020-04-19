package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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

func TestBasicCreateView(t *testing.T) {
	item := &blog.CreateRequest{
		UserID:     "20041600000001",
		Header:     "Clean Code",
		Subtitle:   "How to write smart code",
		Background: "https://smartcodecentral.com/images/bookright.jpg",
		Content:    "You must read clean code book read",
		ReadTime:   1,
	}

	create, err := blogClient.Create(context.Background(), item)
	assert.Nil(t, err)
	fmt.Println("Create OK: ", create)

	view, err := blogClient.View(context.Background(), &blog.ViewRequest{
		BlogID: create.Data.BlogID,
	})
	assert.Nil(t, err)
	fmt.Println("View OK: ", view)
}

func TestView(t *testing.T) {
	view, err := blogClient.View(context.Background(), &blog.ViewRequest{
		BlogID: "2004_00000002",
	})
	assert.Nil(t, err)
	fmt.Println("View OK: ", view)
}

func TestList(t *testing.T) {
	list, err := blogClient.List(context.Background(), &blog.ListRequest{
		Offset: 0,
		Limit:  10,
	})
	assert.Nil(t, err)
	fmt.Println("List OK:", list)
}

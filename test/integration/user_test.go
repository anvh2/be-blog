package integration

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/anvh2/be-blog/grpc-gen/user"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var client = pb.NewUserServiceClient(devConn())

func devConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":55301", opts...)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return conn
}

func TestRegister(t *testing.T) {
	register, err := client.Register(context.Background(), &pb.RegisterRequest{
		Username:        "anvh2",
		Password:        "Hoangan2110",
		ConfirmPassword: "Hoangan2110",
		DName:           "Hoang An",
		Avatar:          "",
		Email:           "anvo.ht209@gmail.com",
	})
	assert.Nil(t, err)
	fmt.Println("Register OK:", register)
}

func TestLogin(t *testing.T) {
	login, err := client.Login(context.Background(), &pb.LoginRequest{
		Username: "anvh2",
		Password: "Hoangan2110",
	})
	assert.Nil(t, err)
	fmt.Println("Login OK:", login)
}

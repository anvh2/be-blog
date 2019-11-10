package proxy

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	gw "github.com/anvh2/be-blog/grpc-gen/blog"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:55101", "grpc server endpoint")
)

// ReverseProxy -
type ReverseProxy struct {
	logger *zap.Logger
}

// NewReverseProxy -
func NewReverseProxy(logger *zap.Logger) *ReverseProxy {
	return &ReverseProxy{
		logger: logger,
	}
}

// Run -
func (p *ReverseProxy) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// register grpc server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// get proxy port
	grpcPort := viper.GetInt("blogs.grpc_port")
	endpoint := fmt.Sprintf(":%d", grpcPort)

	// register grpc server to proxy
	err := gw.RegisterBlogerviceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	// start proxy
	addr := fmt.Sprintf(":%d", viper.GetInt("blogs.proxy_port"))
	if err := http.ListenAndServe(addr, mux); err != nil {
		p.logger.Error("[Run] failed to start proxy", zap.String("address", addr))
		return err
	}
	defer p.logger.Info("[Run] start proxy", zap.String("address", addr))
	return nil
}

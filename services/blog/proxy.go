package backend

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	gw "github.com/anvh2/be-blog/grpc-gen/blog"
	"github.com/anvh2/be-blog/plugins/middlewares"
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

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcPort := viper.GetInt("blog.grpc_port")
	endpoint := fmt.Sprintf(":%d", grpcPort)

	if err := gw.RegisterBlogServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}

	httpPort := viper.GetInt("blog.proxy_port")
	addr := fmt.Sprintf(":%d", httpPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: middlewares.AllowCORS(mux),
	}

	p.logger.Info("[Run] start proxy", zap.Int("port", httpPort))

	if err := srv.ListenAndServe(); err != nil {
		p.logger.Error("[Run] failed to start proxy", zap.Int("port", httpPort))
		return err
	}

	return nil
}

func addLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		ctx := r.Context()
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

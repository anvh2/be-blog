package common

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/anvh2/be-blog/plugins/middlewares"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// GRPCRegister ...
type GRPCRegister func(*grpc.Server)

// HTTPRegister ...
type HTTPRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

// GRPCServer ...
type GRPCServer struct {
	server       *grpc.Server
	listener     net.Listener
	port         int
	logger       *zap.Logger
	rootPath     string
	grpcRegister GRPCRegister
	httpRegister HTTPRegister
	wgServe      *sync.WaitGroup
}

// NewGrpcServer ...
func NewGrpcServer(port int, grpcRegister GRPCRegister) (*GRPCServer, error) {
	return &GRPCServer{
		port:         port,
		grpcRegister: grpcRegister,
		wgServe:      &sync.WaitGroup{},
	}, nil
}

// Run ...
func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	s.listener = lis

	if s.httpRegister != nil {
		m := cmux.New(s.listener)
		grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
		httpListener := m.Match(cmux.HTTP1Fast())

		s.wgServe.Add(1)
		go func() error {
			defer s.wgServe.Done()
			return s.grpcServe(grpcListener)
		}()

		s.wgServe.Add(1)
		go func() error {
			defer s.wgServe.Done()
			return s.httpServe(httpListener)
		}()

		s.wgServe.Add(1)
		go func() error {
			defer s.wgServe.Done()
			return m.Serve()
		}()
	} else {
		s.wgServe.Add(1)
		go func() error {
			defer s.wgServe.Done()
			return s.grpcServe(s.listener)
		}()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan error, 1) // buffer channel to prevent deadlock

	fmt.Println("Server is listening")

	go func() {
		<-sig
		fmt.Println("Shuting down server ...")
		err := s.listener.Close()
		s.wgServe.Wait()
		done <- err
	}()

	fmt.Println("Ctrl-C to interrup ...")
	err = <-done
	fmt.Println("Server is shutdown")

	return err
}

func (s *GRPCServer) grpcServe(l net.Listener) error {
	interceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(),
	))

	server := grpc.NewServer(interceptors)
	s.grpcRegister(server)

	return server.Serve(l)
}

func (s *GRPCServer) httpServe(l net.Listener) error {
	ctx := context.Background()

	mux := runtime.NewServeMux()
	endpoint := fmt.Sprintf(":%d", s.port)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := s.httpRegister(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	httpPort := viper.GetInt("blog.proxy_port")
	addr := fmt.Sprintf(":%d", httpPort)
	server := &http.Server{
		Addr:    addr,
		Handler: middlewares.AllowCORS(mux),
	}

	return server.Serve(l)
}

// EnableHTTP ...
func (s *GRPCServer) EnableHTTP(httpRegister HTTPRegister, rootPath string) *GRPCServer {
	s.httpRegister = httpRegister
	s.rootPath = rootPath
	return s
}

// AddShutdownHook ...
func (s *GRPCServer) AddShutdownHook(fn func()) {

}

func (s *GRPCServer) runShutdownHook() {

}

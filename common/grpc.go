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
	"google.golang.org/grpc"
)

// GRPCRegister ...
type GRPCRegister func(*grpc.Server)

// HTTPRegister ...
type HTTPRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type shutdownHook func()

// GRPCServer ...
type GRPCServer struct {
	server         *grpc.Server
	listener       net.Listener
	port           int
	logger         *WrappedLogger
	rootPath       string
	grpcRegister   GRPCRegister
	httpRegister   HTTPRegister
	hooks          []shutdownHook
	httpAuth       func(r *http.Request)
	excludeMethods []string
}

// NewGrpcServer ...
func NewGrpcServer(port int, grpcRegister GRPCRegister) (*GRPCServer, error) {
	return &GRPCServer{
		port:         port,
		grpcRegister: grpcRegister,
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
		go s.dualServe(s.listener)
	} else {
		go s.grpcServe(s.listener)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan error, 1) // buffer channel to prevent deadlock

	fmt.Println("Server is listening")

	go func() {
		<-sig
		fmt.Println("Shuting down server ...")
		err = s.runHooks()
		done <- err
	}()

	fmt.Println("Ctrl-C to interrup ...")
	err = <-done
	fmt.Println("Server is shutdown")

	return err
}

func (s *GRPCServer) dualServe(l net.Listener) error {
	wg := &sync.WaitGroup{}
	m := cmux.New(l)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	wg.Add(1)
	go func() error {
		defer wg.Done()
		return s.grpcServe(grpcListener)
	}()

	wg.Add(1)
	go func() error {
		defer wg.Done()
		return s.httpServe(httpListener)
	}()

	wg.Add(1)
	go func() error {
		defer wg.Done()
		return m.Serve()
	}()

	wg.Wait()

	return nil
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

	handler := middlewares.AllowCORS(mux)
	handler = middlewares.RecoverHTTPServer(handler)
	handler = middlewares.AuthenHTTPServer(handler, s.httpAuth, s.excludeMethods)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: handler,
	}

	return server.Serve(l)
}

// EnableHTTP ...
func (s *GRPCServer) EnableHTTP(httpRegister HTTPRegister, rootPath string) {
	s.httpRegister = httpRegister
	s.rootPath = rootPath
}

// AddShutdownHook ...
func (s *GRPCServer) AddShutdownHook(fn func()) {
	if fn == nil {
		return
	}

	s.hooks = append(s.hooks, fn)
}

// WithLogger ...
func (s *GRPCServer) WithLogger(logger *WrappedLogger) {
	s.logger = logger
}

// WithHTTPAuthFunc ...
func (s *GRPCServer) WithHTTPAuthFunc(auth func(r *http.Request), excludeMethods []string) {
	s.httpAuth = auth
	s.excludeMethods = excludeMethods
}

func (s *GRPCServer) runHooks() error {
	for _, hook := range s.hooks {
		hook()
	}

	return s.listener.Close()
}

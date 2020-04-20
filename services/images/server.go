package backend

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/anvh2/be-blog/common"
	"github.com/anvh2/be-blog/plugins/middlewares"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	router     *mux.Router
	logger     *common.WrappedLogger
	dirStorage string
}

// NewServer ...
func NewServer() *Server {
	router := mux.NewRouter()

	logger := &zap.Logger{}
	var err error
	if viper.GetString("app.env") == "staging" || viper.GetString("app.env") == "development" {
		fmt.Println("Create new development logger")
		logger, err = common.NewDevelopmentZapLogger()
		if err != nil {
			fmt.Println("Failed to create new development logger")
		}
	} else {
		logger, err = common.NewProductionZapLogger(viper.GetString("blog.log_path"))
		if err != nil {
			log.Fatal("failed to new logger production\n", err)

		}
	}

	wlogger := common.NewWrappedLogger(logger)

	return &Server{
		router:     router,
		logger:     wlogger,
		dirStorage: viper.GetString("images.dir_storage"),
	}
}

// Run ...
func (s *Server) Run() {
	s.setupRouter()
	port := viper.GetInt("images.http_port")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middlewares.AllowCORS(s.router),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			s.logger.Fatal("[Run] failed to start server")
		}
	}()

	s.logger.Info("[Run] Now server is listening ...", zap.Int("port", port))

	sig := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("Shuting down...")
		close(done)
	}()

	fmt.Println("Server is listening\nCtrl-C to interup ...")
	<-done
	fmt.Println("Shutdown")
}

func (s *Server) setupRouter() {
	s.router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/",
			http.FileServer(FileSystem{http.Dir("./public/images")})))

	s.router.HandleFunc("/v1/upload", s.upload)
	s.router.HandleFunc("/v1/file", s.download)
	s.router.HandleFunc("/v1/remove", s.remove)
	s.router.HandleFunc("/v1/move", s.move)
}

// FileSystem override method Open to custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		return nil, errors.New("File only")
	}

	return f, nil
}

package backend

import (
	"fmt"
	"net/http"

	"github.com/anvh2/be-blog/common"
	pb "github.com/anvh2/be-blog/grpc-gen/blog"

	"google.golang.org/grpc"

	"github.com/anvh2/be-blog/plugins/storages/sqlite"

	"github.com/jinzhu/gorm"
	// include gorm sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	blogDb BlogDb
	logger *zap.Logger
}

// NewServer ...
func NewServer() *Server {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{
		viper.GetString("blog.log_path"),
	}
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.MessageKey = "message"

	logger, err := config.Build()
	if err != nil {
		fmt.Println("failed to new logger production")
	}

	db, err := gorm.Open("sqlite3", viper.GetString("blog.data_path"))
	if err != nil {
		logger.Error("failed to connection database", zap.Error(err))
		// return err
	}

	sqlite.NewBlogDb(db, logger)
	return &Server{
		logger: logger,
	}
}

// Run ...
func (s *Server) Run() error {
	port := viper.GetInt("blog.grpc_port")
	server, err := common.NewGrpcServer(port, func(server *grpc.Server) {
		pb.RegisterBlogServiceServer(server, s)
	})
	if err != nil {
		return err
	}

	server.EnableHTTP(pb.RegisterBlogServiceHandlerFromEndpoint, "")
	server.AddShutdownHook(func() {
		s.blogDb.Close()
	})
	server.WithHTTPAuthFunc(s.authen, []string{""})

	return server.Run()
}

func (s *Server) authen(r *http.Request) {

}

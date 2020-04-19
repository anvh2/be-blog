package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anvh2/be-blog/common"
	pb "github.com/anvh2/be-blog/grpc-gen/blog"
	goredis "github.com/go-redis/redis"

	"google.golang.org/grpc"

	"github.com/anvh2/be-blog/plugins/storages/database"
	"github.com/anvh2/be-blog/plugins/storages/mysql"
	"github.com/anvh2/be-blog/plugins/storages/redis"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	blogDB BlogDB
	logger *zap.Logger
}

// NewServer ...
func NewServer() *Server {
	logger, err := common.NewLogger(viper.GetString("blog.log_path"))
	if err != nil {
		if viper.GetString("app.env") == "staging" {
			fmt.Println("Create new development logger")
			logger, err = zap.NewDevelopment()
		} else {
			log.Fatal("failed to new logger production\n", err)
		}
	}

	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&multiStatements=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.db_name"))
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		logger.Fatal("failed to connection database", zap.Error(err))
	}

	storageBlogDB, err := mysql.NewBlogDB(db, logger)
	if err != nil {
		logger.Fatal("failed to init storage blogDB", zap.Error(err))
	}

	redisCli := goredis.NewClient(&goredis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.pass"),
		MaxRetries: viper.GetInt("redis.max_retries"),
	})

	if err := redisCli.Ping().Err(); err != nil {
		logger.Fatal("failed to connect redis", zap.Error(err))
	}

	cacheBlogDB := redis.NewBlogDB(redisCli, logger)

	blogDb := database.NewBlogDB(storageBlogDB, cacheBlogDB, logger)

	return &Server{
		logger: logger,
		blogDB: blogDb,
	}
}

// Run ...
func (s *Server) Run() error {
	if viper.GetString("app.env") == "staging" {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered", r)
			}
		}()
	}

	port := viper.GetInt("blog.grpc_port")
	server, err := common.NewGrpcServer(port, func(server *grpc.Server) {
		pb.RegisterBlogServiceServer(server, s)
	})
	if err != nil {
		return err
	}

	server.EnableHTTP(pb.RegisterBlogServiceHandlerFromEndpoint, "")
	server.AddShutdownHook(func() {
		s.blogDB.Close()
	})
	server.WithHTTPAuthFunc(s.authen, []string{""})

	return server.Run()
}

func (s *Server) authen(r *http.Request) {

}

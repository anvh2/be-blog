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
	logger *common.WrappedLogger
}

// NewServer ...
func NewServer() *Server {
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

	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&multiStatements=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.db_name"))
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		wlogger.Fatal("failed to connection database", zap.Error(err))
	}

	storageBlogDB, err := mysql.NewBlogDB(db, wlogger)
	if err != nil {
		wlogger.Fatal("failed to init storage blogDB", zap.Error(err))
	}

	redisCli := goredis.NewClient(&goredis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.pass"),
		MaxRetries: viper.GetInt("redis.max_retries"),
	})

	if err := redisCli.Ping().Err(); err != nil {
		wlogger.Fatal("failed to connect redis", zap.Error(err))
	}

	cacheBlogDB := redis.NewBlogDB(redisCli, wlogger)

	blogDb := database.NewBlogDB(storageBlogDB, cacheBlogDB, wlogger)

	return &Server{
		logger: wlogger,
		blogDB: blogDb,
	}
}

// Run ...
func (s *Server) Run() error {
	if viper.GetString("app.env") == "staging" {
		defer func() {
			fmt.Println("Crash")
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

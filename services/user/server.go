package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/anvh2/be-blog/common"
	pb "github.com/anvh2/be-blog/grpc-gen/user"
	"github.com/anvh2/be-blog/plugins/storages/database"
	"github.com/anvh2/be-blog/plugins/storages/mysql"
	"github.com/anvh2/be-blog/plugins/storages/redis"
	goredis "github.com/go-redis/redis"
)

// Server ...
type Server struct {
	logger *zap.Logger
	userDB UserDB
}

// NewServer ...
func NewServer() *Server {
	logger, err := common.NewLogger(viper.GetString("user_service.log_path"))
	if err != nil {
		log.Fatal("failed to new logger production\n", err)
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

	storageUserDB, err := mysql.NewUserDB(db, logger)
	if err != nil {
		logger.Fatal("failed to init storage blogDB", zap.Error(err))
	}

	redisCli := goredis.NewClient(&goredis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   "",
		MaxRetries: viper.GetInt("redis.max_retries"),
	})

	if err := redisCli.Ping().Err(); err != nil {
		logger.Fatal("failed to connect redis", zap.Error(err))
	}

	cacheUserDB := redis.NewUserDB(redisCli, logger)

	userDB := database.NewUserDB(storageUserDB, cacheUserDB, logger)

	return &Server{
		logger: logger,
		userDB: userDB,
	}
}

// Run ...
func (s *Server) Run() error {
	port := viper.GetInt("user_service.grpc_port")

	server, err := common.NewGrpcServer(port, func(server *grpc.Server) {
		pb.RegisterUserServiceServer(server, s)
	})
	if err != nil {
		s.logger.Fatal("Can't new grpc server", zap.Error(err))
	}

	server.EnableHTTP(pb.RegisterUserServiceHandlerFromEndpoint, "")
	server.AddShutdownHook(func() {
		s.userDB.Close()
	})
	server.WithHTTPAuthFunc(s.authen, []string{""})

	return server.Run()
}

func (s *Server) authen(r *http.Request) {

}

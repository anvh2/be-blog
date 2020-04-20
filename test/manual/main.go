package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Table ...
type Table struct {
	ID   int64  `json:"id" gorm:"primary_key;auto_increament"`
	Name string `json:"name"`
}

func main() {
	viper.SetConfigName("config.stg")
	viper.AddConfigPath("../..") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&multiStatements=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.db_name"))
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		log.Fatal("failed to connection database", zap.Error(err))
	}

	err = db.AutoMigrate(&Table{}).Error
	if err != nil {
		log.Fatal(err)
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.pass"),
		MaxRetries: viper.GetInt("redis.max_retries"),
	})

	if err := redisCli.Ping().Err(); err != nil {
		log.Fatal("failed to connect redis", zap.Error(err))
	}

	err = redisCli.Set("anvh2", "Hoang An", 0).Err()
	if err != nil {
		log.Println("Set failed", err)
	} else {
		log.Println("Set ok")
	}

	item, err := redisCli.Get("anvh2").Result()
	if err != nil {
		log.Println("Get failed", err)
	}

	log.Println("Item:", item)
}

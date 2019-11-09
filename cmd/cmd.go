package cmd

import (
	"fmt"

	"github.com/anvh2/be-blog/services/blog"
	"github.com/anvh2/be-blog/storages/sqlite"

	"github.com/jinzhu/gorm"
	// include gorm sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var blogCmd = &cobra.Command{
	Use:   "blogs",
	Short: "Starts blogs BlogService",
	Long:  `Starts blogs BlogService`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := zap.NewProductionConfig()

		config.OutputPaths = []string{
			viper.GetString("blogs.log_path"),
		}
		config.EncoderConfig.LevelKey = "level"
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.CallerKey = "caller"
		config.EncoderConfig.MessageKey = "message"

		logger, err := config.Build()

		if err != nil {
			fmt.Println("failed to new logger production")
			return err
		}

		db, err := gorm.Open("sqlite3", viper.GetString("blogs.data_path"))
		if err != nil {
			logger.Error("failed to connection database", zap.Error(err))
			return err
		}

		blogDb := sqlite.NewBlogDb(db, logger)
		server := blog.NewServer(blogDb, logger)
		return server.Run()
	},
}

func init() {
	RootCmd.AddCommand(blogCmd)
}

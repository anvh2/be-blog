package cmd

import (
	"fmt"

	"github.com/anvh2/z-blogs/services/blog"
	"github.com/anvh2/z-blogs/storages/sqlite"
	"github.com/jinzhu/gorm"

	// include gorm sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var blogCmd = &cobra.Command{
	Use:   "blogs",
	Short: "Starts blogs BlogService",
	Long:  `Starts blogs BlogService`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := zap.NewProduction()
		if err != nil {
			fmt.Println("failed to new logger production")
			return err
		}

		db, err := gorm.Open("sqlite3", "data/blogs.db")
		if err != nil {
			logger.Error("failed to connection database", zap.Error(err))
			return err
		}

		blogDb := sqlite.NewBlogDb(db, logger)
		server := blog.NewServer(blogDb, logger)
		return server.Run()
	},
}

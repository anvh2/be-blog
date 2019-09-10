package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var blogCmd = &cobra.Command{
	Use:   "blogs",
	Short: "Starts blogs BlogService",
	Long:  `Starts blogs BlogService`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := zap.NewProduction()
		if err != nil {
			return err
		}
		return nil
	},
}

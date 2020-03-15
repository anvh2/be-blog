package cmd

import (
	backend "github.com/anvh2/be-blog/services/blog"

	"github.com/spf13/cobra"
)

var blogCmd = &cobra.Command{
	Use:   "blog",
	Short: "Starts blogs BlogService",
	Long:  `Starts blogs BlogService`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := backend.NewServer()
		return server.Run()

	},
}

func init() {
	RootCmd.AddCommand(blogCmd)
}

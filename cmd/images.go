package cmd

import (
	backend "github.com/anvh2/be-blog/services/images"
	"github.com/spf13/cobra"
)

var startCmd = cobra.Command{
	Use:   "images",
	Short: "Start image storage server",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := backend.NewServer()
		server.Run()
		return nil
	},
}

func init() {
	RootCmd.AddCommand(&startCmd)
}

package cmd

import (
	backend "github.com/anvh2/be-blog/services/user"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Starts blogs UserService",
	Long:  `Starts blogs UserService`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := backend.NewServer()
		return server.Run()
	},
}

func init() {
	RootCmd.AddCommand(userCmd)
}

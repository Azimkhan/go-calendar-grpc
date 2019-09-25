package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar microservice demo",
}

func init() {
	RootCmd.AddCommand(grpcServerCmd)
	RootCmd.AddCommand(grpcClientCmd)
}

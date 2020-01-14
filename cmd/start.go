package cmd

import (
	"conscience-backend/server"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, args []string) {
	s := server.New()
	s.Start()
}

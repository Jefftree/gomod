package gomod

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Test() {
	cmd := &cobra.Command{
		Use:  "proxy",
		Long: `A gRPC proxy server, receives requests from the API server and forwards to the agent.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("")
		},
	}
	_ = cmd

}


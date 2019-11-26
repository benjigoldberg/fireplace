package server

import (
	"github.com/benjigoldberg/fireplace/pkg/server"
	"github.com/spf13/cobra"
	shHTTP "github.com/spothero/tools/http"
	"github.com/spothero/tools/log"
)

const longDescription = `
Runs a Server for controlling the fireplace
`

// NewCmd constructs a cobra command for running a spotbot server
func NewCmd(gitSHA string) *cobra.Command {
	c := shHTTP.NewDefaultConfig("fireplace")
	c.RegisterHandlers = server.RegisterMuxes
	lc := &log.Config{UseDevelopmentLogger: true}
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Runs a Server for controlling the fireplace",
		Long:  longDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := lc.InitializeLogger(); err != nil {
				return err
			}
			c.NewServer().Run()
			return nil
		},
	}
	// Server Config
	flags := cmd.Flags()
	c.RegisterFlags(flags)
	lc.RegisterFlags(flags)
	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hirokit/devtool"
)

var (
	devCmd = &cobra.Command{
		Use:     "dev",
		Aliases: []string{"d"},
		Short:   "Run development",
		Long:    "",
		Run: func(cmd *cobra.Command, args []string) {
			devtool.Serve()
		},
	}
)

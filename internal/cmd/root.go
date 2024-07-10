package cmd

import (
	"os"
	
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hrx/internal/log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hrx",
		Short: "Hirokit CLI",
		Long:  "Simplified manual operations",
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(migrateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

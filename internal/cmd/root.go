package cmd

import (
	"errors"
	"os"
	
	"github.com/spf13/cobra"
	
	"github.com/daarlabs/hrx/internal/config"
	"github.com/daarlabs/hrx/internal/flag"
	"github.com/daarlabs/hrx/internal/log"
	"github.com/daarlabs/hrx/internal/message"
)

var (
	rootCmd = &cobra.Command{
		Use:   message.RootUse,
		Short: message.RootShort,
		Long:  message.RootLong,
		Run: func(cmd *cobra.Command, args []string) {
			if config.Config.Generate && len(args) == 0 {
				log.Error(errors.New(message.InvalidFilePath))
				return
			}
			if config.Config.Generate && len(args) > 0 {
				if err := generate(args[0]); err != nil {
					log.Error(err)
					os.Exit(1)
				}
				return
			}
			if err := cmd.Help(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&config.Config.Generate, flag.LongGenerateFlag, flag.ShortGenerateFlag, false,
		message.Generate,
	)
	rootCmd.PersistentFlags().BoolVarP(
		&config.Config.Handler, flag.LongHandlerFlag, flag.ShortHandlerFlag, false,
		message.GenerateHandler,
	)
	rootCmd.PersistentFlags().BoolVarP(
		&config.Config.Component, flag.LongComponentFlag, flag.ShortComponentFlag, false,
		message.GenerateComponent,
	)
	rootCmd.PersistentFlags().BoolVarP(
		&config.Config.Form, flag.LongFormFlag, flag.ShortFormFlag, false,
		message.GenerateForm,
	)
	rootCmd.PersistentFlags().BoolVarP(
		&config.Config.Page, flag.LongPageFlag, flag.ShortPageFlag, false,
		message.GeneratePage,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

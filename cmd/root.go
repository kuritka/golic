//package provides cobra commands
package cmd

import (
	"context"
	"os"

	"github.com/kuritka/golic/utils/log"

	"github.com/spf13/cobra"
)

var Verbose bool

var ctx = context.Background()

var logger = log.Log

var rootCmd = &cobra.Command{
	Short: "golic license injector",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Error().Msg("no parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msg("done")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

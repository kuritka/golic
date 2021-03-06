package cmd

import (
	"github.com/kuritka/golic/impl/inject"
	"github.com/spf13/cobra"
	"net/url"
	"os"
)

var injectOptions inject.Options

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(injectOptions.License); os.IsNotExist(err) {
			logger.Error().Msgf("invalid license path '%s'",injectOptions.License)
			_ = cmd.Help()
			os.Exit(0)
		}
		if _, err := os.Stat(injectOptions.Template); os.IsNotExist(err) {
			logger.Error().Msgf("missing template path '%s'",injectOptions.Template)
			_ = cmd.Help()
			os.Exit(0)
		}
		if _,err := url.Parse(injectOptions.ConfigURL); err != nil {
			logger.Error().Msgf("invalid config.yaml url '%s'",injectOptions.ConfigURL)
			_ = cmd.Help()
			os.Exit(0)
		}
		i := inject.New(ctx, injectOptions)
		Command(i).MustRun()
	},
}

func init() {
	injectCmd.Flags().StringVarP(&injectOptions.License, "licignore", "l", "", ".licignore path")
	injectCmd.Flags().StringVarP(&injectOptions.Template, "template", "t", "", "license template path")
	injectCmd.Flags().StringVarP(&injectOptions.Copyright, "copyright", "c", "",
		"e.g.: Copyright 2021 Absa Group Limited")
	injectCmd.Flags().BoolVarP(&injectOptions.Dry, "dry", "d", false, "dry run")
	injectCmd.Flags().StringVarP(&injectOptions.ConfigURL, "config-url", "u", "https://github.com/kuritka/golic/config.yaml", "config URL")
	rootCmd.AddCommand(injectCmd)
}

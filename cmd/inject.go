package cmd

import (
	"github.com/kuritka/golic/impl/inject"
	"github.com/spf13/cobra"
)

var injectOptions inject.Options

var injectCmd = &cobra.Command{
	Use:   "run",
	Short: "traverse from path and inject license",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		i := inject.New(ctx, injectOptions)
		Command(i).MustRun()

	},
}

func init() {
	injectCmd.Flags().StringVarP(&injectOptions.License, "license", "l", "", "License file path")
	rootCmd.AddCommand(injectCmd)
}

package cmd

import (
	"github.com/kuritka/golic/impl/inject"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var injectOptions inject.Options

var injectCmd = &cobra.Command{
	Use:   "run",
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

		i := inject.New(ctx, injectOptions)
		Command(i).MustRun()

	},
}

func init() {
	injectCmd.Flags().StringVarP(&injectOptions.License, "licignore", "l", "", ".licignore path")
	injectCmd.Flags().StringVarP(&injectOptions.Template, "template", "t", "", "license template path")
	injectCmd.Flags().IntVarP(&injectOptions.Year, "year", "y", time.Now().Year(), "year")
	injectCmd.Flags().StringVarP(&injectOptions.Owner, "owner", "o", "",
		"copyright owner or entity authorized by the copyright owner that is granting the License.")
	injectCmd.Flags().BoolVarP(&injectOptions.Dry, "dry", "d", false, "dry run")
	rootCmd.AddCommand(injectCmd)
}

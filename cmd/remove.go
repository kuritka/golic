package cmd

import (
	"github.com/kuritka/golic/impl/remove"
	"github.com/spf13/cobra"
)

var removeOptions remove.Options

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		//if _, err := os.Stat(injectOptions.License); os.IsNotExist(err) {
		//	logger.Error().Msgf("invalid license path '%s'",injectOptions.License)
		//	_ = cmd.Help()
		//	os.Exit(0)
		//}
		//if _, err := os.Stat(injectOptions.Template); os.IsNotExist(err) {
		//	logger.Error().Msgf("missing template path '%s'",injectOptions.Template)
		//	_ = cmd.Help()
		//	os.Exit(0)
		//}

		//i := inject.New(ctx, removeOptions)
		//Command(i).MustRun()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

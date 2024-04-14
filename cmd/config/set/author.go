package set

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Set default author when filtering",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("author", args[0])
		viper.WriteConfig()
	},
}

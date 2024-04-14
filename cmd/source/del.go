package source

import (
	"slices"

	"github.com/sebcej/githis/aggregator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Remove selected source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sources := []aggregator.Source{}

		viper.UnmarshalKey("sources", &sources)

		sources = slices.DeleteFunc(sources, func(s aggregator.Source) bool { return s.Name == args[0] })

		viper.Set("sources", sources)
		viper.WriteConfig()
	},
}

func init() {

}

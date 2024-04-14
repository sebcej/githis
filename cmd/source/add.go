package source

import (
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/sebcej/githis/aggregator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add source folder. A source folder is considered the parent folder of all your projects",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		addedPath, err := filepath.Abs(args[1])
		cobra.CheckErr(err)

		// Check if folder exists and is accessible
		_, err = os.Stat(addedPath)
		cobra.CheckErr(err)

		newSource := aggregator.Source{args[0], addedPath}

		sources := []aggregator.Source{}

		viper.UnmarshalKey("sources", &sources)

		if slices.IndexFunc(sources, func(s aggregator.Source) bool { return s.Name == newSource.Name }) != -1 {
			cobra.CheckErr(errors.New("name already used"))
			return
		}

		sources = append(sources, newSource)

		viper.Set("sources", sources)
		viper.WriteConfig()
	},
}

func init() {

}

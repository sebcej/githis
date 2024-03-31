package source

import (
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "List all available source folders",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		addedPath, err := filepath.Abs(args[1])
		cobra.CheckErr(err)

		// Check if folder exists and is accessible
		_, err = os.Stat(addedPath)
		cobra.CheckErr(err)

		newSource := Source{args[0], addedPath}

		sources := []Source{}

		viper.UnmarshalKey("sources", &sources)

		if slices.IndexFunc[[]Source](sources, func(s Source) bool { return s.Name == newSource.Name }) != -1 {
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

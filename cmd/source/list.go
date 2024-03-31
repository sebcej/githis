package source

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available source folders",
	Run: func(cmd *cobra.Command, args []string) {
		sources := []Source{}

		viper.UnmarshalKey("sources", &sources)

		for _, source := range sources {
			fmt.Println(source.Name, "-", source.Path)
		}
	},
}

func init() {

}

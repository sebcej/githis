package source

import (
	"log"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "List all available source folders",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(cmd)
	},
}

func init() {

}

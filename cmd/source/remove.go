package source

import (
	"log"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove selected source",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(cmd)
	},
}

func init() {

}
